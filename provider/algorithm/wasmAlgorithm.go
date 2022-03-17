package algorithm

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/provider/common"
	"github.com/orange-protocol/orange-server-v1/store"
	"github.com/orange-protocol/orange-server-v1/utils"
	"github.com/orange-protocol/orange-server-v1/wasm"
)

const (
	GAS_LIMIT  = 10000000
	STEP_LIMIT = 10000000
)

type WasmAlgorithmProvider struct {
	did            string
	entris         map[string]*common.Method
	excuter        *wasm.Executor
	executorDID    string
	executorPrikey *ecdsa.PrivateKey
}

func NewWasmAlgorithmProvider(ap *store.AlgorithmProvider) (*WasmAlgorithmProvider, error) {
	methods, err := store.MySqlDB.QueryAlgorithmProviderMethodByDid(ap.Did)
	if err != nil {
		log.Errorf("errors on QueryAlgorithmProviderMethodByDid :%s, error:%s", ap.Did, err.Error())
		return nil, err
	}

	entries := make(map[string]*common.Method)
	for _, method := range methods {
		entries[method.Method] = &common.Method{
			Url:    method.URL,
			Param:  method.Param,
			Result: method.ResultSchema,
		}
	}

	execdid := config.GlobalConfig.WasmExecutor.Did
	chain, err := utils.GetChainFromDID(execdid)
	if chain != "ont" {
		panic("only ont did supported now")
	}
	sdk := ontology_go_sdk.NewOntologySdk()
	account, err := utils.OpenAccount(config.GlobalConfig.WasmExecutor.Wallet, config.GlobalConfig.WasmExecutor.Password, config.GlobalConfig.WasmExecutor.Address, sdk)
	if err != nil {
		panic(err)
	}

	ecdsaPrivkey, err := utils.PrivateKeyToEcdsaPrivkey(keypair.SerializePrivateKey(account.PrivateKey))
	if err != nil {
		panic(err)
	}

	executor := wasm.NewExecutor(GAS_LIMIT, STEP_LIMIT, &wasm.ExecEnv{})
	return &WasmAlgorithmProvider{
		did:            ap.Did,
		entris:         entries,
		excuter:        executor,
		executorDID:    config.GlobalConfig.WasmExecutor.Did,
		executorPrikey: ecdsaPrivkey,
	}, nil
}

func (this *WasmAlgorithmProvider) Invoke(methodName string, paramMap map[string]interface{}) (interface{}, error) {

	method, ok := this.entris[methodName]
	if !ok {
		return nil, fmt.Errorf("no method with %s", methodName)
	}
	addr, err := getCodeAddressFromUrl(method.Url)
	if err != nil {
		return nil, err
	}

	//todo decrypt param
	encrypted := paramMap["%input"]
	if encrypted != nil {
		//deal with composite
		tmpstr := encrypted.(string)
		if strings.Contains(tmpstr, ";;") {
			m := make(map[string]map[string]interface{})
			tmparr := strings.Split(tmpstr, ";;")
			for _, t := range tmparr {
				s := strings.Split(t, "::")
				bts, err := hex.DecodeString(s[1])
				if err != nil {
					log.Errorf("decrypt input failed:%s\n", err.Error())
					return nil, err
				}
				decrypt, err := utils.DecryptMsg(this.executorPrikey, bts)
				if err != nil {
					log.Errorf("decrypt input failed:%s\n", err.Error())
					return nil, err
				}
				tmpMap := make(map[string]interface{})
				err = json.Unmarshal(decrypt, &tmpMap)
				if err != nil {
					return nil, err
				}
				m[s[0]] = tmpMap
			}
			bytes, err := json.Marshal(m)
			if err != nil {
				return nil, err
			}
			log.Debugf("===Invoke decrypt:%s\n", bytes)
			paramMap["%input"] = string(bytes)

		} else {
			bts, err := hex.DecodeString(tmpstr)
			if err != nil {
				log.Errorf("decrypt input failed:%s\n", err.Error())
				return nil, err
			}
			decrypt, err := utils.DecryptMsg(this.executorPrikey, bts)
			if err != nil {
				log.Errorf("decrypt input failed:%s\n", err.Error())
				return nil, err
			}
			log.Debugf("===Invoke decrypt:%s\n", decrypt)
			paramMap["%input"] = string(decrypt)
		}
	}

	wasmcode, err := store.MySqlDB.QueryWasmCodeByDIDAndAddress(this.did, addr)
	if err != nil {
		return nil, err
	}
	if wasmcode == nil {
		return nil, fmt.Errorf("no wasm code found for did:%s", this.did)
	}

	wasmcodebytes, err := hex.DecodeString(wasmcode.Code)
	if err != nil {
		return nil, err
	}
	input := common.ProcessParamMap(method, paramMap)

	return this.excuter.Invoke([]byte(input), wasmcodebytes)
}

func getCodeAddressFromUrl(url string) (string, error) {
	//todo url should be : wasm://address format
	if !strings.HasPrefix(url, "wasm://") {
		return "", fmt.Errorf("not a wasm url:%s", url)
	}
	tmp := strings.Split(url, "://")
	if len(tmp) != 2 {
		return "", fmt.Errorf("format error wasm url:%s", url)
	}
	return tmp[1], nil
}

func (this *WasmAlgorithmProvider) VerifySig(body []byte) (bool, error) {
	//todo
	return true, nil
}
