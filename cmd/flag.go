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

package cmd

import (
	"strings"

	"github.com/urfave/cli"
)

const (
	DEFAULT_LOG_LEVEL           = 1
	DEFAULT_LOG_FILE_PATH       = "./Log/"
	DEFAULT_BLOCK_CHAIN_RPC_URL = "http://localhost:8545"
)

var (
	LogLevelFlag = cli.UintFlag{
		Name:  "loglevel",
		Usage: "Set the log level to `<level>` (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel",
		Value: DEFAULT_LOG_LEVEL,
	}
	LogDirFlag = cli.StringFlag{
		Name:  "log-dir",
		Usage: "log output to the file",
		Value: DEFAULT_LOG_FILE_PATH,
	}
	RpcUrlFlag = cli.StringFlag{
		Name:  "chain-rpc-url",
		Usage: "Set block chain rpc url",
		Value: DEFAULT_BLOCK_CHAIN_RPC_URL,
	}
	DisableLogFileFlag = cli.BoolFlag{
		Name:  "disable-log-file",
		Usage: "Discard log output to file",
	}
	PortFlag = cli.Uint64Flag{
		Name:  "server-port",
		Usage: "server port",
		Value: 8080,
	}
)

//GetFlagName deal with short flag, and return the flag name whether flag name have short name
func GetFlagName(flag cli.Flag) string {
	name := flag.GetName()
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.Split(name, ",")[0])
}
