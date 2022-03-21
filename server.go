/*
 *
 *  * Copyright (C) 2022 The orange protocol Authors
 *  * This file is part of The orange library.
 *  *
 *  * The Orange is free software: you can redistribute it and/or modify
 *  * it under the terms of the GNU Lesser General Public License as published by
 *  * the Free Software Foundation, either version 3 of the License, or
 *  * (at your option) any later version.
 *  *
 *  * The orange is distributed in the hope that it will be useful,
 *  * but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *  * GNU Lesser General Public License for more details.
 *  *
 *  * You should have received a copy of the GNU Lesser General Public License
 *  * along with The orange.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/urfave/cli"

	"github.com/orange-protocol/orange-server-v1/auth"
	"github.com/orange-protocol/orange-server-v1/cache"
	"github.com/orange-protocol/orange-server-v1/cmd"
	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/graph"
	"github.com/orange-protocol/orange-server-v1/graph/generated"
	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/service"
	"github.com/orange-protocol/orange-server-v1/store"

	"github.com/muesli/cache2go"
)

const defaultPort = "8080"

func main() {
	if err := setupAPP().Run(os.Args); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func setupAPP() *cli.App {
	app := cli.NewApp()
	app.Usage = "oscore service"
	app.Action = startAgent
	app.Flags = []cli.Flag{
		cmd.LogLevelFlag,
		cmd.LogDirFlag,
		cmd.RpcUrlFlag,
		cmd.PortFlag,
	}
	app.Before = func(context *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}
	return app
}

func startAgent(ctx *cli.Context) {
	initLog(ctx)
	err := config.LoadConfig("./config.json")
	if err != nil {
		fmt.Println("error on load config")
		panic(err)
	}
	err = store.InitMysql(config.GlobalConfig.Db)
	if err != nil {
		fmt.Println("error on InitMysql")
		panic(err)
	}

	port := fmt.Sprintf("%d", ctx.GlobalInt(cmd.GetFlagName(cmd.PortFlag)))
	cache.GlobalCache = cache2go.Cache("SYSCACHE")

	err = service.InitDidService(config.GlobalConfig)

	err = service.InitSysDataService(config.GlobalConfig.SysDs)
	if err != nil {
		log.Errorf("init sys data service failed!:%s", err.Error())
		panic(err)
	}

	err = service.InitOntloginServcie()
	if err != nil {
		log.Errorf("Init Ontlogin Servcie failed!:%s", err.Error())
		panic(err)
	}

	err = service.InitEmailService(config.GlobalConfig.EmailConfig)
	if err != nil {
		log.Errorf("Init Email Servcie failed!:%s", err.Error())
		panic(err)
	}
	service.InitNftClaimService(config.GlobalConfig)
	taskservice := service.NewTaskService()
	//todo remove uuid_nonce records create_time before 30 min
	go taskservice.Run()
	go CheckLogSize(ctx)

	router := chi.NewRouter()
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		//AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedHeaders:   []string{"Authorization", "Content-Length", "X-CSRF-Token", "Token", "session", "X_Requested_With", "Accept", "Origin", "Host", "Connection", "Accept-Encoding", "Accept-Language", "DNT", "X-CustomHeader", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since", "Cache-Control", "Content-Type", "Pragma"},
		ExposedHeaders:   []string{"Content-Length", "token", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Cache-Control", "Content-Language", "Content-Type", "Expires", "Last-Modified", "Pragma", "FooBar"},
		AllowCredentials: false,
		MaxAge:           172800, // Maximum value not ignored by any of major browsers
		//Debug:true,
	}))
	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	//disable playground on production
	router.Get("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "files"))
	FileServer(router, "/files", filesDir)

	//go log.Fatal(http.ListenAndServe(":8088", http.FileServer(http.Dir("./files/path"))))
	//router.Handle("/file", http.FileServer(http.Dir("./files")) )
	log.Infof("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
	signalHandle()

}

func initLog(ctx *cli.Context) {
	logLevel := ctx.GlobalInt(cmd.GetFlagName(cmd.LogLevelFlag))
	disableLogFile := ctx.GlobalBool(cmd.GetFlagName(cmd.DisableLogFileFlag))
	if disableLogFile {
		log.InitLog(logLevel, log.Stdout)
	} else {
		logFileDir := ctx.String(cmd.GetFlagName(cmd.LogDirFlag))
		logFileDir = filepath.Join(logFileDir, "") + string(os.PathSeparator)
		log.InitLog(logLevel, logFileDir, log.Stdout)
	}
}

func CheckLogSize(ctx *cli.Context) {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			isNeedNewFile := log.CheckIfNeedNewFile()
			if isNeedNewFile {
				log.ClosePrintLog()
				log.InitLog(ctx.GlobalInt(cmd.GetFlagName(cmd.LogLevelFlag)), log.PATH, log.Stdout)
			}
		}
	}
}

func signalHandle() {
	var (
		ch = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			fmt.Println("get a signal: stop the rest gateway process", si.String())
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
