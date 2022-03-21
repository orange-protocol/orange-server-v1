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

package service

import (
	"fmt"
	"regexp"

	"strings"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/provider/common"
	"github.com/orange-protocol/orange-server-v1/store"
)

var (
	PARAM_KEY_APIKEY           = "$API_KEY"
	PARAM_KEY_ARRAY            = "$ARRAY"
	PARAM_KEY_CHAIN            = "$CHAIN_NAME"
	PARAM_KEY_USERADDRESS      = "$USER_ADDRESS"
	PARAM_KEY_ENCRYPTED        = "$ENCRYPTED"
	PARAM_KEY_USER_DID         = "$USER_DID"
	PARAM_DEFAULT_CHAIN_NAME   = "$DEFAULT_CHAIN_NAME"
	PARAM_DEFAULT_USER_ADDRESS = "$DEFAULT_USER_ADDRESS"
	PARAM_AP_DID               = "$AP_DID"
)

//
//func (ts *TaskService) createCommonParam(task *store.TaskInfo, isPOC bool) (string, error) {
//
//	return "", nil
//}

func getReplacedStr(str string, isGql bool, httpMethod string) string {
	if isGql {
		return fmt.Sprintf(`\"%s\"`, str)
	} else {
		if strings.EqualFold(httpMethod, "POST") {
			return fmt.Sprintf(`"%s"`, str)
		} else { //GET
			return fmt.Sprintf("%s", str)
		}
	}
}

func ParseInputParam(paramStr string, apikey string, userAddrInfos []*store.UserAddressInfo, isPoc bool, apdid string, dptype int, httpMethod string) (string, string, error) {

	commentsArr := make([]string, 0)

	isGql := true
	if dptype == common.DP_TYPE_CUSTOM {
		isGql = false
	}

	if strings.Contains(paramStr, PARAM_KEY_APIKEY) {
		//paramStr = strings.ReplaceAll(paramStr, PARAM_KEY_APIKEY, fmt.Sprintf(`\"%s\"`, apikey))
		paramStr = strings.ReplaceAll(paramStr, PARAM_KEY_APIKEY, getReplacedStr(apikey, isGql, httpMethod))
	}

	if strings.Contains(paramStr, PARAM_KEY_ENCRYPTED) {
		encrypted := "true"
		//if isPoc {
		//	encrypted = "false"
		//}
		paramStr = strings.ReplaceAll(paramStr, PARAM_KEY_ENCRYPTED, encrypted)
	}

	if strings.Contains(paramStr, PARAM_KEY_USER_DID) {
		paramStr = strings.ReplaceAll(paramStr, PARAM_KEY_USER_DID, fmt.Sprintf(`\"%s\"`, config.GlobalConfig.WasmExecutor.Did))
	}

	if strings.Contains(paramStr, PARAM_DEFAULT_CHAIN_NAME) {
		//paramStr = strings.ReplaceAll(paramStr, PARAM_DEFAULT_CHAIN_NAME, fmt.Sprintf(`\"%s\"`, userAddrInfos[0].Chain))
		paramStr = strings.ReplaceAll(paramStr, PARAM_DEFAULT_CHAIN_NAME, getReplacedStr(userAddrInfos[0].Chain, isGql, httpMethod))
	}
	if strings.Contains(paramStr, PARAM_DEFAULT_USER_ADDRESS) {
		//paramStr = strings.ReplaceAll(paramStr, PARAM_DEFAULT_USER_ADDRESS, fmt.Sprintf(`\"%s\"`, userAddrInfos[0].Address))
		paramStr = strings.ReplaceAll(paramStr, PARAM_DEFAULT_USER_ADDRESS, getReplacedStr(userAddrInfos[0].Address, isGql, httpMethod))
		commentsArr = append(commentsArr, userAddrInfos[0].Address)
	}
	if strings.Contains(paramStr, PARAM_AP_DID) {
		//paramStr = strings.ReplaceAll(paramStr,PARAM_AP_DID,fmt.Sprintf(`"%s"`,apdid))
		paramStr = strings.ReplaceAll(paramStr, PARAM_AP_DID, getReplacedStr(apdid, isGql, httpMethod))
	}

	r, err := regexp.Compile("\\$ARRAY\\[(.*?)\\]")
	if err != nil {
		return "", "", err
	}

	arr := r.FindAllStringSubmatch(paramStr, 2)
	for _, tmp := range arr {
		if len(tmp) > 0 {
			originStr := tmp[0]
			matchStr := tmp[1]
			strToRelpace := ""
			for _, ua := range userAddrInfos {
				//tmp := strings.ReplaceAll(matchStr, PARAM_KEY_CHAIN, fmt.Sprintf(`\"%s\"`, ua.Chain))
				//tmp = strings.ReplaceAll(tmp, PARAM_KEY_USERADDRESS, fmt.Sprintf(`\"%s\"`, ua.Address))
				tmp := strings.ReplaceAll(matchStr, PARAM_KEY_CHAIN, getReplacedStr(ua.Chain, isGql, httpMethod))
				tmp = strings.ReplaceAll(tmp, PARAM_KEY_USERADDRESS, getReplacedStr(ua.Address, isGql, httpMethod))
				commentsArr = append(commentsArr, ua.Address)

				if len(strToRelpace) == 0 {
					strToRelpace = tmp
				} else {
					strToRelpace = strToRelpace + "," + tmp
				}
			}
			paramStr = strings.ReplaceAll(paramStr, originStr, fmt.Sprintf("[%s]", strToRelpace))

		}
	}
	return paramStr, strings.Join(commentsArr, ","), nil
}
