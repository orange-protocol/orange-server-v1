package service

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/smartcontract"
	"github.com/orange-protocol/orange-server-v1/store"
)

var GlobalNftClaimService *NftClaimService

const ClaimBatchCount = 100

type NftClaimService struct {
	ks                *keystore.KeyStore
	pwd               string
	clients           map[string]*ethclient.Client
	exitChan          chan int
	theGraphService   *TheGraphService
	nftTimeoutMinutes time.Duration
}

func InitNftClaimService(cfg *config.SysConfig) {
	GlobalNftClaimService = NewNftClaimService(cfg)
	go GlobalNftClaimService.MonitorTx()
	go GlobalNftClaimService.DealWithTimeoutTx()
}

func NewNftClaimService(cfg *config.SysConfig) *NftClaimService {
	service := &NftClaimService{}
	capitalKeyStore := keystore.NewKeyStore(cfg.ETHWallet.KeyStore, keystore.StandardScryptN,
		keystore.StandardScryptP)

	accArr := capitalKeyStore.Accounts()
	if len(accArr) == 0 {
		log.Fatal("eth wallet has no account")
		acct, err := capitalKeyStore.NewAccount(cfg.ETHWallet.Password)
		fmt.Printf("addr is %s\n", acct.Address.String())
		if err != nil {
			panic(err)
		}

	}
	str := ""
	for i, v := range accArr {
		str += fmt.Sprintf("(no.%d acc: %s), ", i+1, v.Address.String())
	}
	log.Infof("server are using accounts: [ %s ]", str)

	service.ks = capitalKeyStore
	service.pwd = cfg.ETHWallet.Password

	if cfg.NFTConfig != nil {
		m := make(map[string]*ethclient.Client)
		for k, v := range cfg.NFTConfig.NftInfos {
			c, err := ethclient.Dial(v.Rpc)
			if err != nil {
				panic(err)
			}
			m[k] = c
		}
		service.clients = m
	}
	service.theGraphService = NewTheGraphService(cfg.GraphConfig)
	service.nftTimeoutMinutes = time.Duration(cfg.NFTTimeOutMinutes) * time.Minute
	return service
}

func (ns *NftClaimService) SignMsg(msghash []byte) ([]byte, error) {
	sig, err := ns.ks.SignHashWithPassphrase(ns.ks.Accounts()[0], ns.pwd, msghash)
	if err != nil {
		return nil, err
	}
	if len(sig) != 65 {
		return nil, fmt.Errorf("sig length is not 65")
	}
	sig[64] += 27
	return sig, nil
}

func (ns *NftClaimService) GetUserClaimHash(userAddr string, nftType int, score uint64) ([]byte, error) {

	nfthash := crypto.Keccak256Hash(
		common.HexToAddress(userAddr).Bytes(),
		common.LeftPadBytes(big.NewInt(int64(nftType)).Bytes(), 32),
		common.LeftPadBytes(big.NewInt(int64(score)).Bytes(), 32),
	)

	hash := crypto.Keccak256Hash(
		[]byte("\x19Ethereum Signed Message:\n32"),
		nfthash.Bytes(),
	)

	return hash.Bytes(), nil
}

func (ns *NftClaimService) MonitorTx() {
	ticker := time.NewTicker(TASK_TICKER)
	for {
		select {
		case <-ticker.C:
			go ns.monitorClaimTxStatus()
			go ns.resolveTimeoutTx()
		case <-ns.exitChan:
			return
		}
	}
}

func (ns *NftClaimService) DealWithTimeoutTx() {
	ticker := time.NewTicker(ns.nftTimeoutMinutes)
	for {
		select {
		case <-ticker.C:
			go ns.DealWithTimeoutNFT()
		case <-ns.exitChan:
			return
		}
	}
}

func (ns *NftClaimService) resolveTimeoutTx() {
	timeTasks, err := store.MySqlDB.QueryClaimRecordsByCondition(fmt.Sprintf(" where status = %d and sysdate() - create_time > %d ", store.CLAIM_NFT_STATUS_RESOLVING, 600))
	if err != nil {
		log.Errorf("errors on QueryClaimRecordsByCondition:%s", err.Error())
		return
	}

	for _, task := range timeTasks {
		_, err = store.MySqlDB.UpdateClaimRecordStatusByPK(store.CLAIM_NFT_STATUS_RESOLVING, store.CLAIM_NFT_STATUS_INIT, task.TxHash, task.Chain)
		if err != nil {
			log.Errorf("errors on UpdateClaimRecordStatusByPK:%s", err.Error())
		}
	}
}

