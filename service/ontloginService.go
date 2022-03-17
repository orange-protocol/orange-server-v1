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
