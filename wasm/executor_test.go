package wasm

//func TestNewExecutor(t *testing.T) {
//	code, err := ioutil.ReadFile("data/hello.wasm")
//	ensure(err)
//	start := time.Now()
//	input := `
//{
//	"user_did":"mydid",
//	"asset_infos": [{
//		"token_name": "ONT",
//		"balance": "100000000000",
//		"xday_sum": {"amount": "1000000000000000", "days": 10},
//		"price": "37000000000"
//	}, {
//		"token_name": "ETH",
//		"balance": "100000000000",
//		"xday_sum": {"amount": "1000000000000000", "days": 10},
//		"price": "37000000000"
//	}]
//}
//`
//	//var assetInfo types.AssetInfoData
//	//err = json.Unmarshal([]byte(input), &assetInfo)
//	//ensure(err)
//	executor := NewExecutor([]byte(input), 100000, 1000000, &ExecEnv{nil})
//	result, err := executor.Invoke(code)
//	ensure(err)
//
//	fmt.Printf("execute time : %f\n", float64(time.Now().Sub(start))/float64(time.Millisecond))
//	fmt.Printf("execute: result %v", result)
//}
