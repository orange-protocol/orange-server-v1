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

package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bufio"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/google/uuid"
	ontsdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontology-tech/ontlogin-sdk-go/modules"
	"github.com/orange-protocol/orange-server-v1/auth"
	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/graph/generated"
	"github.com/orange-protocol/orange-server-v1/graph/model"
	"github.com/orange-protocol/orange-server-v1/jwt"
	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/provider/algorithm"
	"github.com/orange-protocol/orange-server-v1/provider/common"
	"github.com/orange-protocol/orange-server-v1/provider/data"
	"github.com/orange-protocol/orange-server-v1/service"
	"github.com/orange-protocol/orange-server-v1/store"
	"github.com/orange-protocol/orange-server-v1/utils"
)

func (r *mutationResolver) SubmitChanllenge(ctx context.Context, input model.ClientResponse) (string, error) {
	lr := &modules.ClientResponse{
		Ver:   input.Ver,
		Type:  input.Type,
		Did:   input.Did,
		Nonce: input.Nonce,
		Proof: &modules.Proof{
			Type:               input.Proof.Type,
			VerificationMethod: input.Proof.VerificationMethod,
			Created:            uint64(input.Proof.Created),
			Value:              input.Proof.Value,
		},
		//VPs:   nil,
	}
	err := service.OloginService.OntloginSdk.ValidateClientResponse(lr)
	if err != nil {
		log.Errorf("ValidateClientResponse failed:%s", err.Error())
		return "", fmt.Errorf("ValidateClientResponse failed,please make sure your did has been registered")
	}

	//bind if the did is ethereum address
	tmparr := strings.Split(input.Did, ":")
	if len(tmparr) == 3 && tmparr[1] == "etho" {
		address := "0x" + tmparr[2]
		exist, err := store.MySqlDB.IsUserBindLoginAddress(input.Did, address)
		if err != nil {
			return "", err
		}
		if !exist {
			err = store.MySqlDB.AddUserAddressInfo(&store.UserAddressInfo{
				Did:        input.Did,
				Chain:      "eth",
				Address:    address,
				Pubkey:     "",
				CreateTime: time.Now(),
				Visible:    true,
			})
			if err != nil {
				return "", err
			}
		}
	}

	return jwt.GenerateToken(input.Did)
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	udid := input.Did
	inputtime := input.Time
	sig := input.Sig

	if int64(math.Abs(float64(time.Now().Unix()-inputtime))) > 5*60 {
		return "", fmt.Errorf("input time should within 5 min")
	}

	v, err := service.SysDidService.ValidateSig(udid, fmt.Sprintf("%d", inputtime), sig)
	if err != nil {
		return "", err
	}
	if !v {
		return "", fmt.Errorf("verify sig failed!")
	}

	return jwt.GenerateToken(udid)
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	oldtoken := input.Token
	udid, err := jwt.ParseToken(oldtoken)
	if err != nil {
		return "", err
	}
	return jwt.GenerateToken(udid)
}

func (r *mutationResolver) BindAddress(ctx context.Context, input model.BindAddress) (string, error) {
	if err := auth.CheckLogin(ctx, input.Did); err != nil {
		log.Debugf("err:%s\n", err.Error())
		return "", err
	}
	v := utils.VerifyDIDSigs(input.Chain, input.Address, input.Did, input.Sig, input.Pubkey)
	if !v {
		return "", fmt.Errorf("invalid signature")
	}

	info := &store.UserAddressInfo{
		Did:     input.Did,
		Chain:   input.Chain,
		Address: input.Address,
		Pubkey:  input.Pubkey,
		Visible: true,
	}
	//log.Infof("%v",info)
	err := store.MySqlDB.AddUserAddressInfo(info)
	if err != nil {
		return err.Error(), err
	}
	return fmt.Sprintf("binded:%s-%s-%s", input.Did, input.Chain, input.Address), nil
}

func (r *mutationResolver) UnbindAddress(ctx context.Context, input model.UnBindAddress) (string, error) {
	if err := auth.CheckLogin(ctx, input.Did); err != nil {
		return "", err
	}
	log.Debugf("%s,%s,%s", input.Did, input.Chain, input.Address)
	err := store.MySqlDB.DeleteUserAddressInfo(input.Did, input.Chain, input.Address)
	if err != nil {
		return "", fmt.Errorf("DB error:%s", err.Error())
	}
	return fmt.Sprintf("unbinded:%s-%s-%s", input.Did, input.Chain, input.Address), nil
}

func (r *mutationResolver) AddTask(ctx context.Context, input model.AddTask, overwrite bool) (int64, error) {
	if err := auth.CheckLogin(ctx, input.UserDid); err != nil {
		return -1, err
	}
	matched, err := store.MySqlDB.CheckProviderMethodMatch(input.ApDid, input.ApMethod, input.DpDid, input.DpMethod)
	if err != nil {
		return -1, err
	}
	if !matched {
		return -1, fmt.Errorf("dp method result schema and ap method param schema are not matched")
	}

	//todo check bind_info
	bind_addrs := strings.Split(input.BindInfo, ",")
	infos, err := store.MySqlDB.GetUserAddressInfo(input.UserDid)
	if err != nil {
		return 0, err
	}
	userBindAddrs := make([]string, 0)
	for _, info := range infos {
		if info.Visible {
			userBindAddrs = append(userBindAddrs, info.Address)
		}
	}
	if len(userBindAddrs) == 0 {
		return 0, fmt.Errorf("no bind address found")
	}
	ubaStr := strings.Join(userBindAddrs, ",")

	for _, addr := range bind_addrs {
		if strings.Index(ubaStr, addr) < 0 {
			return 0, fmt.Errorf("address:%s is not binded or is not visible", addr)
		}
	}

	taskid, err := store.MySqlDB.AddTask(input.UserDid, input.ApDid, input.ApMethod, input.DpDid, input.DpMethod, input.UserDid, input.BindInfo)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if !overwrite {
				return -2, fmt.Errorf("same task with ap & dp already existed, please wait for complete and remove the task")
			} else {
				err = store.MySqlDB.DeleteTask(input.UserDid, input.ApDid, input.ApMethod, input.DpDid, input.DpMethod)
				if err != nil {
					return -1, err
				} else {
					taskid, err := store.MySqlDB.AddTask(input.UserDid, input.ApDid, input.ApMethod, input.DpDid, input.DpMethod, input.UserDid, input.BindInfo)
					if err != nil {
						return -1, err
					}
					return taskid, nil
				}
			}
		}
		return -1, fmt.Errorf("DB error:%s", err.Error())
	}
	return taskid, nil
}

func (r *mutationResolver) ChangeAddressVisible(ctx context.Context, userDid string, chain string, address string, visible bool) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	err := store.MySqlDB.SetAddressVisible(userDid, chain, address, visible)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) RequestEmailVCode(ctx context.Context, did string, email string) (bool, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return false, err
	}
	code, err := service.GlobalEmailService.RequestEmailVCode(did, email)
	if err != nil {
		return false, err
	}

	err = service.GlobalEmailService.SendVerificationCode(email, code)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) AddUserBasicInfo(ctx context.Context, input model.AddUserBasicInfoReq) (bool, error) {
	if err := auth.CheckLogin(ctx, input.Did); err != nil {
		return false, err
	}

	code, err := store.MySqlDB.GetEmailVerificationCode(input.Did, input.Email)
	if err != nil {
		return false, err
	}
	if !strings.EqualFold(code, input.Vcode) {
		return false, fmt.Errorf("invalid verification code")
	}

	err = store.MySqlDB.AddNewBasicInfo(input.Did, input.NickName, input.Email, "")
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateUserBasicInfo(ctx context.Context, input *model.UpdateUserBasicInfoReq) (bool, error) {
	if err := auth.CheckLogin(ctx, input.Did); err != nil {
		return false, err
	}
	fmt.Println("========================")
	if input.Vcode != nil {
		code, err := store.MySqlDB.GetEmailVerificationCode(input.Did, *input.Email)
		if err != nil {
			return false, err
		}

		if !strings.EqualFold(code, *input.Vcode) {
			return false, fmt.Errorf("invalid verification code")
		}
	}
	nickname, email := "", ""
	if input.NickName != nil {
		nickname = *input.NickName
	}
	if input.Email != nil {
		email = *input.Email
	}
	if len(nickname)+len(email) == 0 {
		return false, fmt.Errorf("nick name and email cann't be both null")
	}

	err := store.MySqlDB.UpdateBasicInfo(input.Did, nickname, email)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateApplicationInfo(ctx context.Context, input *model.ApplicationInfoReq) (bool, error) {
	if err := auth.CheckLogin(ctx, input.Did); err != nil {
		return false, err
	}

	err := store.MySqlDB.SaveApplicationInfo(input.Did, input.Name, input.Website)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) EditNickNameBasicInfo(ctx context.Context, did string, nickName string) (bool, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return false, err
	}
	err := store.MySqlDB.EditNickNameBasicInfo(did, nickName)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) EditEmailAddrBasicInfo(ctx context.Context, did string, email string, verifyCode string) (bool, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return false, err
	}
	err := store.MySqlDB.EditEmailAddrBasicInfo(did, email, verifyCode)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) EditAppNameAppInfo(ctx context.Context, did string, appName string) (bool, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return false, err
	}
	err := store.MySqlDB.EditAppNameAppInfo(did, appName)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) EditWebsiteAppInfo(ctx context.Context, did string, website string) (bool, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return false, err
	}
	err := store.MySqlDB.EditWebsiteAppInfo(did, website)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) RequestOscore(ctx context.Context, input *model.RequestOscoreReq) (int64, error) {
	appdid := input.Appdid
	authinfo, err := store.MySqlDB.QueryAuthInfoByDid(appdid)
	if err != nil {
		return -1, err
	}
	if authinfo == nil || authinfo.State != store.AUTH_INFO_STATE_APPROVED {
		return -1, fmt.Errorf("appdid :%s is has no authentication to access", appdid)
	}

	//verify appdid sig
	msg, err := json.Marshal(input.Data)
	if err != nil {
		return -1, err
	}

	f, err := service.SysDidService.ValidateSig(appdid, string(msg), input.Sig)
	if err != nil {
		return -1, err
	}
	if !f {
		return 0, fmt.Errorf("signature is invalid")
	}

	//verify sig
	userdid := input.Data.Userdid
	for _, w := range input.Data.Wallets {
		f := utils.VerifyDIDSigs(w.Chain, w.Address, userdid, w.Sig, w.Pubkey)
		if !f {
			return -1, fmt.Errorf("invalid sig of chain:%s,address:%s", w.Chain, w.Address)
		}
		err := store.MySqlDB.AddUserAddressInfo(&store.UserAddressInfo{
			Did:        userdid,
			Chain:      w.Chain,
			Address:    w.Address,
			CreateTime: time.Time{},
			Visible:    true,
		})
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				log.Infof("already added address: chain:%s - address:%s", w.Chain, w.Address)
			} else {
				log.Errorf("AddUserAddressInfo failed:%s", err.Error())
				return -1, err
			}
		}
	}

	t, err := store.MySqlDB.QueryTaskByUniqueKey(userdid, input.Data.Apdid, input.Data.Apmethod, input.Data.Dpdid, input.Data.Dpmethod)
	if err != nil {
		return -1, err
	}
	if t != nil {
		if input.Data.OverwriteOld == true {
			err = store.MySqlDB.DeleteTaskById(t.TaskId)
			if err != nil {
				return -1, err
			}
		} else {
			return -1, fmt.Errorf("task with same algorithm and data provider already exist!")
		}
	}

	matched, err := store.MySqlDB.CheckProviderMethodMatch(input.Data.Apdid, input.Data.Apmethod, input.Data.Dpdid, input.Data.Dpmethod)
	if err != nil {
		return -1, err
	}
	if !matched {
		return -1, fmt.Errorf("dp method result schema and ap method param schema are not matched")
	}

	taskid, err := store.MySqlDB.AddTask(userdid, input.Data.Apdid, input.Data.Apmethod, input.Data.Dpdid, input.Data.Dpmethod, appdid, "")
	if err != nil {
		log.Error("errors on AddTask:%v", input)
		return -1, err
	}

	return taskid, nil
}

