package common

import (
	"fmt"
	"strings"

	"github.com/orange-protocol/orange-server-v1/store"
)

const (
	AP_TYPE_OUTER = 0
	AP_TYPE_WASM  = 1

	DP_TYPE_OUTER     = 0
	DP_TYPE_COMPOSITE = 1
	DP_TYPE_CUSTOM    = 2
)

type Method struct {
	Url        string
	Param      string
	Result     string
	HttpMethod string
}

func ProcessParamMap(m *Method, paramMap map[string]interface{}) string {

	paramstr := ""
	if len(m.Param) > 0 {
		paramstr = m.Param
		for k, v := range paramMap {
			//todo fixme
			paramstr = strings.ReplaceAll(paramstr, k, fmt.Sprintf("%v", v))
		}
	}
	return paramstr
}

func TransformTaskStatus(status int) string {
	res := ""
	switch status {
	case store.TASK_STATUS_INIT:
		res = "INIT"
	case store.TASK_STATUS_RESOLVING:
		res = "RESOLVING"
	case store.TASK_STATUS_DP_QUERYING:
		res = "QUERYING DATA"
	case store.TASK_STATUS_DP_FINISHED:
		res = "DATA COLLECTED"
	case store.TASK_STATUS_DP_FAILED:
		res = "DATA COLLECT FAILED"
	case store.TASK_STATUS_AP_RESOLVING:
		res = "ALGORITHM RESOLVING"
	case store.TASK_STATUS_AP_QUERYING:
		res = "ALGORITHM CALCULATING"
	case store.TASK_STATUS_AP_FINISHED:
		res = "FINISHED"
	case store.TASK_STATUS_AP_FAILED:
		res = "ALGORITHM FAILED"
	case store.TASK_STATUS_VC_STARTING:
		res = "CREDENTIAL STARTING"
	case store.TASK_STATUS_VC_GENERATING:
		res = "CREDENTIAL GENERATING"
	case store.TASK_STATUS_VC_FAILED:
		res = "CREDENTIAL FAILED"
	case store.TASK_STATUS_DONE:
		res = "TASK SUCCEED"
	default:
		res = ""
	}
	return res
}
