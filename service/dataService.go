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
	"encoding/json"
	"fmt"

	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/provider/common"
	"github.com/orange-protocol/orange-server-v1/provider/data"
	"github.com/orange-protocol/orange-server-v1/store"
)

//this is fixed provided by system
type SysDataService struct {
	SysDP  *data.HttpDataProvider
	Apikey string
	Did    string
}

var SysDS *SysDataService

type UserAssetInfo struct {
	Name         string `json:"name"`
	TokenAddress string `json:"tokenAddress"`
	Icon         string `json:"icon"`
	Chain        string `json:"chain"`
	Balance      string `json:"balance"`
	Price        string `json:"price"`
	Value        string `json:"value"`
}

type AssetParam struct {
	Chain   string `json:"chain"`
	Address string `json:"address"`
}

func InitSysDataService(did string) error {
	ds, err := NewSysDataService(did)
	if err != nil {
		return err
	}
	SysDS = ds
	return nil
}

func NewSysDataService(did string) (*SysDataService, error) {

	sysdp, err := store.MySqlDB.QueryDataProviderByDid(did)
	if err != nil {
		return nil, err
	}
	if sysdp == nil {
		return nil, fmt.Errorf("no dp with did:%s", did)
	}

	syshttpdp, err := data.NewHttpDataProvider(sysdp)
	if err != nil {
		return nil, err
	}
	return &SysDataService{
		SysDP:  syshttpdp,
		Apikey: sysdp.Apikey,
		Did:    did,
	}, nil
}

func (s *SysDataService) GetUserAssetsDetail(did string) ([]*UserAssetInfo, error) {
	userAddrs, err := store.MySqlDB.GetUserVisibleAddressInfo(did)
	if err != nil {
		return nil, err
	}

	sysdpm, err := store.MySqlDB.QueryDPMethodByDIDAndMethod(s.Did, "queryAssets")
	if err != nil {
		log.Errorf("errors on QueryDPMethodByDIDAndMethod:%s", err.Error())
		return nil, err
	}
	if sysdpm == nil {
		return nil, fmt.Errorf("system data provider is not existed")
	}

	paramStr, _, err := ParseInputParam(sysdpm.Param, s.Apikey, userAddrs, true, "", common.DP_TYPE_OUTER, "POST")
	if err != nil {
		log.Errorf("errors on ParseInputParam:%s", err.Error())
		return nil, err
	}

	res, err := s.SysDP.InvokeMethodWithParamStr("queryAssets", paramStr)
	if err != nil {
		return nil, err
	}
	return s.AnalyzeUserAssetInfoRes(res)
}

func (s *SysDataService) AnalyzeUserAssetInfoRes(in []byte) ([]*UserAssetInfo, error) {
	//todo analyze the result

	m := make(map[string]interface{})

	err := json.Unmarshal(in, &m)
	if err != nil {
		return nil, err
	}
	t, ok := m["data"]
	if !ok {
		return nil, fmt.Errorf("error result format no 'data' field")
	}
	if t == nil {
		return nil, fmt.Errorf("error result is nil")
	}

	t2, ok := t.(map[string]interface{})["queryAssets"]
	if !ok {
		return nil, fmt.Errorf("error result format no 'queryXdaysSum' field")
	}

	assets := t2.([]interface{})
	result := make([]*UserAssetInfo, 0)

	for _, as := range assets {
		tmpas := as.(map[string]interface{})
		t := &UserAssetInfo{
			Name:         tmpas["name"].(string),
			TokenAddress: tmpas["address"].(string),
			Icon:         tmpas["icon"].(string),
			Chain:        tmpas["chain"].(string),
			Balance:      tmpas["balance"].(string),
			Price:        tmpas["price"].(string),
			Value:        tmpas["value"].(string),
		}
		result = append(result, t)
	}

	return result, nil
}