func (r *mutationResolver) AddNewOuterTask(ctx context.Context, input model.AddNewOuterTaskReq) (int64, error) {
	if strings.HasPrefix(input.CallerDid, "did:etho:") {
		if !utils.AddressInArray(input.CallerDid, config.GlobalConfig.CallerAddrs) {
			return 0, fmt.Errorf("unauthorized caller did:%s", input.CallerDid)
		}
		if config.GlobalConfig.SigAuth {
			ethaddr := strings.ReplaceAll(input.CallerDid, "did:etho:", "0x")
			msg, err := json.Marshal(input.Data)
			if err != nil {
				return 0, err
			}

			verify := utils.ETHVerifySig(ethaddr, input.Sig, msg)
			if !verify {
				return 0, fmt.Errorf("verify sig failed")
			}
		}
		userdid := strings.ReplaceAll(input.Data.Wallet.Address, "0x", "did:etho:")
		taskid, err := store.MySqlDB.AddTask(userdid, input.Data.ApDid, input.Data.ApMethod, input.Data.DpDid, input.Data.DpMethod, input.CallerDid, input.Data.Wallet.Address)
		if err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				err = store.MySqlDB.DeleteTask(userdid, input.Data.ApDid, input.Data.ApMethod, input.Data.DpDid, input.Data.DpMethod)
				if err != nil {
					return -1, err
				} else {
					taskid, err := store.MySqlDB.AddTask(userdid, input.Data.ApDid, input.Data.ApMethod, input.Data.DpDid, input.Data.DpMethod, input.CallerDid, "")
					if err != nil {
						return -1, err
					}
					return taskid, nil
				}

			}
			return -1, fmt.Errorf("DB error")
		}
		return taskid, nil

	} else {
		return 0, fmt.Errorf("not a supported did: %s", input.CallerDid)
	}
}

func (r *mutationResolver) SaveThirdPartyVc(ctx context.Context, did string, mediaType string, credential string) (bool, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return false, err
	}
	err := store.MySqlDB.SaveThirdPartyVc(did, mediaType, credential)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UnBindThirdParty(ctx context.Context, did string, mediaType string) (bool, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return false, err
	}
	err := store.MySqlDB.UnBindThirdParty(did, mediaType)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) SaveUserKycInfo(ctx context.Context, did string, kyc string) (bool, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return false, err
	}
	err := store.MySqlDB.SaveUserKycInfo(did, kyc)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) ClaimNft(ctx context.Context, did string, chain string, nftType int64) (*model.ClaimNFTResponse, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return nil, err
	}

	nftsettings, err := store.MySqlDB.GetNFTSettingByCondition(fmt.Sprintf("where id = %d", nftType))
	if err != nil {
		return nil, err
	}
	if len(nftsettings) != 1 {
		return nil, fmt.Errorf("nft type:%d is not exist", nftType)
	}
	nftsetting := nftsettings[0]

	task, err := store.MySqlDB.QueryTaskByUniqueKey(did, nftsetting.ApDID, nftsetting.ApMethod, nftsetting.DpDID, nftsetting.DpMethod)
	if err != nil {
		return nil, err
	}
	if task == nil || task.TaskStatus != store.TASK_STATUS_DONE {
		code := 0
		if task == nil {
			code = 10
		} else if task.TaskStatus != store.TASK_STATUS_DONE {
			code = 20
		}
		return &model.ClaimNFTResponse{
			ErrorCode: int64(code),
			Address:   "",
			Param:     nil,
		}, nil
	}
	cvValid := int64(config.GlobalConfig.DidConf[0].CredentialExpirationDays)
	ValidTo := task.CreateTime.Add(time.Duration(cvValid*24) * time.Hour).Unix()
	if time.Now().Unix() > ValidTo {
		return &model.ClaimNFTResponse{
			ErrorCode: int64(60),
			Address:   "",
			Param:     nil,
		}, nil
	}
	//get user bind ether address
	//fixme now the address must be eth(bsc...)address

	addr := ""
	if strings.Contains(did, "ont") {
		visibleAddrs, err := store.MySqlDB.GetUserVisibleAddressInfo(did)
		if err != nil {
			return nil, err
		}
		if len(visibleAddrs) != 1 {
			return &model.ClaimNFTResponse{
				ErrorCode: int64(40),
				Address:   "",
				Param:     nil,
			}, nil
		}
		addr = visibleAddrs[0].Address
	} else {
		addr = strings.ReplaceAll(did, "did:etho:", "0x")
	}
	if !strings.EqualFold(addr, task.Comments) {
		return &model.ClaimNFTResponse{
			ErrorCode: int64(50),
			Address:   "",
			Param:     nil,
		}, nil
	}

	score, err := strconv.ParseInt(task.TaskResult, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("score is invalid:%s", task.TaskResult)
	}
	log.Debugf("taskid:%d,taskscore:%d,limit:%d\n", task.TaskId, score, nftsetting.LowestScore)
	if score < int64(nftsetting.LowestScore) {
		return &model.ClaimNFTResponse{
			ErrorCode: int64(30),
			Address:   "",
			Param:     nil,
		}, nil
	}

	hashbytes, err := service.GlobalNftClaimService.GetUserClaimHash(addr, int(nftType), uint64(score))
	if err != nil {
		return nil, fmt.Errorf("errors on GetUserClaimHash:%s", err.Error())
	}

	sigbytes, err := service.GlobalNftClaimService.SignMsg(hashbytes)
	if err != nil {
		return nil, fmt.Errorf("errors on SignMsg:%s", err.Error())
	}

	return &model.ClaimNFTResponse{
		ErrorCode: 0,
		Address:   addr,
		Param: &model.NFTParam{
			Hash:    hexutil.Encode(hashbytes),
			Sig:     hexutil.Encode(sigbytes),
			NftType: nftType,
			Score:   score,
		},
	}, nil
}

func (r *mutationResolver) SubmitClaimTxInfo(ctx context.Context, did string, chain string, addr string, nftType int64, txhash string) (bool, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return false, err
	}

	contractInfo, ok := config.GlobalConfig.NFTConfig.NftInfos[chain]
	if !ok {
		return false, fmt.Errorf("chain:%s is not supported", chain)
	}

	nftsetting, err := store.MySqlDB.GetNFTSettingByCondition(fmt.Sprintf(" where id=%d ", nftType))
	if err != nil {
		return false, err
	}
	if nftsetting == nil || len(nftsetting) == 0 {
		return false, fmt.Errorf("nft setting of type :%d is not exist", nftType)
	}

	task, err := store.MySqlDB.QueryTaskByUniqueKey(did, nftsetting[0].ApDID, nftsetting[0].ApMethod, nftsetting[0].DpDID, nftsetting[0].DpMethod)
	if err != nil {
		return false, err
	}
	if task == nil || task.TaskStatus != store.TASK_STATUS_DONE {
		return false, fmt.Errorf("task status is not DONE")
	}

	score, err := strconv.Atoi(task.TaskResult)
	if err != nil {
		return false, err
	}
	record := &store.ClaimNFTRecord{
		TxHash:          txhash,
		Chain:           chain,
		ContractAddress: contractInfo.ContractAddress,
		NftType:         nftType,
		UserDID:         did,
		UserAddress:     addr,
		Score:           int64(score),
	}

	err = store.MySqlDB.AddNewClaimRecord(record)
	return err == nil, err
}

func (r *mutationResolver) SaveDPInfo(ctx context.Context, userDid string, input *model.SubmitDpInfo) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	if input.DpDid != "" {
		return false, fmt.Errorf("Empty DID Input")
	}
	if !ontsdk.VerifyID(input.DpDid) {
		return false, fmt.Errorf("Invalid DID")
	}
	if input.DpName != "" {
		res, userDID, err := store.MySqlDB.CheckDpNameExist(input.DpName)
		if err != nil {
			return false, err
		}
		if input.DpInfoID == 0 {
			if res != nil && res.DpInfoID != 0 {
				return false, fmt.Errorf("DP name duplicate")
			}
		} else {
			if res != nil && userDID != userDid {
				return false, fmt.Errorf("DP name duplicate")
			}
		}
	}
	return store.MySqlDB.SaveDPInfo(userDid, input)
}

func (r *mutationResolver) SubmitDPInfo(ctx context.Context, userDid string, input *model.SubmitDpInfo) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	if input.DpDid == "" || input.DpName == "" || input.DpDesc == "" {
		return false, fmt.Errorf("Empty DID Input")
	}
	if len(input.DpName) > 50 || len(input.DpDesc) > 200 {
		return false, fmt.Errorf("DP Name or DP Description too long")
	}
	if !ontsdk.VerifyID(input.DpDid) {
		return false, fmt.Errorf("Invalid DID")
	}
	user_did, err := store.MySqlDB.QueryUserDIDByDPDID(input.DpDid, utils.VERIFYING)
	if err != nil {
		return false, err
	}
	if user_did != "" && user_did != userDid {
		return false, fmt.Errorf("DID already existed:%s", input.DpDid)
	}
	userDID, err := store.MySqlDB.QueryUserDIDByDPDID(input.DpDid, utils.PUBLISHED)
	if err != nil {
		return false, err
	}
	if userDID != "" && userDID != userDid {
		return false, fmt.Errorf("DID already existed:%s", input.DpDid)
	}
	dataProviderInfo, err := store.MySqlDB.QueryDataProviderTitle(input.DpName)
	if err != nil {
		return false, err
	}
	if dataProviderInfo != nil && dataProviderInfo.Did != input.DpDid {
		return false, fmt.Errorf("Duplicate DP Name:%s", input.DpName)
	}
	res, userDID, err := store.MySqlDB.CheckDpNameExist(input.DpName)
	if err != nil {
		return false, err
	}
	if input.DpInfoID == 0 {
		if res != nil && res.DpInfoID != 0 {
			return false, fmt.Errorf("Duplicate DP Name")
		}
	} else {
		if res != nil && userDID != userDid {
			return false, fmt.Errorf("Duplicate DP Name")
		}
	}
	dpInfo, err := store.MySqlDB.QueryDPInfo(userDid, utils.PUBLISHED)
	if err != nil {
		return false, err
	}
	if dpInfo != nil {
		if *dpInfo.DpDid != input.DpDid {
			return false, fmt.Errorf("DID doesn't match with the old one")
		}
	}
	return store.MySqlDB.SubmitDPInfo(userDid, input)
}

func (r *mutationResolver) RevokeDPInfo(ctx context.Context, userDid string, dpInfoID int64) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	return store.MySqlDB.RevokeDPInfo(userDid, utils.DRAFT, utils.VERIFYING, dpInfoID)
}

func (r *mutationResolver) UploadAvatar(ctx context.Context, file graphql.Upload) (string, error) {
	log.Debugf("content type:%s\n", file.ContentType)
	log.Debugf("file name:%s\n", file.Filename)
	log.Debugf("file size:%d\n", file.Size)
	if file.Size > 2*1024*1024 {
		return "", fmt.Errorf("File size large than 2M")
	}
	br := bufio.NewReader(file.File)
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	res := strings.Split(file.ContentType, "/")
	if len(res) != 2 {
		return "", fmt.Errorf("Invalid Filename")
	}
	if res[1] != "png" && res[1] != "jpeg" && res[1] != "jpg" && res[1] != "svg" {
		return "", fmt.Errorf("only support png,jpeg,jpg,svg image format")
	}
	fileName := uid.String() + "." + res[1]
	fo, err := os.Create("files/" + config.GlobalConfig.AvatarFilePath + fileName)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := fo.Close(); err != nil {
			log.Errorf("Upload Avatar err:%s", err)
			panic(err)
		}
	}()
	// make a write buffer
	w := bufio.NewWriter(fo)

	// make a buffer to keep chunks that are read
	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := br.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 {
			break
		}

		// write a chunk
		if _, err := w.Write(buf[:n]); err != nil {
			return "", err
		}
	}

	if err = w.Flush(); err != nil {
		return "", err
	}
	return config.GlobalConfig.FilePath + config.GlobalConfig.AvatarFilePath + fileName, nil
}