func (ns *NftClaimService) getClaimTokenId(chain, contractAddress string, blockNum uint64, userAddress string) (int64, error) {
	contractAddr := common.HexToAddress(contractAddress)
	nftabi, err := smartcontract.NewSmartcontract(contractAddr, ns.clients[chain])
	if err != nil {
		log.Errorf("errors on NewSmartcontract:%s", err.Error())
		return 0, err
	}
	ite, err := nftabi.FilterTransfer(&bind.FilterOpts{
		Start: blockNum,
		End:   &blockNum,
	}, nil, []common.Address{common.HexToAddress(userAddress)}, nil)
	if err != nil {
		log.Errorf("errors on FilterTransfer:%s", err.Error())
		return 0, err
	}
	tokenid := int64(0)
	if ite.Next() {
		tokenid = ite.Event.TokenId.Int64()
	}
	//test

	return tokenid, err
}

func (ns *NftClaimService) monitorClaimTxStatus() {
	cnt, err := store.MySqlDB.QueryClaimRecordCountByStatus(store.CLAIM_NFT_STATUS_QUERYING)
	if err != nil {
		log.Errorf("errors on QueryClaimRecordCountByStatus:%s", err.Error())
		return
	}

	if cnt >= ClaimBatchCount {
		return
	}

	_, err = store.MySqlDB.UpdateClaimRecordStatusByLimit(store.CLAIM_NFT_STATUS_INIT, store.CLAIM_NFT_STATUS_QUERYING, int(ClaimBatchCount-cnt))
	if err != nil {
		log.Errorf("errors on UpdateClaimRecordStatusByLimit:%s", err.Error())
		return
	}

	list, err := store.MySqlDB.QueryClaimRecordsByStatus(store.CLAIM_NFT_STATUS_QUERYING)
	if err != nil {
		log.Errorf("errors on QueryClaimRecordsByStatus:%s", err.Error())
		return
	}

	for _, record := range list {
		_, err := store.MySqlDB.UpdateClaimRecordStatusByPK(store.CLAIM_NFT_STATUS_QUERYING, store.CLAIM_NFT_STATUS_RESOLVING, record.TxHash, record.Chain)
		if err != nil {
			log.Errorf("errors on UpdateClaimRecordStatusByPK:%s", err.Error())
			continue //skip this record
		}

		c, ok := ns.clients[record.Chain]
		if !ok {
			log.Error("no client for chain:%s", record.Chain)
			continue
		}

		hash := common.HexToHash(record.TxHash)
		_, isPending, err := c.TransactionByHash(context.Background(), hash)
		if err != nil {
			log.Errorf("errors on TransactionByHash:%s,hash:%s", err.Error(), record.TxHash)
			_, err = store.MySqlDB.UpdateClaimRecordStatusByPK(store.CLAIM_NFT_STATUS_RESOLVING, store.CLAIM_NFT_STATUS_QUERYING, record.TxHash, record.Chain)
			if err != nil {
				log.Errorf("errors on UpdateClaimRecordStatusByPK:%s", err.Error())
			}
			continue
		}

		if isPending {
			_, err = store.MySqlDB.SetClaimRecordResultByPK(store.CLAIM_NFT_STATUS_RESOLVING, store.CLAIM_NFT_STATUS_QUERYING, record.TxHash, record.Chain, "PENDING", 0)
			if err != nil {
				log.Error("errors on SetClaimRecordResultByPK:%s", err.Error())
			}
			continue
		}

		reciept, err := c.TransactionReceipt(context.Background(), hash)
		if err != nil {
			log.Error("errors on TransactionReceipt:%s", err.Error())
			_, err = store.MySqlDB.UpdateClaimRecordStatusByPK(store.CLAIM_NFT_STATUS_RESOLVING, store.CLAIM_NFT_STATUS_QUERYING, record.TxHash, record.Chain)
			if err != nil {
				log.Errorf("errors on UpdateClaimRecordStatusByPK:%s", err.Error())
			}
			continue
		}
		result := "SUCCEED"
		tokenId := int64(0)
		if reciept.Status == 0 {
			//failed
			result = "FAILED"
		} else {
			tokenId, err = ns.getClaimTokenId(record.Chain, record.ContractAddress, reciept.BlockNumber.Uint64(), record.UserAddress)
			if err != nil {
				log.Error("errors on getClaimTokenId:%s", err.Error())
				continue
			}
		}

		_, err = store.MySqlDB.SetClaimRecordResultByPK(store.CLAIM_NFT_STATUS_RESOLVING, store.CLAIM_NFT_STATUS_DONE, record.TxHash, record.Chain, result, tokenId)
		if err != nil {
			log.Error("errors on SetClaimRecordResultByPK:%s", err.Error())
		}
	}

}

