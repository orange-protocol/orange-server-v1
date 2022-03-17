package provider

import (
	"fmt"
	"time"

	"github.com/orange-protocol/orange-server-v1/cache"
	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/provider/algorithm"
	"github.com/orange-protocol/orange-server-v1/provider/common"
	"github.com/orange-protocol/orange-server-v1/provider/data"
	"github.com/orange-protocol/orange-server-v1/store"
)

type AlgorithmProviderInf interface {
	Invoke(methodName string, paramMap map[string]interface{}) (interface{}, error)
	VerifySig(body []byte) (bool, error)
}

func GetAlgorithmProvider(did string) (AlgorithmProviderInf, error) {

	ap, err := store.MySqlDB.QueryAlgorithmProviderByDid(did)
	if err != nil {
		log.Errorf("errors on QueryAlgorithmProviderByDid :%d,error:%s", did, err.Error())
		return nil, err
	}
	if ap == nil {
		return nil, fmt.Errorf("no algorithrm with did:%s", did)
	}

	if ap.APType == common.AP_TYPE_OUTER {
		res, err := algorithm.NewHttpAlgorithmProvider(ap)
		if err != nil {
			return nil, err
		}
		//cache.GlobalCache.Add(key, 30*time.Minute, res)
		return res, nil
	} else {
		res, err := algorithm.NewWasmAlgorithmProvider(ap)
		if err != nil {
			return nil, err
		}
		//cache.GlobalCache.Add(key, 30*time.Minute, res)
		return res, nil
	}
}

func GetDataProvider(did string) (DataProviderInf, error) {

	key := fmt.Sprintf("DP-%s", did)
	cap, err := cache.GlobalCache.Value(key)
	if err == nil {
		return cap.Data().(DataProviderInf), err
	}

	dp, err := store.MySqlDB.QueryDataProviderByDid(did)
	if err != nil {
		return nil, err
	}
	if dp == nil {
		return nil, fmt.Errorf("no dp with did:%s", did)
	}
	if dp.DpType == common.DP_TYPE_OUTER {
		res, err := data.NewHttpDataProvider(dp)
		if err != nil {
			return nil, err
		}
		cache.GlobalCache.Add(key, 30*time.Minute, res)
		return res, nil
	}
	if dp.DpType == common.DP_TYPE_COMPOSITE {
		//todo
		//no logic here
	}
	if dp.DpType == common.DP_TYPE_CUSTOM {
		res, err := data.NewHttpDataProvider(dp)
		if err != nil {
			return nil, err
		}
		cache.GlobalCache.Add(key, 30*time.Minute, res)
		return res, nil
	}
	return nil, fmt.Errorf("not a supported type:%d in did:%s", dp.DpType, dp.Did)
}

type DataProviderInf interface {
	InvokeMethod(methodName string, paramMap map[string]interface{}) ([]byte, error)
	InvokeMethodWithParamStr(methodName string, paramStr string) ([]byte, error)
	VerifyDataSig(body []byte) (bool, error)
}