func (r *mutationResolver) SaveDPDataSet(ctx context.Context, userDid string, input *model.DpDataSetInfo) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	if input.DataSetName == "" || len(input.DataSetName) > 50 {
		return false, fmt.Errorf("Invalid DataSet Name")
	}
	if input.DataSetID == 0 {
		dpDatSet, err := store.MySqlDB.CheckDPDataSetNameExist(userDid, input.DataSetName)
		if err != nil {
			return false, err
		}
		if dpDatSet != nil {
			return false, fmt.Errorf("Duplicate DataSet Name")
		}
		if input.DataSetMethodName != "" {
			dataSet, err := store.MySqlDB.CheckDPDataSetMethodNameExist(userDid, input.DataSetMethodName)
			if err != nil {
				return false, err
			}
			if dataSet != nil {
				return false, fmt.Errorf("Duplicate DataSet MethodName")
			}
		}
	} else if input.DataSetID != 0 {
		dpDatSet, err := store.MySqlDB.CheckDPDataSetNameExist(userDid, input.DataSetName)
		if err != nil {
			return false, err
		}
		if dpDatSet != nil && dpDatSet.DataSetID != input.DataSetID {
			return false, fmt.Errorf("Duplicate DataSet Name")
		}
		if input.DataSetMethodName != "" {
			dataSet, err := store.MySqlDB.CheckDPDataSetMethodNameExist(userDid, input.DataSetMethodName)
			if err != nil {
				return false, err
			}
			if dataSet != nil && dataSet.DataSetID != input.DataSetID {
				return false, fmt.Errorf("Duplicate DataSet MethodName")
			}
		}

	}
	return store.MySqlDB.SubmitDPDataSet(input, userDid, utils.DRAFT)
}

func (r *mutationResolver) PublishDPDataSet(ctx context.Context, userDid string, input *model.DpDataSetInfo) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	if input.DataSetName == "" || input.DataSetMethodName == "" || input.DataSetDesc == "" || input.HTTPMethod == "" || input.HTTPURL == "" {
		return false, fmt.Errorf("Input content can not be empty")
	}
	if len(input.DataSetName) > 50 || len(input.DataSetMethodName) > 50 || len(input.DataSetDesc) > 200 {
		return false, fmt.Errorf("Input content too long")
	}
	if input.Labels != nil {
		if len(input.Labels.BlockChain) > 3 || len(input.Labels.Category) > 3 || len(input.Labels.Category) > 3 {
			return false, fmt.Errorf("Number of Labels larger than 3")
		}
	}
	err := service.SysDS.SysDP.CheckUrl(input.HTTPURL)
	if err != nil {
		return false, err
	}
	res, err := service.SysDS.SysDP.CheckHttpStatus(input.HTTPMethod, input.HTTPURL, input.Params)
	if err != nil {
		return false, err
	}
	if res == false {
		return false, nil
	}
	flag, err := store.MySqlDB.CheckDuplicateDPDataSetName(input, userDid)
	if err != nil {
		return false, err
	}
	if flag == false {
		return false, nil
	}
	return store.MySqlDB.SubmitDPDataSet(input, userDid, utils.VERIFYING)
}

func (r *mutationResolver) RevokeDPDataSet(ctx context.Context, userDid string, dataSetID int64) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	return store.MySqlDB.RevokeDPDataSet(userDid, utils.DRAFT, utils.VERIFYING, dataSetID)
}

func (r *mutationResolver) DeleteDPDataSet(ctx context.Context, userDid string, dataSetID int64) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	return store.MySqlDB.DeleteDPDataSet(userDid, dataSetID)
}

func (r *mutationResolver) RevokePublishedDPDataSet(ctx context.Context, userDid string, dataSetID int64) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	return store.MySqlDB.RevokePublishedDPDataSet(userDid, dataSetID)
}

func (r *mutationResolver) SaveAPInfo(ctx context.Context, userDid string, input *model.SubmitApInfo) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	if input.ApDid == "" {
		return false, fmt.Errorf("Input content is null")
	}
	if !ontsdk.VerifyID(input.ApDid) {
		return false, fmt.Errorf("Invalid AP DID")
	}
	if input.ApName != "" {
		res, userDID, err := store.MySqlDB.CheckApNameExist(input.ApName)
		if err != nil {
			return false, err
		}
		if input.ApInfoID == 0 {
			if res != nil && res.ApInfoID != 0 {
				return false, fmt.Errorf("Duplicate AP Name")
			}
		} else {
			if res != nil && userDID != userDid {
				return false, fmt.Errorf("Duplicate AP Name")
			}
		}
	}
	return store.MySqlDB.SaveAPInfo(userDid, input)
}

func (r *mutationResolver) SubmitAPInfo(ctx context.Context, userDid string, input *model.SubmitApInfo) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	if input.ApDid == "" || input.ApName == "" || input.ApDesc == "" {
		return false, fmt.Errorf("input content is null")
	}
	if len(input.ApName) > 50 || len(input.ApDesc) > 200 {
		return false, fmt.Errorf("input content too long")
	}
	if !ontsdk.VerifyID(input.ApDid) {
		return false, fmt.Errorf("apDID invalid")
	}
	user_did, err := store.MySqlDB.QueryUserDIDByAPDID(input.ApDid, utils.VERIFYING)
	if err != nil {
		return false, err
	}
	if user_did != "" && user_did != userDid {
		return false, fmt.Errorf("apDId has exist:%s", input.ApDid)
	}
	userDID, err := store.MySqlDB.QueryUserDIDByAPDID(input.ApDid, utils.PUBLISHED)
	if err != nil {
		return false, err
	}
	if userDID != "" && userDID != userDid {
		return false, fmt.Errorf("apDId has exist:%s", input.ApDid)
	}
	algorithmProviderInfo, err := store.MySqlDB.QueryAlgorithmProviderTitle(input.ApName)
	if err != nil {
		return false, err
	}
	if algorithmProviderInfo != nil && algorithmProviderInfo.Did != input.ApDid {
		return false, fmt.Errorf("ap name has exist:%s", input.ApName)
	}
	res, userDID, err := store.MySqlDB.CheckApNameExist(input.ApName)
	if err != nil {
		return false, err
	}
	if input.ApInfoID == 0 {
		if res != nil && res.ApInfoID != 0 {
			return false, fmt.Errorf("ap name duplicate")
		}
	} else {
		if res != nil && userDID != userDid {
			return false, fmt.Errorf("ap name duplicate")
		}
	}
	apInfo, err := store.MySqlDB.QueryAPInfo(userDid, utils.PUBLISHED)
	if err != nil {
		return false, err
	}
	if apInfo != nil {
		if *apInfo.ApDid != input.ApDid {
			return false, fmt.Errorf("SubmitAPInfo ap did not match")
		}
	}
	return store.MySqlDB.SubmitAPInfo(userDid, input)
}

func (r *mutationResolver) RevokeAPInfo(ctx context.Context, userDid string, apInfoID int64) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	return store.MySqlDB.RevokeAPInfo(userDid, utils.DRAFT, utils.VERIFYING, apInfoID)
}

func (r *mutationResolver) SaveAPDataSet(ctx context.Context, userDid string, input *model.ApDataSetInfo) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	if input.DataSetName == "" || len(input.DataSetName) > 50 {
		return false, fmt.Errorf("Invalid DataSetName")
	}
	if input.DataSetID == 0 {
		apDataSet, err := store.MySqlDB.CheckAPDataSetNameExist(userDid, input.DataSetName)
		if err != nil {
			return false, err
		}
		if apDataSet != nil {
			return false, fmt.Errorf("Duplicate DataSet Name")
		}
		if input.DataSetMethodName != "" {
			dataSet, err := store.MySqlDB.CheckAPDataSetMethodNameExist(userDid, input.DataSetMethodName)
			if err != nil {
				return false, err
			}
			if dataSet != nil {
				return false, fmt.Errorf("Duplicate DataSet MethodName")
			}
		}
	} else if input.DataSetID != 0 {
		apDatSet, err := store.MySqlDB.CheckAPDataSetNameExist(userDid, input.DataSetName)
		if err != nil {
			return false, err
		}
		if apDatSet != nil && apDatSet.DataSetID != input.DataSetID {
			return false, fmt.Errorf("Duplicate DataSet Name")
		}
		if input.DataSetMethodName != "" {
			dataSet, err := store.MySqlDB.CheckAPDataSetMethodNameExist(userDid, input.DataSetMethodName)
			if err != nil {
				return false, err
			}
			if dataSet != nil && dataSet.DataSetID != input.DataSetID {
				return false, fmt.Errorf("Duplicate DataSet MethodName")
			}
		}
	}
	return store.MySqlDB.SubmitAPDataSet(input, userDid, utils.DRAFT)
}

func (r *mutationResolver) PublishAPDataSet(ctx context.Context, userDid string, input *model.ApDataSetInfo) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	if input.DataSetName == "" || input.DataSetMethodName == "" || input.DataSetDesc == "" || input.HTTPMethod == "" || input.HTTPURL == "" {
		return false, fmt.Errorf("Input content can not be empty")
	}
	if len(input.DataSetName) > 50 || len(input.DataSetMethodName) > 50 || len(input.DataSetDesc) > 200 {
		return false, fmt.Errorf("Input content too long")
	}
	if input.Labels != nil {
		if len(input.Labels.BlockChain) > 3 || len(input.Labels.Category) > 3 || len(input.Labels.Category) > 3 {
			return false, fmt.Errorf("Number of Labels larger than 3")
		}
	}
	err := service.SysDS.SysDP.CheckUrl(input.HTTPURL)
	if err != nil {
		return false, err
	}

	flag, err := store.MySqlDB.CheckDuplicateAPDataSetName(input, userDid)
	if err != nil {
		return false, err
	}
	if flag == false {
		return false, nil
	}
	return store.MySqlDB.SubmitAPDataSet(input, userDid, utils.VERIFYING)
}

func (r *mutationResolver) RevokeAPDataSet(ctx context.Context, userDid string, dataSetID int64) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	return store.MySqlDB.RevokeAPDataSet(userDid, utils.DRAFT, utils.VERIFYING, dataSetID)
}

func (r *mutationResolver) DeleteAPDataSet(ctx context.Context, userDid string, dataSetID int64) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	return store.MySqlDB.DeleteAPDataSet(userDid, dataSetID)
}

func (r *mutationResolver) RevokePublishedAPDataSet(ctx context.Context, userDid string, dataSetID int64) (bool, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return false, err
	}
	return store.MySqlDB.RevokePublishedAPDataSet(userDid, dataSetID)
}

