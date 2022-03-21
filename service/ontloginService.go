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

	"github.com/google/uuid"
	"github.com/ontology-tech/ontlogin-sdk-go/did"
	"github.com/ontology-tech/ontlogin-sdk-go/did/ont"
	"github.com/ontology-tech/ontlogin-sdk-go/modules"
	"github.com/ontology-tech/ontlogin-sdk-go/sdk"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/store"

	eth2 "github.com/ontology-tech/ontlogin-sdk-go/did/eth"
)

type OntloginService struct {
	OntloginSdk *sdk.OntLoginSdk
}

var OloginService *OntloginService

func InitOntloginServcie() error {
	processors := make(map[string]did.DidProcessor)
	//todo refactor me if config changed
	ontProcessor, err := ont.NewOntProcessor(false, config.GlobalConfig.DidConf[0].URL, config.GlobalConfig.DidConf[0].DIDContract, "", "")
	if err != nil {
		return err
	}
	processors["ont"] = ontProcessor

	ethProcessor := eth2.NewEthProcessor()
	processors["eth"] = ethProcessor
	vcfilters := make(map[int][]*modules.VCFilter)
	vcfilters[modules.ACTION_CERTIFICATION] = []*modules.VCFilter{}
	vcfilters[modules.ACTION_AUTHORIZATION] = []*modules.VCFilter{}

	conf := &sdk.SDKConfig{
		Chain:      config.GlobalConfig.OntloginConf.Chain,
		Alg:        config.GlobalConfig.OntloginConf.Alg,
		ServerInfo: config.GlobalConfig.OntloginConf.ServerInfo,
		VCFilters:  vcfilters,
	}
	loginsdk, err := sdk.NewOntLoginSdk(conf, processors, GenUUID, CheckNonce)
	OloginService = &OntloginService{
		OntloginSdk: loginsdk,
	}
	return nil
}

func GenUUID(action int) string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		log.Errorf("NewUUID failed:%s", err.Error())
		return ""
	}
	err = store.MySqlDB.AddUUIDNonce(uuid.String(), action)
	if err != nil {
		log.Errorf("AddUUIDNonce failed:%s", err.Error())
		return ""
	}
	return uuid.String()
}

func CheckNonce(nonce string) (int, error) {
	action, err := store.MySqlDB.QueryUUIDAction(nonce)
	if err != nil {
		return -1, fmt.Errorf("no nonce found")
	}
	if action < 0 {
		return -1, fmt.Errorf("no nonce found")
	}

	return action, nil
}
