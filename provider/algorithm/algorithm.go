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

package algorithm

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/provider/common"
	"github.com/orange-protocol/orange-server-v1/store"
)

type HttpAlgorithmProvider struct {
	did    string
	entris map[string]*common.Method
	client *http.Client
}

func NewHttpAlgorithmProvider(ap *store.AlgorithmProvider) (*HttpAlgorithmProvider, error) {

	methods, err := store.MySqlDB.QueryAlgorithmProviderMethodByDid(ap.Did)
	if err != nil {
		log.Errorf("errors on QueryAlgorithmProviderMethodByDid :%s, error:%s", ap.Did, err.Error())
		return nil, err
	}

	entries := make(map[string]*common.Method)
	for _, method := range methods {
		entries[method.Method] = &common.Method{
			Url:        method.URL,
			Param:      method.Param,
			Result:     method.ResultSchema,
			HttpMethod: method.HttpMethod,
		}
	}

	return &HttpAlgorithmProvider{
		did:    ap.Did,
		entris: entries,
		client: new(http.Client),
	}, nil
}

func (ap *HttpAlgorithmProvider) Invoke(methodName string, paramMap map[string]interface{}) (interface{}, error) {
	method, ok := ap.entris[methodName]
	if !ok {
		return nil, fmt.Errorf("no method:%s found", methodName)
	}
	paramStr := ""
	encrypted := paramMap["%input"]
	if encrypted != nil {
		paramStr = paramMap["%input"].(string)
	}

	log.Debugf("paramStr:%s\n", paramStr)
	/////////////////
	url := method.Url
	log.Debugf("url:%s\n", url)

	httpmethod := strings.ToUpper(method.HttpMethod)
	log.Debugf("httpmethod:%s\n", httpmethod)

	var req *http.Request
	if httpmethod == "POST" {
		payload := strings.NewReader(paramStr)
		r, err := http.NewRequest(httpmethod, url, payload)
		if err != nil {
			log.Errorf("errors on NewRequest", err.Error())
			return nil, err
		}
		r.Header.Add("Content-Type", "application/json")
		req = r
	} else if httpmethod == "GET" {
		r, err := http.NewRequest(httpmethod, url+"?"+paramStr, nil)
		if err != nil {
			log.Errorf("errors on NewRequest", err.Error())
			return nil, err
		}
		req = r
	}

	res, err := ap.client.Do(req)
	if err != nil {
		log.Errorf("errors on Do client", err.Error())
		return nil, err
	}
	if err != nil {
		log.Errorf("errors on Post:%s, error:%s", method.Url, err.Error())
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		log.Errorf("errors on Post:%s, code:%d", method.Url, res.StatusCode)
		return nil, fmt.Errorf("post error:code:%d", res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

func (this *HttpAlgorithmProvider) VerifySig(body []byte) (bool, error) {
	return true, nil
}
func TransformAPType(aptype int) string {
	res := ""
	switch aptype {
	case common.AP_TYPE_OUTER:
		res = "OUTER PROVIDER"
	case common.AP_TYPE_WASM:
		res = "WASM PROVIDER"
	default:
		res = "unknown type"
	}
	return res
}
