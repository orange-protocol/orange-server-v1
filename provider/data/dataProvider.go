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

package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/provider/common"
	"github.com/orange-protocol/orange-server-v1/store"
)

type MultiDPResult struct {
	Data interface{}
	Sig  string
}

type EncryptedDPResult struct {
	Data      interface{} `json:"data"`
	Encrypted string      `json:"encrypted"`
}

type HttpDataProvider struct {
	//name string
	did    string
	entris map[string]*common.Method
	client *http.Client
}

type DataProviderResp struct {
	Did  string `json:"did"`
	Data string `json:"data"`
	Sig  string `json:"sig"`
}

func NewHttpDataProvider(dp *store.DataProvider) (*HttpDataProvider, error) {
	methods, err := store.MySqlDB.QueryDataProviderMethodByDid(dp.Did)
	if err != nil {
		log.Errorf("errors on QueryDataProviderMethodByDid:%s, err:%s", dp.Did, err.Error())
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
	//give a long timeout
	client := &http.Client{Timeout: 60 * time.Minute}

	return &HttpDataProvider{
		did:    dp.Did,
		entris: entries,
		client: client,
	}, nil
}

//deprecated
func (dp *HttpDataProvider) InvokeMethod(methodName string, paramMap map[string]interface{}) ([]byte, error) {
	method, ok := dp.entris[methodName]
	if !ok {
		return nil, fmt.Errorf("no method:%s found", methodName)
	}

	/////////////////
	url := method.Url
	httpmethod := "POST"
	paramstr := common.ProcessParamMap(method, paramMap)
	log.Debugf("para:%s\n", paramstr)
	payload := strings.NewReader(paramstr)
	req, err := http.NewRequest(httpmethod, url, payload)
	if err != nil {
		log.Errorf("errors on NewRequest", err.Error())
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := dp.client.Do(req)
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

func (dp *HttpDataProvider) InvokeMethodWithParamStr(methodName string, paramStr string) ([]byte, error) {

	method, ok := dp.entris[methodName]
	if !ok {
		return nil, fmt.Errorf("no method:%s found", methodName)
	}

	/////////////////
	url := method.Url
	httpmethod := strings.ToUpper(method.HttpMethod)
	log.Debugf("para:%s\n", paramStr)
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

	res, err := dp.client.Do(req)
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

func (dp *HttpDataProvider) VerifyDataSig(body []byte) (bool, error) {
	//todo verify did sig
	return true, nil
}

func TransformDPType(dptype int) string {
	res := ""
	switch dptype {
	case common.DP_TYPE_OUTER:
		res = "OUTER PROVIDER"
	case common.DP_TYPE_COMPOSITE:
		res = "COMPOSITE PROVIDER"
	default:
		res = "unknown type"
	}
	return res
}

func (dp *HttpDataProvider) CheckUrl(urlpath string) error {
	_, err := url.ParseRequestURI(urlpath)
	return err
}

func GetServiceConfig(key string) (string, error) {
	sql := "select config_content from service_config where config_type=?"
	r, err := store.MySqlDB.Dbconnect.Query(sql, key)
	if err != nil {
		return "", err
	}
	defer r.Close()
	content := ""
	if r.Next() {
		err = r.Scan(&content)
		if err != nil {
			return "", err
		}
	}
	return content, nil
}

func UpdateHttpParam(params string) (string, error) {
	paramsMap := make(map[string]string, 0)
	err := json.Unmarshal([]byte(params), &paramsMap)
	if err != nil {
		return "", err
	}
	//param := "{\n\t\"$API_KEY\": \"a1bd9450-cabc-463e-8c1d-9c833b6c83e4\",\n\t\"$CHAIN_NAME\": \"eth\",\n\t\"$USER_ADDRESS\": \"0x45929D79A6DDdaA3C8154D4F245d17d1D80DbBcc\",\n\t\"$ENCRYPTED\": \"false\",\n\t\"$USER_DID\": \"did:ont:ARNzB1pTkG61NDwxwzJfNJF8BqcZjpfNev\",\n\t\"$DEFAULT_CHAIN_NAME\": \"eth\",\n\t\"$DEFAULT_USER_ADDRESS\": \"0x45929D79A6DDdaA3C8154D4F245d17d1D80DbBcc\",\n\t\"$AP_DID\": \"did:ont:ASwHNVY8jvtuJoxbFKDcz1KkVCxcYUvSj2\",\n\t\"$ORANGE_DID\": \"did:ont:AXdmdzbyf3WZKQzRtrNQwAR91ZxMUfhXkt\"\n}"
	serviceCfg, err := GetServiceConfig("dp_method")
	if err != nil {
		return "", err
	}
	paramMap := make(map[string]interface{}, 0)
	err = json.Unmarshal([]byte(serviceCfg), &paramMap)
	if err != nil {
		return "", err
	}
	parMap := make(map[string]interface{}, 0)
	for k, v := range paramsMap {
		if k == "encrypt" {
			parMap[k] = false
		} else {
			parMap[k] = paramMap[v]
		}
	}
	jsonString, err := json.Marshal(parMap)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

func (dp *HttpDataProvider) CheckHttpStatus(httpMethod, urlpath, params string) (bool, error) {
	var req *http.Request
	var err error
	if httpMethod == "GET" {
		req, err = http.NewRequest("GET", urlpath, nil)
		if err != nil {
			return false, err
		}
	} else if httpMethod == "POST" {
		httpParams, err := UpdateHttpParam(params)
		if err != nil {
			return false, err
		}
		req, err = http.NewRequest("POST", urlpath, bytes.NewReader([]byte(httpParams)))
		if err != nil {
			return false, err
		}
	} else {
		return false, fmt.Errorf("HTTP method err")
	}
	if req.Header == nil {
		return false, fmt.Errorf("check http err")
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := dp.client.Do(req)
	if err != nil {
		return false, err
	}
	if res == nil || res.Body == nil {
		return false, nil
	}
	if res.StatusCode != http.StatusOK {
		return false, fmt.Errorf("http status code:%d", res.StatusCode)
	}
	io.Copy(io.Discard, res.Body)
	defer res.Body.Close()
	return true, nil
}

func (dp *HttpDataProvider) ParseHttpResult(httpMethod, urlpath, params string) (string, error) {
	flag, err := dp.CheckHttpStatus(httpMethod, urlpath, params)
	if err != nil {
		return "", err
	}
	if flag == false {
		return "", fmt.Errorf("ParseHttpResult error")
	}
	var req *http.Request
	if httpMethod == "GET" {
		req, err = http.NewRequest("GET", urlpath, nil)
		if err != nil {
			return "", err
		}
	} else if httpMethod == "POST" {
		req, err = http.NewRequest("POST", urlpath, bytes.NewReader([]byte(params)))
		if err != nil {
			return "", err
		}
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := dp.client.Do(req)
	if err != nil {
		return "", err
	}
	if res == nil || res.Body == nil {
		return "", nil
	}
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("http status code:%d", res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		return "", err
	}
	mJson, err := json.Marshal(jsonMap)
	if err != nil {
		return "", err
	}
	return string(mJson), nil
}