func (r *queryResolver) GetAllAlgorithmProviders(ctx context.Context, first *int64, skip *int64, where *model.AlgorithmProviderWhere, orderBy *string, orderDirection *string) (*model.GetAllAlgorithmProvidersResp, error) {
	strWhere := ""
	if where != nil && where.NameLike != nil {
		strWhere = " where title like '%" + *where.NameLike + "%' "
	}

	count, err := store.MySqlDB.QueryAllAlgorithmProvidersCountByCondition(strWhere)
	if err != nil {
		return nil, err
	}
	if count == int64(0) {
		return &model.GetAllAlgorithmProvidersResp{
			Total: 0,
			Data:  nil,
		}, nil
	}

	if orderBy != nil {
		strWhere = strWhere + " order by " + *orderBy
	}
	if orderDirection != nil {
		if !strings.EqualFold("asc", *orderDirection) && !strings.EqualFold("desc", *orderDirection) {
			return nil, fmt.Errorf("orderDirection only accept 'asc' or 'desc'")
		}
		strWhere = strWhere + " " + *orderDirection
	}
	if first != nil {
		strWhere = strWhere + fmt.Sprintf(" limit %d", *first)
	}
	if skip != nil {
		strWhere = strWhere + fmt.Sprintf(" offset %d", *skip)
	}

	aps, err := store.MySqlDB.QueryAllAlgorithmProvidersByCondition(strWhere)

	if err != nil {
		return nil, err
	}
	res := make([]*model.AlgorithmProvider, 0)
	for _, ap := range aps {
		t := &model.AlgorithmProvider{
			Name:            ap.Name,
			Type:            algorithm.TransformAPType(ap.APType),
			Introduction:    ap.Introduction,
			Did:             ap.Did,
			CreateTime:      ap.CreateTime.Unix(),
			Title:           ap.Title,
			Provider:        ap.Provider,
			InvokeFrequency: int64(ap.InvokeFrequency),
			APIState:        int64(ap.ApiState),
			Author:          ap.Author,
			Popularity:      int64(ap.Popularity),
			Delay:           int64(ap.Delay),
			Icon:            ap.Icon,
		}
		res = append(res, t)
	}
	return &model.GetAllAlgorithmProvidersResp{
		Total: count,
		Data:  res,
	}, nil
}

func (r *queryResolver) GetAllDataProviders(ctx context.Context, first *int64, skip *int64, where *model.DataProviderWhere, orderBy *string, orderDirection *string) (*model.GetAllDataProviders, error) {
	strWhere := ""
	if where != nil && where.NameLike != nil {
		strWhere = " where title like '%" + *where.NameLike + "%' "
	}
	count, err := store.MySqlDB.QueryAllDataProvidersCountByCondition(strWhere)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &model.GetAllDataProviders{
			Total: 0,
			Data:  nil,
		}, nil
	}

	if orderBy != nil {
		strWhere = strWhere + " order by " + *orderBy
	}
	if orderDirection != nil {
		if !strings.EqualFold("asc", *orderDirection) && !strings.EqualFold("desc", *orderDirection) {
			return nil, fmt.Errorf("orderDirection only accept 'asc' or 'desc'")
		}
		strWhere = strWhere + " " + *orderDirection
	}
	if first != nil {
		strWhere = strWhere + fmt.Sprintf(" limit %d", *first)
	}
	if skip != nil {
		strWhere = strWhere + fmt.Sprintf(" offset %d", *skip)
	}

	dps, err := store.MySqlDB.QueryAllDataProvidersByCondition(strWhere)
	if err != nil {
		return nil, err
	}
	res := make([]*model.DataProvider, 0)
	for _, dp := range dps {
		t := &model.DataProvider{
			Name:            dp.Name,
			Type:            data.TransformDPType(dp.DpType),
			Introduction:    dp.Introduction,
			Did:             dp.Did,
			CreateTime:      dp.CreateTime.Unix(),
			Title:           dp.Title,
			Provider:        dp.Provider,
			InvokeFrequency: int64(dp.InvokeFrequency),
			APIState:        int64(dp.ApiState),
			Author:          dp.Author,
			Popularity:      int64(dp.Popularity),
			Delay:           int64(dp.Delay),
			Icon:            dp.Icon,
		}
		res = append(res, t)
	}
	return &model.GetAllDataProviders{
		Total: count,
		Data:  res,
	}, nil
}

func (r *queryResolver) GetDataProvidersByAp(ctx context.Context, did string, method string) ([]*model.DPAndMethod, error) {
	dps, dms, err := store.MySqlDB.QueryDPAndMethodsByAP(did, method)
	if err != nil {
		return nil, err
	}
	res := make([]*model.DPAndMethod, 0)
	for i := 0; i < len(dps); i++ {
		tp := &model.DataProvider{
			Name:            dps[i].Name,
			Type:            data.TransformDPType(dps[i].DpType),
			Introduction:    dps[i].Introduction,
			Did:             dps[i].Did,
			CreateTime:      dps[i].CreateTime.Unix(),
			Title:           dps[i].Title,
			Provider:        dps[i].Provider,
			InvokeFrequency: int64(dps[i].InvokeFrequency),
			APIState:        int64(dps[i].ApiState),
			Author:          dps[i].Author,
			Popularity:      int64(dps[i].Popularity),
			Delay:           int64(dps[i].Delay),
			Icon:            dps[i].Icon,
		}
		isSupportMultiAddr := false
		if strings.Contains(dms[i].Param, "$ARRAY") {
			isSupportMultiAddr = true
		}
		tm := &model.ProviderMethod{
			Name:         dms[i].Method,
			ParamSchema:  dms[i].ParamSchema,
			ResultSchema: dms[i].ResultSchema,
			Title:        dms[i].Name,
			Description:  dms[i].Description,
			SupportMulti: isSupportMultiAddr,
		}
		tmp := &model.DPAndMethod{
			Dp:     tp,
			Method: tm,
		}
		res = append(res, tmp)
	}

	return res, nil
}

func (r *queryResolver) GetUserAssetBalance(ctx context.Context, did string) ([]*model.UserAsset, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return nil, err
	}
	//todo mock return
	res := make([]*model.UserAsset, 0)
	results, err := service.SysDS.GetUserAssetsDetail(did)
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		t := &model.UserAsset{
			Name:         result.Name,
			TokenAddress: result.TokenAddress,
			Icon:         result.Icon,
			Chain:        result.Chain,
			Balance:      result.Balance,
			Price:        result.Price,
			Value:        result.Value,
		}
		res = append(res, t)
	}

	return res, nil
}

func (r *queryResolver) GetUserTotalValue(ctx context.Context, did string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTokenPrice(ctx context.Context, input model.TokenPrice) (string, error) {
	return "1.00", nil
}

func (r *queryResolver) GetUserBindAddressInfo(ctx context.Context, input string) ([]*model.UserAddressInfo, error) {
	if err := auth.CheckLogin(ctx, input); err != nil {
		return nil, err
	}

	infos, err := store.MySqlDB.GetUserAddressInfo(input)
	if err != nil {
		return nil, err
	}

	res := make([]*model.UserAddressInfo, 0)
	for _, info := range infos {
		t := &model.UserAddressInfo{}
		t.Address = info.Address
		t.Chain = info.Chain
		t.CreateTime = info.CreateTime.Unix()
		t.Visible = info.Visible
		res = append(res, t)
	}

	return res, nil
}

func (r *queryResolver) GetUserVisibleBindAddressInfo(ctx context.Context, input string) ([]*model.UserAddressInfo, error) {
	if err := auth.CheckLogin(ctx, input); err != nil {
		return nil, err
	}

	infos, err := store.MySqlDB.GetUserAddressInfo(input)
	if err != nil {
		return nil, err
	}

	res := make([]*model.UserAddressInfo, 0)
	for _, info := range infos {
		if !info.Visible {
			continue
		}
		t := &model.UserAddressInfo{}
		t.Address = info.Address
		t.Chain = info.Chain
		t.CreateTime = info.CreateTime.Unix()
		t.Visible = info.Visible
		res = append(res, t)
	}

	return res, nil
}

func (r *queryResolver) GetUserTasks(ctx context.Context, first *int64, skip *int64, where *model.GetUserTasksWhere, orderBy *string, orderDirection *string) (*model.GetUserTasksResp, error) {
	did := where.UserDid
	if err := auth.CheckLogin(ctx, did); err != nil {
		return nil, err
	}

	strwhere := "where t1.user_did = '" + did + "' "
	if where.DpDid != nil {
		strwhere = strwhere + " and t1.dp_did = '" + *where.DpDid + "' "
	}
	if where.DpMethod != nil {
		strwhere = strwhere + " and t1.dp_method = '" + *where.DpMethod + "' "
	}
	if where.ApDid != nil {
		strwhere = strwhere + " and t1.ap_did = '" + *where.ApDid + "' "
	}
	if where.ApMethod != nil {
		strwhere = strwhere + " and t1.ap_method = '" + *where.ApMethod + "' "
	}

	count, err := store.MySqlDB.GetUserCredentialsCountByCondition(strwhere)
	if err != nil {
		return nil, err
	}
	res := &model.GetUserTasksResp{TotalCount: count}
	if count == 0 {
		return res, nil
	}

	if orderBy != nil {
		strwhere = strwhere + " order by t1." + *orderBy
	}
	if orderDirection != nil {
		if !strings.EqualFold("asc", *orderDirection) && !strings.EqualFold("desc", *orderDirection) {
			return nil, fmt.Errorf("orderDirection only accept 'asc' or 'desc'")
		}
		strwhere = strwhere + " " + *orderDirection
	}
	if first != nil {
		strwhere = strwhere + fmt.Sprintf(" limit %d", *first)
	}
	if skip != nil {
		strwhere = strwhere + fmt.Sprintf(" offset %d", *skip)
	}

	tasks, err := store.MySqlDB.GetUserCredentialsByCondition(strwhere)
	if err != nil {
		return nil, err
	}

	cvValid := int64(config.GlobalConfig.DidConf[0].CredentialExpirationDays)
	credentials := make([]*model.UserTasks, 0)
	for _, task := range tasks {
		filepath := ""
		if len(task.ResultFile) > 0 {
			filepath = config.GlobalConfig.FilePath + task.ResultFile
		}
		t := &model.UserTasks{
			TaskID:       fmt.Sprintf("%d", task.TaskId),
			UserDid:      task.UserDID,
			ApDid:        task.ApDID,
			ApName:       task.ApTitle,
			ApIcon:       task.ApIcon,
			ApMethod:     task.ApMethod,
			ApMethodName: task.APMethodName,
			DpDid:        task.DpDID,
			DpName:       task.DpTitle,
			DpIcon:       task.DpIcon,
			DpMethod:     task.DpMethod,
			DpMethodName: task.DpMethodName,
			CreateTime:   fmt.Sprintf("%d", task.CreateTime.Unix()),
			UpdateTime:   fmt.Sprintf("%d", task.UpdateTime.Unix()),
			TaskStatus:   common.TransformTaskStatus(task.TaskStatus),
			TaskResult:   &task.TaskResult,
			ResultFile:   &filepath,
			IssueTxhash:  &task.IssueTxhash,
			ValidTo:      task.CreateTime.Add(time.Duration(cvValid*24) * time.Hour).Unix(),
		}
		credentials = append(credentials, t)
	}

	res.Tasks = credentials
	return res, nil
}

func (r *queryResolver) QueryTaskExist(ctx context.Context, input *model.AddTask) (bool, error) {
	if err := auth.CheckLogin(ctx, input.UserDid); err != nil {
		return false, err
	}

	t, err := store.MySqlDB.QueryTaskByUniqueKey(input.UserDid, input.ApDid, input.ApMethod, input.DpDid, input.DpMethod)
	if err != nil {
		return false, err
	}
	if t != nil {
		return true, nil
	}
	return false, nil
}

func (r *queryResolver) GetLatestOscore(ctx context.Context, did string) (*model.UserLatestOscore, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return nil, err
	}
	lo, err := store.MySqlDB.QueryUserLatestOScoreInfo(did)
	if err != nil {
		log.Errorf("errors on QueryUserLatestOScoreInfo:%s", err.Error())
		return nil, err
	}
	if lo == nil {
		return &model.UserLatestOscore{
			Oscore:     int64(0),
			ApDid:      "",
			DpDid:      "",
			CreateTime: time.Now().Unix(),
		}, nil
	}
	return &model.UserLatestOscore{
		Oscore:     int64(lo.Oscore),
		ApDid:      lo.ApDid,
		DpDid:      lo.DpDid,
		CreateTime: lo.CreateTime.Unix(),
	}, nil
}

