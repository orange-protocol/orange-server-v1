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

package config

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/ontology-tech/ontlogin-sdk-go/modules"
)

const (
	ETH_MONITOR_INTERVAL = 3 * time.Second
	ETH_USEFUL_BLOCK_NUM = 6
)

var GlobalConfig *SysConfig

type DB struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	DBAddr   string `json:"DBAddr"`
	DbName   string `json:"DBName"`
}

type DidConf struct {
	Chain                    string `json:"chain"`
	Wallet                   string `json:"wallet"`
	Password                 string `json:"password"`
	URL                      string `json:"url"`
	DID                      string `json:"did"`
	DIDContract              string `json:"DIDContract"`
	CredentialExpirationDays int    `json:"CredentialExpirationDays"`
	Gasprice                 uint64 `json:"gasprice"`
	Gaslimit                 uint64 `json:"gaslimit"`
	Commit                   bool   `json:"commit"`
}

type WasmExecutor struct {
	Did      string `json:"did"`
	Address  string `json:"address"`
	Wallet   string `json:"wallet"`
	Password string `json:"password"`
}

type SysConfig struct {
	Chain                string                `json:"chain"`
	ChainRpc             string                `json:"chain_rpc"`
	SysDs                string                `json:"sys_data_service"`
	FilePath             string                `json:"file_path"`
	AvatarFilePath       string                `json:"avatar_file_path"`
	WasmExecutor         WasmExecutor          `json:"wasm_executor"`
	Db                   *DB                   `json:"db"`
	DidConf              []*DidConf            `json:"did_config"`
	OntloginConf         *OntloginConfig       `json:"ontlogin_config"`
	CallerAddrs          []string              `json:"outer_task_caller"`
	SigAuth              bool                  `json:"sig_auth"`
	MNftBlock            int                   `json:"m_nft_block"`
	METHBlock            int                   `json:"m_eth_block"`
	EmailConfig          *EmailConfig          `json:"mail_config"`
	SnapShotAssetsConfig *SnapShotAssetsConfig `json:"snapshot_assets_config"`
	BatchTaskCount       int                   `json:"batch_task_count"`
	TimeOutSeconds       int                   `json:"task_timeout_seconds"`
	ETHWallet            *ETHWallet            `json:"eth_wallet"`
	NFTConfig            *NFTConfig            `json:"nft_config"`
	GraphConfig          *GraphConfig          `json:"graph_config"`
	NFTTimeOutMinutes    int                   `json:"nft_timeout_minutes"`
}

type GraphConfig struct {
	Eth     string `json:"eth"`
	Bsc     string `json:"bsc"`
	Polygon string `json:"polygon"`
}

type NFTConfig struct {
	NftInfos map[string]NFTInfo `json:"nft_infos"`
}

type NFTInfo struct {
	ContractAddress string `json:"contract_address"`
	Rpc             string `json:"rpc"`
}

type ETHWallet struct {
	KeyStore string `json:"key_store_path"`
	Password string `json:"password"`
}

type OntloginConfig struct {
	Chain      []string            `json:"chain"`
	Alg        []string            `json:"alg"`
	ServerInfo *modules.ServerInfo `json:"serverInfo"`
}

type EmailConfig struct {
	MailAddress string `json:"mail_address"`
	Host        string `json:"host"`
	SmtpPort    int    `json:"smtp_port"`
	Password    string `json:"password"`
	Subject     string `json:"subject"`
	Content     string `json:"content"`
}

type SnapShotAssetsConfig struct {
	ApDID    string `json:"ap_did"`
	ApMethod string `json:"ap_method"`
	DpDID    string `json:"dp_did"`
	DpMethod string `json:"dp_method"`
}

func LoadConfig(filepath string) error {
	return loadConfigFromFile(filepath)
}

func loadConfigFromFile(filepath string) error {
	configData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	cfg := &SysConfig{}
	err = json.Unmarshal(configData, cfg)
	if err != nil {
		return err
	}
	GlobalConfig = cfg
	return nil
}
