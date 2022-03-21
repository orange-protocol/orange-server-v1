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