func (r *queryResolver) GetAlgorithmMethods(ctx context.Context, did string) ([]*model.ProviderMethod, error) {
	methods, err := store.MySqlDB.QueryAlgorithmProviderMethodByDid(did)
	if err != nil {
		return nil, err
	}
	res := make([]*model.ProviderMethod, 0)
	for _, method := range methods {
		t := &model.ProviderMethod{
			Name:         method.Method,
			ParamSchema:  method.ParamSchema,
			ResultSchema: method.ResultSchema,
			Title:        method.Name,
			Description:  method.Description,
		}
		res = append(res, t)
	}
	return res, nil
}

func (r *queryResolver) GetDataMethods(ctx context.Context, did string) ([]*model.ProviderMethod, error) {
	methods, err := store.MySqlDB.QueryDataProviderMethodByDid(did)
	if err != nil {
		return nil, err
	}
	res := make([]*model.ProviderMethod, 0)
	for _, method := range methods {
		t := &model.ProviderMethod{
			Name:         method.Method,
			ParamSchema:  method.ParamSchema,
			ResultSchema: method.ResultSchema,
			Title:        method.Name,
			Description:  method.Description,
		}
		res = append(res, t)
	}
	return res, nil
}

func (r *queryResolver) GetAlgorithmProvider(ctx context.Context, did string) (*model.AlgorithmProvider, error) {
	ap, err := store.MySqlDB.QueryAlgorithmProviderByDid(did)
	if err != nil {
		return nil, err
	}
	if ap == nil {
		return nil, nil
	}

	appopularity, err := store.MySqlDB.GetAPPopularity(did)
	if err != nil {
		log.Errorf("GetAPPopularity failed:%s", err.Error())
		return nil, err
	}

	return &model.AlgorithmProvider{
		Name:            ap.Name,
		Type:            algorithm.TransformAPType(ap.APType),
		Introduction:    ap.Introduction,
		Did:             ap.Did,
		CreateTime:      ap.CreateTime.Unix(),
		Title:           ap.Title,
		Provider:        ap.Provider,
		InvokeFrequency: int64(ap.InvokeFrequency),
		APIState:        int64(ap.ApiState),
		Author:          ap.Author,
		Popularity:      appopularity,
		Delay:           int64(ap.Delay),
		Icon:            ap.Icon,
	}, nil
}

func (r *queryResolver) GetDataProvider(ctx context.Context, did string) (*model.DataProvider, error) {
	dp, err := store.MySqlDB.QueryDataProviderByDid(did)
	if err != nil {
		return nil, err
	}
	if dp == nil {
		return nil, nil
	}

	dppopularity, err := store.MySqlDB.GetDPPopularity(did)
	if err != nil {
		log.Errorf("GetDPPopularity failed:%s", err.Error())
		return nil, err
	}

	return &model.DataProvider{
		Name:            dp.Name,
		Type:            data.TransformDPType(dp.DpType),
		Introduction:    dp.Introduction,
		Did:             dp.Did,
		CreateTime:      dp.CreateTime.Unix(),
		Title:           dp.Title,
		Provider:        dp.Provider,
		InvokeFrequency: int64(dp.InvokeFrequency),
		APIState:        int64(dp.ApiState),
		Author:          dp.Author,
		Popularity:      dppopularity,
		Delay:           int64(dp.Delay),
		Icon:            dp.Icon,
	}, nil
}

func (r *queryResolver) GetUserTask(ctx context.Context, key string, taskID int64) (*model.UserTasks, error) {
	f, err := store.MySqlDB.CheckAPIKey(key)
	if err != nil {
		return nil, err
	}
	if !f {
		return nil, fmt.Errorf("apikey is invalid")
	}

	task, err := store.MySqlDB.QueryTaskByPK(taskID)
	if err != nil {
		log.Errorf("errors on QueryTaskByPK:%s", err.Error())
		return nil, err
	}
	if task == nil {
		return nil, nil
	}

	nfttype := 0
	nftsetting, err := store.MySqlDB.GetNFTSettingByCondition(fmt.Sprintf(" where ap_did='%s' and ap_method='%s' ", task.ApDID, task.ApMethod))
	if err != nil {
		log.Errorf("erros on GetNFTSettingByCondition:%s", err.Error())
	}
	if nftsetting == nil || len(nftsetting) == 0 {
		nfttype = 0
	} else {
		nfttype = nftsetting[0].Id
	}

	filepath := ""
	if len(task.ResultFile) > 0 {
		filepath = config.GlobalConfig.FilePath + task.ResultFile
	}
	cvValid := config.GlobalConfig.DidConf[0].CredentialExpirationDays
	return &model.UserTasks{
		TaskID:       fmt.Sprintf("%d", task.TaskId),
		UserDid:      task.UserDID,
		ApDid:        task.ApDID,
		ApName:       task.ApTitle,
		ApIcon:       task.ApIcon,
		ApMethod:     task.ApMethod,
		ApMethodName: task.APMethodName,
		DpDid:        task.DpDID,
		DpName:       task.ApTitle,
		DpIcon:       task.DpIcon,
		DpMethod:     task.DpMethod,
		DpMethodName: task.DpMethodName,
		CreateTime:   fmt.Sprintf("%d", task.CreateTime.Unix()),
		UpdateTime:   fmt.Sprintf("%d", task.UpdateTime.Unix()),
		TaskStatus:   common.TransformTaskStatus(task.TaskStatus),
		//todo test me if nil case
		TaskResult:        &task.TaskResult,
		ResultFile:        &filepath,
		IssueTxhash:       &task.IssueTxhash,
		ValidTo:           task.CreateTime.Add(time.Duration(cvValid*24) * time.Hour).Unix(),
		NftType:           int64(nfttype),
		InvolvedAddresses: task.Comments,
	}, nil
}

func (r *queryResolver) GetAlgorithmProviderMethod(ctx context.Context, did string, name string) (*model.ProviderMethod, error) {
	am, err := store.MySqlDB.QueryAPMethodByDIDAndMethod(did, name)
	if err != nil {
		log.Errorf("QueryAPMethodByDIDAndMethod error:%s", err.Error())
		return nil, err
	}
	if am == nil {
		return nil, nil
	}
	nfttype := 0
	nftsetting, err := store.MySqlDB.GetNFTSettingByCondition(fmt.Sprintf(" where ap_did='%s' and ap_method='%s' ", did, name))
	if err != nil {
		log.Errorf("erros on GetNFTSettingByCondition:%s", err.Error())
	}
	if nftsetting == nil || len(nftsetting) == 0 {
		nfttype = 0
	} else {
		nfttype = nftsetting[0].Id
	}
	hasDataSet, err := store.MySqlDB.CheckAPDataSetExistByAPDID(did)
	if err != nil {
		return nil, err
	}
	providerMethod := &model.ProviderMethod{
		Name:         am.Method,
		ParamSchema:  am.ParamSchema,
		ResultSchema: am.ResultSchema,
		Title:        am.Name,
		Description:  am.Description,
		CreateTime:   am.CreateTime.Unix(),
		NftType:      int64(nfttype),
		TotalUsed:    int64(am.Invoked),
		HasDataSet:   hasDataSet,
		Status:       am.Status,
		Labels: &model.LabelsInfos{
			BlockChain: make([]string, 0),
			Category:   make([]string, 0),
			Scenario:   make([]string, 0),
		},
	}
	if am.Labels != nil {
		if am.Labels.BlockChain != "" {
			providerMethod.Labels.BlockChain = strings.Split(am.Labels.BlockChain, ",")
		}
		if am.Labels.Category != "" {
			providerMethod.Labels.Category = strings.Split(am.Labels.Category, ",")
		}
		if am.Labels.Scenario != "" {
			providerMethod.Labels.Scenario = strings.Split(am.Labels.Scenario, ",")
		}
	}
	return providerMethod, nil
}

func (r *queryResolver) GetDataProviderMethod(ctx context.Context, did string, name string) (*model.ProviderMethod, error) {
	dm, err := store.MySqlDB.QueryDPMethodByDIDAndMethod(did, name)
	if err != nil {
		log.Errorf("QueryDPMethodByDIDAndMethod error:%s", err.Error())
		return nil, err
	}
	if dm == nil {
		return nil, nil
	}
	hasDataSet, err := store.MySqlDB.CheckDPDataSetExistByDPDID(did)
	if err != nil {
		return nil, err
	}
	if dm.Status == store.METHOD_REMOVED {
		return nil, fmt.Errorf("has revoked")
	}
	providerMethod := &model.ProviderMethod{
		Name:         dm.Method,
		ParamSchema:  dm.ParamSchema,
		ResultSchema: dm.ResultSchema,
		Title:        dm.Name,
		Description:  dm.Description,
		CreateTime:   dm.CreateTime.Unix(),
		TotalUsed:    int64(dm.Invoked),
		HasDataSet:   hasDataSet,
		Status:       dm.Status,
		Labels: &model.LabelsInfos{
			BlockChain: make([]string, 0),
			Category:   make([]string, 0),
			Scenario:   make([]string, 0),
		},
	}
	if dm.Labels != nil {
		if dm.Labels.BlockChain != "" {
			providerMethod.Labels.BlockChain = strings.Split(dm.Labels.BlockChain, ",")
		}
		if dm.Labels.Category != "" {
			providerMethod.Labels.Category = strings.Split(dm.Labels.Category, ",")
		}
		if dm.Labels.Scenario != "" {
			providerMethod.Labels.Scenario = strings.Split(dm.Labels.Scenario, ",")
		}
	}
	return providerMethod, nil
}

func (r *queryResolver) GetUserGenNFTCount(ctx context.Context, did string) (*model.GenNFTCountResp, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return nil, err
	}

	count, err := store.MySqlDB.GetGenNFTCountByDID(did)
	return &model.GenNFTCountResp{
		Count: count,
	}, err
}

func (r *queryResolver) GetUserGenReputationCount(ctx context.Context, did string) (*model.GenReputationCountResp, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return nil, err
	}
	count, err := store.MySqlDB.QueryTaskHistoryCountByUserDID(did)
	return &model.GenReputationCountResp{Count: count}, err
}

func (r *queryResolver) GetUserCredentials(ctx context.Context, first *int64, skip *int64, where model.UserCredntialWhere, orderBy *string, orderDirection *string) (*model.UserCredentials, error) {
	did := where.UserDid
	if err := auth.CheckLogin(ctx, did); err != nil {
		return nil, err
	}

	strwhere := "where t1.user_did = '" + did + "' and t1.task_status = 14 "
	if where.DpDid != nil {
		strwhere = strwhere + " and t1.dp_did = '" + *where.DpDid + "' "
	}
	if where.DpMethod != nil {
		strwhere = strwhere + " and t1.dp_method = '" + *where.DpMethod + "' "
	}
	if where.ApDid != nil {
		strwhere = strwhere + " and t1.ap_did = '" + *where.ApDid + "' "
	}
	if where.ApMethod != nil {
		strwhere = strwhere + " and t1.ap_method = '" + *where.ApMethod + "' "
	}

	count, err := store.MySqlDB.GetUserCredentialsCountByCondition(strwhere)
	if err != nil {
		return nil, err
	}
	res := &model.UserCredentials{TotalCount: count}
	if count == 0 {
		return res, nil
	}

	if orderBy != nil {
		strwhere = strwhere + " order by t1." + *orderBy
	}
	if orderDirection != nil {
		if !strings.EqualFold("asc", *orderDirection) && !strings.EqualFold("desc", *orderDirection) {
			return nil, fmt.Errorf("orderDirection only accept 'asc' or 'desc'")
		}
		strwhere = strwhere + " " + *orderDirection
	}
	if first != nil {
		strwhere = strwhere + fmt.Sprintf(" limit %d", *first)
	}
	if skip != nil {
		strwhere = strwhere + fmt.Sprintf(" offset %d", *skip)
	}

	tasks, err := store.MySqlDB.GetUserCredentialsByCondition(strwhere)
	if err != nil {
		return nil, err
	}

	cvValid := int64(config.GlobalConfig.DidConf[0].CredentialExpirationDays)
	credentials := make([]*model.UserCredential, 0)
	for _, task := range tasks {
		t := &model.UserCredential{
			Score:        task.TaskResult,
			DpName:       task.DpTitle,
			DpMethodName: task.DpMethodName,
			ApName:       task.ApTitle,
			ApMethodName: task.APMethodName,
			CreateTime:   task.CreateTime.Format("2006-01-02"),
			ValidTo:      task.CreateTime.Add(time.Duration(cvValid*24) * time.Hour).Unix(),
		}
		credentials = append(credentials, t)
	}

	res.Data = credentials
	return res, nil
}