func (ns *NftClaimService) GetNFTDetail(chain, contractAddress string, tokenid int64) (smartcontract.OrangeReputationtokenDetail, error) {
	contractAddr := common.HexToAddress(contractAddress)
	nftabi, err := smartcontract.NewSmartcontract(contractAddr, ns.clients[chain])
	if err != nil {
		log.Errorf("errors on NewSmartcontract:%s", err.Error())
		return smartcontract.OrangeReputationtokenDetail{}, err
	}

	return nftabi.TokenProperty(nil, big.NewInt(tokenid))
}

func (ns *NftClaimService) DealWithTimeoutNFT() {
	unDoneTask, err := ns.GetNFTTimeoutRecord()
	if err != nil {
		log.Errorf("DealWithTimeoutNFT GetNFTTimeoutRecord err:%s", err)
		return
	}

	err = ns.UpdateClaimNftRecord(unDoneTask)
	if err != nil {
		log.Errorf("DealWithTimeoutNFT UpdateClaimNftRecord err:%s", err)
		return
	}
}

func (ns *NftClaimService) GetNFTTimeoutRecord() ([]*store.ClaimNFTRecord, error) {
	unDoneTask, err := store.MySqlDB.QueryClaimRecordsByCondition(fmt.Sprintf(" where status != %d and sysdate() - create_time > %d ", store.CLAIM_NFT_STATUS_DONE, 600))
	if err != nil {
		log.Errorf("errors on QueryClaimRecordsByCondition:%s", err.Error())
		return nil, err
	}
	return unDoneTask, nil
}

func (ns *NftClaimService) UpdateClaimNftRecord(nftRecords []*store.ClaimNFTRecord) error {
	for _, nftRecord := range nftRecords {
		tokens, err := ns.theGraphService.QueryHashByAddressFromChain(nftRecord.UserAddress, nftRecord.Chain)
		if err != nil {
			return err
		}
		for _, token := range tokens {
			claimNFTRecords, err := store.MySqlDB.QueryClaimRecordsByCondition(fmt.Sprintf("where status = %d and tx_hash = '%s'", store.CLAIM_NFT_STATUS_DONE, token.MintTx))
			if err != nil {
				log.Errorf("errors on UpdateClaimNftRecord:%s", err.Error())
				return err
			}
			if len(claimNFTRecords) == 1 {
				continue
			}
			tokenId, _ := strconv.ParseInt(token.Id, 10, 64)
			tokenDetail, err := ns.GetNFTDetail(nftRecord.Chain, nftRecord.ContractAddress, tokenId)
			if err != nil {
				return err
			}
			if tokenDetail.Category.Apdid == "" {
				return fmt.Errorf("UpdateClaimNftRecord getTokenDetail is null")
			}
			nftSettings, err := store.MySqlDB.GetNFTSettingByCondition(fmt.Sprintf(" where dp_did='%s' and dp_method='%s' and ap_did='%s' and ap_method='%s' ", tokenDetail.Category.Dpdid,
				tokenDetail.Category.Dpmethod, tokenDetail.Category.Apdid, tokenDetail.Category.Apmethod))
			if err != nil {
				log.Errorf("erros on GetNFTSettingByCondition:%s", err.Error())
				return err
			}
			for _, nftsetting := range nftSettings {
				if nftsetting.Id == int(nftRecord.NftType) {
					_, err = store.MySqlDB.UpdateClaimRecordResult(nftRecord.Status, store.CLAIM_NFT_STATUS_DONE, token.MintTx, nftRecord.Chain, "SUCCEED", nftRecord.NftType, tokenId, nftRecord.TxHash)
					if err != nil {
						log.Errorf("UpdateClaimNftRecord UpdateClaimRecordResult:%s", err.Error())
						return fmt.Errorf("UpdateClaimNftRecord UpdateClaimRecordResult:%s", err.Error())
					}
					log.Infof("UpdateClaimNftRecord UpdateClaimRecordResult hash:%s,tokenId:%d,nftType:%d,chain:%s", token.MintTx, tokenId, nftRecord.NftType, nftRecord.Chain)
				}
			}
		}
	}
	return nil
}
