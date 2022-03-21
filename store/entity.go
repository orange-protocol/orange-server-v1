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

package store

import (
	"time"
)

var (
	TASK_STATUS_INIT          = 0
	TASK_STATUS_RESOLVING     = 1
	TASK_STATUS_DP_QUERYING   = 2
	TASK_STATUS_DP_FINISHED   = 3
	TASK_STATUS_DP_FAILED     = 4
	TASK_STATUS_DP_TIMEOUT    = 5
	TASK_STATUS_AP_RESOLVING  = 6
	TASK_STATUS_AP_QUERYING   = 7
	TASK_STATUS_AP_FAILED     = 8
	TASK_STATUS_AP_TIMEOUT    = 9
	TASK_STATUS_AP_FINISHED   = 10
	TASK_STATUS_VC_STARTING   = 11
	TASK_STATUS_VC_GENERATING = 12
	TASK_STATUS_VC_FAILED     = 13
	TASK_STATUS_DONE          = 14

	TASK_STATUS_PARTITIAL_FINISHED = 15
	TASK_STATUS_MUTLI_DP_FINISHED  = 16

	//BATCH_COUNT = 100

	AUTH_INFO_STATE_INIT     = 0
	AUTH_INFO_STATE_APPROVED = 1
	AUTH_INFO_STATE_REJECTED = 2

	CLAIM_NFT_STATUS_INIT      = 0
	CLAIM_NFT_STATUS_QUERYING  = 1
	CLAIM_NFT_STATUS_RESOLVING = 2
	CLAIM_NFT_STATUS_DONE      = 3
	CLAIM_NFT_STATUS_TIMEOUT   = 4

	METHOD_REMOVED   = "REMOVED"
	METHOD_PUBLISHED = "published"
	METHOD_REVOKED   = "revoking"
)

type DataProvider struct {
	Did             string
	Introduction    string
	CreateTime      time.Time
	DpType          int
	Name            string
	Apikey          string
	Title           string
	Provider        string
	InvokeFrequency int
	ApiState        int
	Author          string
	Popularity      int
	Delay           int
	Icon            string
}

type DPMethod struct {
	Did              string
	Method           string
	ParamSchema      string
	ResultSchema     string
	URL              string
	CompositeSetting string
	Param            string
	Name             string
	Description      string
	Invoked          int
	Latency          int
	CreateTime       time.Time
	Labels           *LabelsInfo
	HttpMethod       string
	Status           string
}

type DPMethodWithDPInfo struct {
	Dp       *DataProvider
	DpMethod *DPMethod
}

type APMethodWithAPInfo struct {
	Ap       *AlgorithmProvider
	ApMethod *APMethod
}

type UserAddressInfo struct {
	Did        string
	Chain      string
	Address    string
	Pubkey     string
	CreateTime time.Time
	Visible    bool
}

type TaskInfo struct {
	TaskId       int64
	UserDID      string
	ApDID        string
	ApName       string
	ApTitle      string
	ApIcon       string
	ApMethod     string
	APMethodName string
	DpDID        string
	DpName       string
	DpTitle      string
	DpIcon       string
	DpMethod     string
	DpMethodName string
	DpResult     string
	CreateTime   time.Time
	UpdateTime   time.Time
	TaskStatus   int
	TaskResult   string
	ResultFile   string
	IssueTxhash  string
	RevokeTxhash string
	CallerDid    string
	Comments     string
	TaskBindInfo string
}

type AlgorithmProvider struct {
	Did             string
	APType          int
	Name            string
	Introduction    string
	CreateTime      time.Time
	ApiKey          string
	Title           string
	Provider        string
	InvokeFrequency int
	ApiState        int
	Author          string
	Popularity      int
	Delay           int
	Icon            string
}

type APMethod struct {
	Did          string
	Method       string
	ParamSchema  string
	ResultSchema string
	URL          string
	Param        string
	Name         string
	Description  string
	Invoked      int
	Latency      int
	CreateTime   time.Time
	HttpMethod   string
	Labels       *LabelsInfo
	Status       string
}

type LabelsInfo struct {
	BlockChain string `json:"blockChain"`
	Category   string `json:"category"`
	Scenario   string `json:"scenario"`
}

type WasmCodeInfo struct {
	OwnerDID   string
	Address    string
	Code       string
	CreateTime time.Time
	Comments   string
}

type UserOscoreInfo struct {
	Did        string
	Oscore     int
	ApDid      string
	DpDid      string
	CreateTime time.Time
}

type TaskHistory struct {
	TaskId       int64
	UserDID      string
	ApDID        string
	ApMethod     string
	DpDID        string
	DpMethod     string
	CreateTime   time.Time
	UpdateTime   time.Time
	TaskStatus   int
	TaskResult   string
	IssueTxhash  string
	RevokeTxhash string
	callerDid    string
}

type AuthInfo struct {
	Did           string
	AppName       string
	DataAuth      string
	AlgorithmAuth string
	State         int
}

type GenNFT struct {
	Chain           string
	ContractAddress string
	Name            string
	Description     string
	TokenID         string
	UserDID         string
	WalletAddress   string
	CreateTime      time.Time
}

type UserBasicInfo struct {
	Did        string
	NickName   string
	Avatar     string
	Email      string
	CreateTime time.Time
	UpdateTime time.Time
}

type ApplicationInfo struct {
	Did     string
	Name    string
	Website string
}

type ClaimNFTRecord struct {
	TxHash          string
	Chain           string
	ContractAddress string
	NftType         int64
	UserDID         string
	UserAddress     string
	CreateTime      time.Time
	UpdateTime      time.Time
	Status          int
	Result          string
	TokenId         int64
	Score           int64
}

type NFTSetting struct {
	Id          int
	Name        string
	Description string
	Image       string
	DpDID       string
	DpMethod    string
	ApDID       string
	ApMethod    string
	LowestScore int
	ValidDays   int
	Restriction string
	IssueBy     string
	AltImage    string
}

type FeedBack struct {
	UserDID    string
	Email      string
	Title      string
	Content    string
	CreateTime time.Time
}