func (r *queryResolver) GetAllDataProviderMethod(ctx context.Context, first *int64, skip *int64, where *model.DataProviderMethodWhere, orderBy *string, orderDirection *string, labels model.LabelsInfo) (*model.GetAllDataProviderMethodsResp, error) {
	strWhere := ""
	if where != nil {
		if where.NameLike != nil {
			strWhere = " where t1.name like '%" + *where.NameLike + "%' "
		}
		if where.DidLike != nil {
			if len(strWhere) == 0 {
				strWhere = " where t1.did like '%" + *where.DidLike + "%' "
			} else {
				strWhere = strWhere + " and t1.did like '%" + *where.DidLike + "%' "
			}
		}
		strWhere = strWhere + " and t1.status !='" + store.METHOD_REMOVED + "' "
	}
	strWhere = strWhere + " and composite_setting = 'NONE' "
	whereStr := " and"
	for _, blockChainLabels := range labels.BlockChain {
		whereStr += " block_chain_labels like '%" + blockChainLabels + "%' or "
	}
	for _, catgegoryLabels := range labels.Category {
		whereStr += " category_labels like '%" + catgegoryLabels + "%' or "
	}
	for _, catgegoryLabels := range labels.Scenario {
		whereStr += " category_labels like '%" + catgegoryLabels + "%' or "
	}
	whereStr = strings.TrimRight(whereStr, "or ")
	if whereStr == " and" {
		whereStr = " "
	}
	count, err := store.MySqlDB.QueryAllDPMethodCountByCondition(strWhere + " " + whereStr)
	if err != nil {
		return nil, err
	}
	if count == int64(0) {
		return &model.GetAllDataProviderMethodsResp{
			Total: 0,
			Data:  nil,
		}, nil
	}
	strWhere = strWhere + " " + whereStr
	if orderBy != nil {
		strWhere = strWhere + " order by t1." + *orderBy
	}
	if orderDirection != nil {
		if !strings.EqualFold("asc", *orderDirection) && !strings.EqualFold("desc", *orderDirection) {
			return nil, fmt.Errorf("orderDirection only accept 'asc' or 'desc'")
		}
		strWhere = strWhere + " " + *orderDirection
	}
	if first != nil {
		strWhere = strWhere + fmt.Sprintf(" limit %d", *first)
	}
	if skip != nil {
		strWhere = strWhere + fmt.Sprintf(" offset %d", *skip)
	}

	dpms, err := store.MySqlDB.QueryAllDPMethodByCondition(strWhere)
	if err != nil {
		return nil, err
	}

	res := make([]*model.DPMethodWithDp, 0)
	for _, dpm := range dpms {
		dpRes := &model.DPMethodWithDp{
			Did:               dpm.Dp.Did,
			Name:              dpm.Dp.Title,
			Method:            dpm.DpMethod.Method,
			MethodName:        dpm.DpMethod.Name,
			MethodDescription: dpm.DpMethod.Description,
			Icon:              dpm.Dp.Icon,
			Used:              int64(dpm.DpMethod.Invoked),
			Labels: &model.LabelsInfos{
				BlockChain: make([]string, 0),
				Category:   make([]string, 0),
				Scenario:   make([]string, 0),
			},
		}
		if dpm.DpMethod.Labels.BlockChain != "" {
			dpRes.Labels.BlockChain = strings.Split(dpm.DpMethod.Labels.BlockChain, ",")
		}
		if dpm.DpMethod.Labels.Category != "" {
			dpRes.Labels.Category = strings.Split(dpm.DpMethod.Labels.Category, ",")
		}
		if dpm.DpMethod.Labels.Scenario != "" {
			dpRes.Labels.Scenario = strings.Split(dpm.DpMethod.Labels.Scenario, ",")
		}
		res = append(res, dpRes)
	}
	return &model.GetAllDataProviderMethodsResp{
		Total: count,
		Data:  res,
	}, nil
}

func (r *queryResolver) GetAllAlgorithmProviderMethod(ctx context.Context, first *int64, skip *int64, where *model.AlgorithmProviderMethodWhere, orderBy *string, orderDirection *string, labels model.LabelsInfo) (*model.GetAllAlgorithmProviderMethodsResp, error) {
	strWhere := ""
	if where != nil {
		if where.NameLike != nil {
			strWhere = " where t1.name like '%" + *where.NameLike + "%' "
		}
		if where.DidLike != nil {
			if len(strWhere) == 0 {
				strWhere = " where t1.did like '%" + *where.DidLike + "%' "
			} else {
				strWhere = strWhere + " and t1.did like '%" + *where.DidLike + "%' "
			}
		}
		strWhere = strWhere + " and t1.status != '" + store.METHOD_REMOVED + "'"
	}
	whereStr := " and"
	for _, blockChainLabels := range labels.BlockChain {
		whereStr += " block_chain_labels like '%" + blockChainLabels + "%' or "
	}
	for _, catgegoryLabels := range labels.Category {
		whereStr += " category_labels like '%" + catgegoryLabels + "%' or "
	}
	for _, catgegoryLabels := range labels.Scenario {
		whereStr += " category_labels like '%" + catgegoryLabels + "%' or "
	}
	whereStr = strings.TrimRight(whereStr, "or ")
	if whereStr == " and" {
		whereStr = " "
	}
	count, err := store.MySqlDB.QueryAllAlgorithmProviderMethodsCountByCondition(strWhere + " " + whereStr)
	if err != nil {
		return nil, err
	}
	if count == int64(0) {
		return &model.GetAllAlgorithmProviderMethodsResp{
			Total: 0,
			Data:  nil,
		}, nil
	}
	strWhere = strWhere + " " + whereStr
	if orderBy != nil {
		strWhere = strWhere + " order by t1." + *orderBy
	}
	if orderDirection != nil {
		if !strings.EqualFold("asc", *orderDirection) && !strings.EqualFold("desc", *orderDirection) {
			return nil, fmt.Errorf("orderDirection only accept 'asc' or 'desc'")
		}
		strWhere = strWhere + " " + *orderDirection
	}
	if first != nil {
		strWhere = strWhere + fmt.Sprintf(" limit %d", *first)
	}
	if skip != nil {
		strWhere = strWhere + fmt.Sprintf(" offset %d", *skip)
	}

	apms, err := store.MySqlDB.QueryAllAlgorithmProviderMethodsByCondition(strWhere)
	if err != nil {
		return nil, err
	}
	res := make([]*model.APMethodWithAp, 0)
	for _, apm := range apms {
		apRes := &model.APMethodWithAp{
			Did:               apm.Ap.Did,
			Name:              apm.Ap.Title,
			Method:            apm.ApMethod.Method,
			MethodName:        apm.ApMethod.Name,
			MethodDescription: apm.ApMethod.Description,
			Icon:              apm.Ap.Icon,
			Used:              int64(apm.ApMethod.Invoked),
			Labels: &model.LabelsInfos{
				BlockChain: make([]string, 0),
				Category:   make([]string, 0),
				Scenario:   make([]string, 0),
			},
		}
		if apm.ApMethod.Labels.BlockChain != "" {
			apRes.Labels.BlockChain = strings.Split(apm.ApMethod.Labels.BlockChain, ",")
		}
		if apm.ApMethod.Labels.Category != "" {
			apRes.Labels.Category = strings.Split(apm.ApMethod.Labels.Category, ",")
		}
		if apm.ApMethod.Labels.Scenario != "" {
			apRes.Labels.Scenario = strings.Split(apm.ApMethod.Labels.Scenario, ",")
		}
		res = append(res, apRes)
	}
	return &model.GetAllAlgorithmProviderMethodsResp{
		Total: count,
		Data:  res,
	}, nil
}

func (r *queryResolver) GetCompositeDpInfo(ctx context.Context, did string, method string) ([]*model.MethodInfo, error) {
	return store.MySqlDB.GetCompositeDpInfo(did, method)
}

func (r *queryResolver) GetAllAPInfo(ctx context.Context) ([]*model.MethodInfo, error) {
	return store.MySqlDB.GetAllApInfo()
}

func (r *queryResolver) GetBasedVotingStrategy(ctx context.Context, addrs []string, space string, snapshot string, network string, options *model.SnapShotOptions) ([]*model.StrategyResult, error) {
	resp := make([]*model.StrategyResult, 0)
	for _, addr := range addrs {
		res, err := store.MySqlDB.QuerySnapShotAssetsScore(addr)
		if err != nil {
			return nil, err
		}
		score, err := strconv.ParseInt(res, 10, 64)
		if err != nil {
			return nil, err
		}
		resp = append(resp, &model.StrategyResult{
			Address: addr,
			Score:   score,
		})
	}
	return resp, nil
}

func (r *queryResolver) GetUserBasicInfo(ctx context.Context, did string) (*model.UserBasicInfo, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return nil, err
	}

	bi, err := store.MySqlDB.GetUserBasicInfo(did)
	if err != nil {
		return nil, err
	}
	if bi != nil {
		res := &model.UserBasicInfo{
			Did:      did,
			NickName: bi.NickName,
			Avatar:   &bi.Avatar,
			Email:    bi.Email,
		}
		return res, nil
	}
	return &model.UserBasicInfo{
		Did:      did,
		NickName: "",
		Avatar:   nil,
		Email:    "",
	}, nil
}

func (r *queryResolver) GetApplicationInfo(ctx context.Context, did string) (*model.ApplicationInfo, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return nil, err
	}

	ai, err := store.MySqlDB.QueryApplicationInfo(did)
	if err != nil {
		return nil, err
	}
	if ai == nil {
		return &model.ApplicationInfo{
			Did:     did,
			Name:    "",
			Website: "",
		}, nil
	}
	return &model.ApplicationInfo{
		Did:     ai.Did,
		Name:    ai.Name,
		Website: ai.Website,
	}, nil
}

func (r *queryResolver) QueryOuterTask(ctx context.Context, input *model.OuterTaskReq) (*model.UserTasks, error) {
	if strings.HasPrefix(input.CallerDid, "did:etho:") {

		if !utils.AddressInArray(input.CallerDid, config.GlobalConfig.CallerAddrs) {
			return nil, fmt.Errorf("unauthorized caller did:%s", input.CallerDid)
		}

		if config.GlobalConfig.SigAuth {
			ethaddr := strings.ReplaceAll(input.CallerDid, "did:etho:", "0x")
			msg, err := json.Marshal(input.Data)
			if err != nil {
				return nil, err
			}

			verify := utils.ETHVerifySig(ethaddr, input.Sig, msg)
			if !verify {
				return nil, fmt.Errorf("verify sig failed")
			}
		}

		task, err := store.MySqlDB.QueryTaskByPK(input.Data.TaskID)
		if err != nil {
			log.Errorf("errors on QueryTaskByPK:%s", err.Error())
			return nil, err
		}
		if task == nil {
			return nil, nil
		}
		filepath := ""
		if len(task.ResultFile) > 0 {
			filepath = config.GlobalConfig.FilePath + task.ResultFile
		}

		return &model.UserTasks{
			TaskID:      fmt.Sprintf("%d", task.TaskId),
			UserDid:     task.UserDID,
			ApDid:       task.ApDID,
			ApName:      task.ApName,
			ApMethod:    task.ApMethod,
			DpDid:       task.DpDID,
			DpName:      task.DpName,
			DpMethod:    task.DpMethod,
			CreateTime:  fmt.Sprintf("%d", task.CreateTime.Unix()),
			UpdateTime:  fmt.Sprintf("%d", task.UpdateTime.Unix()),
			TaskStatus:  common.TransformTaskStatus(task.TaskStatus),
			TaskResult:  &task.TaskResult,
			ResultFile:  &filepath,
			IssueTxhash: &task.IssueTxhash,
		}, nil

	} else {
		return nil, fmt.Errorf("not a supported did:%s", input.CallerDid)
	}
}

func (r *queryResolver) RequestChanllenge(ctx context.Context, input *model.ClientHello) (*model.ServerHello, error) {
	cr := &modules.ClientHello{
		Ver:    input.Ver,
		Type:   input.Type,
		Action: int(input.Action),
	}

	serverHello, err := service.OloginService.OntloginSdk.GenerateChallenge(cr)
	if err != nil {
		log.Errorf("GenerateChallenge failed:%s", err.Error())
		return nil, err
	}

	res := &model.ServerHello{
		Ver:   serverHello.Ver,
		Type:  serverHello.Type,
		Nonce: serverHello.Nonce,
		Server: &model.ServerInfo{
			Name: serverHello.Server.Name,
			Icon: serverHello.Server.Icon,
			URL:  serverHello.Server.Url,
			Did:  serverHello.Server.Did,
		},
		Chain:     serverHello.Chain,
		Alg:       serverHello.Alg,
		VcFilters: make([]*model.VCFilter, 0),
	}

	return res, nil
}

func (r *queryResolver) QueryThirdPartyVc(ctx context.Context, did string, mediaType string) (string, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return "", err
	}
	return store.MySqlDB.QueryThirdPartyVc(did, mediaType)
}

func (r *queryResolver) QueryAllThirdPartyVcStatus(ctx context.Context, did string) ([]*model.ThirdPartyVcStatus, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return nil, err
	}
	return store.MySqlDB.QueryAllThirdPartyVcStatus(did)
}

func (r *queryResolver) QueryUserKycInfo(ctx context.Context, did string) (string, error) {
	if err := auth.CheckLogin(ctx, did); err != nil {
		return "", err
	}
	return store.MySqlDB.QueryUserKycInfo(did)
}

func (r *queryResolver) QueryUserPublishedDp(ctx context.Context, userDid string) (*model.DpInfoRes, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return nil, err
	}
	return store.MySqlDB.QueryDPInfo(userDid, utils.PUBLISHED)
}

func (r *queryResolver) QueryUserLatestDPInfo(ctx context.Context, userDid string) (*model.DpInfoRes, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return nil, err
	}
	return store.MySqlDB.QueryDPInfo(userDid, "")
}

func (r *queryResolver) QueryDPDataSetInfo(ctx context.Context, userDid string, dataSetID int64) (*model.DpDataSetRes, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return nil, err
	}
	return store.MySqlDB.QueryDPDataSetInfo(userDid, dataSetID)
}

func (r *queryResolver) QueryUserDPDataSetList(ctx context.Context, userDid string, dataSetName string, status string, page int64, size int64, labels model.LabelsInfo) (*model.DPDataSetList, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return nil, err
	}
	return store.MySqlDB.QueryUserDPDataSetLists(userDid, dataSetName, status, page, size, labels)
}

func (r *queryResolver) QueryDPLabels(ctx context.Context) (*model.LabelsInfos, error) {
	return store.MySqlDB.GetDataSetLabels(utils.DP_LABELS)
}

func (r *queryResolver) QuerySig(ctx context.Context, addr string, nftType int64, score int64) (*model.SigResp, error) {
	panic("this interface is not executable!")
	hashbts, err := service.GlobalNftClaimService.GetUserClaimHash(addr, int(nftType), uint64(score))
	if err != nil {
		return nil, err
	}
	sig, err := service.GlobalNftClaimService.SignMsg(hashbts)
	if err != nil {
		return nil, err
	}
	return &model.SigResp{
		Hash: hexutil.Encode(hashbts),
		Sig:  hexutil.Encode(sig),
	}, nil
}

func (r *queryResolver) GetClaimNFTRecords(ctx context.Context, first *int64, skip *int64, where *model.ClaimNFTWhere, orderBy *string, orderDirection *string) (*model.ClaimNFTRecordsResp, error) {
	strWhere := ""
	if where != nil {
		strWhere = " where status=3 and "
		and := ""
		if where.Chain != nil {
			strWhere = strWhere + fmt.Sprintf(" chain = '%s'", *where.Chain)
			and = " and "
		}
		if where.UserDid != nil {
			strWhere = strWhere + and + fmt.Sprintf(" user_did = '%s'", *where.UserDid)
			and = " and "
		}
		if where.UserAddress != nil {
			strWhere = strWhere + and + fmt.Sprintf(" user_address = '%s'", *where.UserAddress)
			and = " and "
		}
		if where.NftType != nil {
			strWhere = strWhere + and + fmt.Sprintf(" nft_type = %d", *where.NftType)
			and = " and "
		}
		if where.Result != nil {
			strWhere = strWhere + and + fmt.Sprintf(" result = '%s'", *where.Result)
		}
	}

	count, err := store.MySqlDB.QueryClaimRecordCountByCondition(strWhere)
	if err != nil {
		return nil, err
	}
	if count == int64(0) {
		return &model.ClaimNFTRecordsResp{
			Count:   0,
			Records: make([]*model.ClaimNFTRecord, 0),
		}, nil
	}

	if orderBy != nil {
		strWhere = strWhere + " order by " + *orderBy
	}
	if orderDirection != nil {
		if !strings.EqualFold("asc", *orderDirection) && !strings.EqualFold("desc", *orderDirection) {
			return nil, fmt.Errorf("orderDirection only accept 'asc' or 'desc'")
		}
		strWhere = strWhere + " " + *orderDirection + ", create_time asc"
	}
	if first != nil {
		strWhere = strWhere + fmt.Sprintf(" limit %d", *first)
	}
	if skip != nil {
		strWhere = strWhere + fmt.Sprintf(" offset %d", *skip)
	}

	list, err := store.MySqlDB.QueryClaimRecordsByCondition(strWhere)
	if err != nil {
		return nil, err
	}
	records := make([]*model.ClaimNFTRecord, 0)
	for _, l := range list {
		r := &model.ClaimNFTRecord{
			TxHash:          l.TxHash,
			Chain:           l.Chain,
			ContractAddress: l.ContractAddress,
			NftType:         l.NftType,
			UserDid:         l.UserDID,
			UserAddress:     l.UserAddress,
			CreateTime:      l.CreateTime.Unix(),
			Result:          l.Result,
			Score:           fmt.Sprintf("%d", l.Score),
		}
		records = append(records, r)
	}
	return &model.ClaimNFTRecordsResp{
		Count:   count,
		Records: records,
	}, nil
}

func (r *queryResolver) GetNFTSettings(ctx context.Context, first *int64, skip *int64, where *model.NFTSettingWhere, orderBy *string, orderDirection *string) (*model.NFTSettingResp, error) {
	strWhere := ""
	if where != nil {
		strWhere = " where "
		if where.ID != nil {
			strWhere = strWhere + fmt.Sprintf(" id = %d", *where.ID)
		}
	}

	count, err := store.MySqlDB.GetNFTSettingCountByCondition(strWhere)
	if err != nil {
		return nil, err
	}
	if count == int64(0) {
		return &model.NFTSettingResp{
			Count:   0,
			Records: make([]*model.NFTSetting, 0),
		}, nil
	}

	if orderBy != nil {
		strWhere = strWhere + " order by " + *orderBy
	}
	if orderDirection != nil {
		if !strings.EqualFold("asc", *orderDirection) && !strings.EqualFold("desc", *orderDirection) {
			return nil, fmt.Errorf("orderDirection only accept 'asc' or 'desc'")
		}
		strWhere = strWhere + " " + *orderDirection
	}
	if first != nil {
		strWhere = strWhere + fmt.Sprintf(" limit %d", *first)
	}
	if skip != nil {
		strWhere = strWhere + fmt.Sprintf(" offset %d", *skip)
	}

	list, err := store.MySqlDB.GetNFTSettingByCondition(strWhere)
	if err != nil {
		return nil, err
	}
	res := make([]*model.NFTSetting, 0)

	for _, l := range list {
		t := &model.NFTSetting{}
		t.ID = int64(l.Id)
		t.Name = l.Name
		t.Description = l.Description
		//t.Image = l.Image
		t.Image = l.AltImage
		t.DpDid = l.DpDID
		t.DpMethod = l.DpMethod
		t.ApDid = l.ApDID
		t.ApMethod = l.ApMethod
		t.LowestScore = int64(l.LowestScore)
		t.ValidDays = int64(l.ValidDays)
		t.Restriction = l.Restriction
		t.IssueBy = l.IssueBy
		chainarrs := make([]*model.ChainAddress, 0)
		for k, v := range config.GlobalConfig.NFTConfig.NftInfos {
			ca := &model.ChainAddress{
				Chain:           k,
				ContractAddress: v.ContractAddress,
			}
			chainarrs = append(chainarrs, ca)
		}
		t.ChainAddresses = chainarrs

		apinfo, err := store.MySqlDB.QueryAlgorithmProviderByDid(l.ApDID)
		if err != nil {
			return nil, err
		}
		if apinfo == nil {
			return nil, fmt.Errorf("no ap info found")
		}
		t.ApName = apinfo.Title

		apminfo, err := store.MySqlDB.QueryAPMethodByDIDAndMethod(l.ApDID, l.ApMethod)
		if err != nil {
			return nil, err
		}
		if apminfo == nil {
			return nil, fmt.Errorf("no ap info found")
		}
		t.ApMethodName = apminfo.Name
		res = append(res, t)
	}
	return &model.NFTSettingResp{
		Count:   count,
		Records: res,
	}, nil
}

func (r *queryResolver) GetUserClaimedNft(ctx context.Context, first *int64, skip *int64, where *model.UserClaimedNFTWhere, orderBy *string, orderDirection *string) (*model.UserClaimedNFTResp, error) {
	//todo add auth check
	strWhere := ""
	if where != nil {
		strWhere = fmt.Sprintf(" where status = %d and result='SUCCEED'", store.CLAIM_NFT_STATUS_DONE)
		and := " and "
		if where.UserDid != nil {
			strWhere = strWhere + and + fmt.Sprintf(" user_did = '%s'", *where.UserDid)
		}
		if where.Address != nil {
			strWhere = strWhere + and + fmt.Sprintf(" user_address = '%s'", *where.Address)
		}
		if where.NftType != nil {
			strWhere = strWhere + and + fmt.Sprintf(" nft_type = %d", *where.NftType)
		}
		if where.TokenID != nil {
			strWhere = strWhere + and + fmt.Sprintf(" token_id = %d", *where.TokenID)
		}
		if where.Chain != nil {
			strWhere = strWhere + and + fmt.Sprintf(" chain = '%s'", *where.Chain)
		}
	}

	count, err := store.MySqlDB.QueryClaimRecordCountByCondition(strWhere)
	if err != nil {
		return nil, err
	}
	if count == int64(0) {
		return &model.UserClaimedNFTResp{
			Count:   0,
			Records: make([]*model.UserClaimedNft, 0),
		}, nil
	}

	if orderBy != nil {
		strWhere = strWhere + " order by " + *orderBy
	}
	if orderDirection != nil {
		if !strings.EqualFold("asc", *orderDirection) && !strings.EqualFold("desc", *orderDirection) {
			return nil, fmt.Errorf("orderDirection only accept 'asc' or 'desc'")
		}
		strWhere = strWhere + " " + *orderDirection
	}
	if first != nil {
		strWhere = strWhere + fmt.Sprintf(" limit %d", *first)
	}
	if skip != nil {
		strWhere = strWhere + fmt.Sprintf(" offset %d", *skip)
	}

	list, err := store.MySqlDB.QueryClaimRecordsByCondition(strWhere)
	if err != nil {
		return nil, err
	}

	data := make([]*model.UserClaimedNft, 0)

	for _, l := range list {
		t := &model.UserClaimedNft{}
		detail, err := service.GlobalNftClaimService.GetNFTDetail(l.Chain, l.ContractAddress, l.TokenId)
		if err != nil {
			return nil, err
		}
		nftsettings, err := store.MySqlDB.GetNFTSettingByCondition(fmt.Sprintf("where id = %d", l.NftType))
		if err != nil {
			return nil, err
		}
		if len(nftsettings) != 1 {
			return nil, fmt.Errorf("nft type:%d is not exist", l.NftType)
		}
		nftsetting := nftsettings[0]

		if !strings.EqualFold(l.UserAddress, detail.OriginOwner.Hex()) {
			return nil, fmt.Errorf("tokenid %d is not belong to user:%s", l.TokenId, l.UserAddress)
		}
		dp, err := store.MySqlDB.QueryDataProviderByDid(nftsetting.DpDID)
		if err != nil {
			return nil, err
		}
		if dp == nil {
			return nil, fmt.Errorf("no dp with did:%s\n", nftsetting.DpDID)
		}

		ap, err := store.MySqlDB.QueryAlgorithmProviderByDid(nftsetting.ApDID)
		if err != nil {
			return nil, err
		}
		if ap == nil {
			return nil, fmt.Errorf("no ap with did:%s\n", nftsetting.ApDID)
		}

		dpm, err := store.MySqlDB.QueryDPMethodByDIDAndMethod(nftsetting.DpDID, nftsetting.DpMethod)
		if err != nil {
			return nil, err
		}
		if dpm == nil {
			return nil, fmt.Errorf("no dp method with method:%s\n", nftsetting.DpMethod)
		}

		apm, err := store.MySqlDB.QueryAPMethodByDIDAndMethod(nftsetting.ApDID, nftsetting.ApMethod)
		if err != nil {
			return nil, err
		}
		if apm == nil {
			return nil, fmt.Errorf("no ap method with method:%s\n", nftsetting.ApMethod)
		}

		t.Owner = l.UserAddress
		t.TokenID = l.TokenId
		t.DpMethod = detail.Category.Dpmethod
		t.ApMethod = detail.Category.Apmethod
		t.ApDid = detail.Category.Apdid
		t.DpDid = detail.Category.Dpdid
		//t.DpTitle = detail.Category.DpTitle
		t.DpTitle = dp.Title
		//t.ApTitle = detail.Category.ApTitle
		t.ApTitle = ap.Title
		//t.ApMethodTitle = detail.Category.ApmethodTitle
		//t.DpMethodTitle = detail.Category.DpmethodTitle
		t.ApMethodTitle = apm.Name
		t.DpMethodTitle = dpm.Name

		t.ValidDays = detail.Category.ValidDays.Int64()
		t.ValidTo = detail.ValidTo.Int64()
		//t.Image = detail.Category.Image
		t.Image = nftsetting.AltImage
		t.TxHash = l.TxHash
		t.IssueBy = nftsetting.IssueBy
		t.Name = nftsetting.Name
		t.Chain = l.Chain
		t.LowestScore = int64(nftsetting.LowestScore)
		t.Description = nftsetting.Description
		t.ClaimTime = detail.ValidTo.Int64() - detail.Category.ValidDays.Int64()*60*60*24
		t.ContractAddress = l.ContractAddress
		t.NftType = l.NftType
		t.NftScore = detail.Score.Int64()
		t.IsExpired = time.Now().Unix() >= t.ValidTo
		data = append(data, t)
	}
	return &model.UserClaimedNFTResp{
		Count:   count,
		Records: data,
	}, nil
}

func (r *queryResolver) GetNFTClaimedCount(ctx context.Context, nftType int64, userDid string) (*model.NFTClaimedCountResp, error) {
	//todo auth check

	totalCount, err := store.MySqlDB.QueryClaimRecordCountByCondition(fmt.Sprintf(" where nft_type=%d and status=%d", nftType, store.CLAIM_NFT_STATUS_DONE))
	if err != nil {
		return nil, err
	}
	userCount, err := store.MySqlDB.QueryClaimRecordCountByCondition(fmt.Sprintf(" where user_did='%s' and nft_type=%d and status=%d", userDid, nftType, store.CLAIM_NFT_STATUS_DONE))
	if err != nil {
		return nil, err
	}

	return &model.NFTClaimedCountResp{
		TotalCount: totalCount,
		UserCount:  userCount,
	}, nil
}

func (r *queryResolver) QueryUserPublishedAp(ctx context.Context, userDid string) (*model.ApInfoRes, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return nil, err
	}
	return store.MySqlDB.QueryAPInfo(userDid, utils.PUBLISHED)
}

func (r *queryResolver) QueryUserLatestAPInfo(ctx context.Context, userDid string) (*model.ApInfoRes, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return nil, err
	}
	return store.MySqlDB.QueryAPInfo(userDid, "")
}

func (r *queryResolver) QueryAPDataSetInfo(ctx context.Context, userDid string, dataSetID int64) (*model.ApDataSetRes, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return nil, err
	}
	return store.MySqlDB.QueryAPDataSetInfo(userDid, dataSetID)
}

func (r *queryResolver) QueryUserAPDataSetList(ctx context.Context, userDid string, dataSetName string, status string, page int64, size int64, labels model.LabelsInfo) (*model.APDataSetList, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return nil, err
	}
	return store.MySqlDB.QueryUserAPDataSetLists(userDid, dataSetName, status, page, size, labels)
}

func (r *queryResolver) QueryAPLabels(ctx context.Context) (*model.LabelsInfos, error) {
	return store.MySqlDB.GetDataSetLabels(utils.AP_LABELS)
}

func (r *queryResolver) QueryUserSNSBinding(ctx context.Context, callerDid string, address string, encrypt bool) (*model.SNSBindingResp, error) {
	//todo add authentication check

	//did from address
	userdid := ""
	if strings.HasPrefix(address, "0x") {
		userdid = strings.ReplaceAll(address, "0x", "did:etho:")
	} else {
		userdid = "did:ont:" + address
	}

	rs, err := store.MySqlDB.QueryAllThirdPartyVcStatus(userdid)
	if err != nil {
		return nil, err
	}
	data := &model.SNSBindingData{}

	for _, r := range rs {
		if r.MediaType == "Github" {
			data.Github = true
		}
		if r.MediaType == "Linkedin" {
			data.Linkedin = true
		}
		if r.MediaType == "Twitter" {
			data.Tweeter = true
		}
		if r.MediaType == "Facebook" {
			data.Facebook = true
		}
		if r.MediaType == "ShuftiPro" {
			data.ShuftiPro = true
		}
		if r.MediaType == "BrightID" {
			data.BrightID = true
		}
	}

	dataToSign, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	sig, err := service.SysDidService.SignData(config.GlobalConfig.DidConf[0].DID, dataToSign)
	if err != nil {
		return nil, err
	}
	dataWithSig := &model.SNSBindingDataWithSig{
		Data: data,
		Sig:  hex.EncodeToString(sig),
	}

	if encrypt {
		databytes, err := json.Marshal(dataWithSig)
		if err != nil {
			return nil, err
		}
		encrypted, err := service.SysDidService.EncryptDataWithDID(databytes, callerDid)
		if err != nil {
			return nil, err
		}
		enhex := hex.EncodeToString(encrypted)

		return &model.SNSBindingResp{
			Data:      nil,
			Encrypted: &enhex,
		}, nil

	} else {
		return &model.SNSBindingResp{
			Data:      dataWithSig,
			Encrypted: nil,
		}, nil
	}
}

func (r *queryResolver) QueryUserBasicInfo(ctx context.Context, userDid string) (*model.UserBasicInfoResp, error) {
	if err := auth.CheckLogin(ctx, userDid); err != nil {
		return nil, err
	}

	addrCnt, err := store.MySqlDB.GetUserAddressInfoCount(userDid)
	if err != nil {
		return nil, err
	}
	tpVCCnt, err := store.MySqlDB.GetThirdPartyVCCounts(userDid)
	if err != nil {
		return nil, err
	}
	repVCCnt, err := store.MySqlDB.GetUserReputationCount(userDid)
	if err != nil {
		return nil, err
	}
	nftCnt, err := store.MySqlDB.QueryClaimRecordCountByCondition("where user_did = '" + userDid + "' ")
	if err != nil {
		return nil, err
	}
	apmCnt, err := store.MySqlDB.QueryAPDatatSetCountByCondition(userDid)
	if err != nil {
		return nil, err
	}
	dpmCnt, err := store.MySqlDB.QueryDPDatatSetCountByCondition(userDid)
	if err != nil {
		return nil, err
	}

	return &model.UserBasicInfoResp{
		WalletAddress:    addrCnt,
		Verifications:    tpVCCnt,
		Credentials:      repVCCnt,
		Nfts:             nftCnt,
		ModelPublished:   apmCnt,
		DatasetPublished: dpmCnt,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) SingleUpload(ctx context.Context, file graphql.Upload) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *queryResolver) GetStrategy(ctx context.Context, addrs []string, space string) ([]*model.StrategyResult, error) {
	resp := make([]*model.StrategyResult, 0)
	//todo add for fixed
	//dpdid := ""
	//dpmethod := ""
	//apdid := ""
	//apmethod := ""

	for _, addr := range addrs {
		tmp := &model.StrategyResult{
			Address: addr,
			Score:   0,
		}
		//for test
		score := rand.Int31n(1000)
		tmp.Score = int64(score)
		resp = append(resp, tmp)
		//userDID := strings.ReplaceAll(addr, "0x", "did:etho:")
		//task, err := store.MySqlDB.QueryTaskByUniqueKey(userDID, dpdid, dpmethod, apdid, apmethod)
		//if err != nil {
		//	resp = append(resp, tmp)
		//	continue
		//}
		//if task == nil {
		//	resp = append(resp, tmp)
		//	continue
		//}
		//if task.TaskStatus != store.TASK_STATUS_DONE {
		//	resp = append(resp, tmp)
		//	continue
		//}
		//
		//t, err := strconv.ParseInt(task.TaskResult, 10, 64)
		//if err != nil {
		//	resp = append(resp, tmp)
		//	continue
		//}
		//tmp.Score = t
		//resp = append(resp, tmp)
	}
	return resp, nil
}
func (r *mutationResolver) AddTaskTest(ctx context.Context, input model.AddTask, overwrite bool) (int64, error) {
	//panic("not implemented")
	taskid, err := store.MySqlDB.AddTask(input.UserDid, input.ApDid, input.ApMethod, input.DpDid, input.DpMethod, input.UserDid, "")
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			if !overwrite {
				return -2, fmt.Errorf("same task with ap & dp already existed, please wait for complete and remove the task")
			} else {
				err = store.MySqlDB.DeleteTask(input.UserDid, input.ApDid, input.ApMethod, input.DpDid, input.DpMethod)
				if err != nil {
					return -1, err
				} else {
					taskid, err := store.MySqlDB.AddTask(input.UserDid, input.ApDid, input.ApMethod, input.DpDid, input.DpMethod, input.UserDid, "")
					if err != nil {
						return -1, err
					}
					return taskid, nil
				}
			}
		}
		return -1, fmt.Errorf("DB error")
	}
	return taskid, nil
}
