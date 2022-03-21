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

package service

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/ontio/ontology/common"
	"github.com/stretchr/testify/assert"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/store"
)

func Test_SysDataService(t *testing.T) {
	url := "http://localhost:8081/query"
	method := "POST"

	payload := strings.NewReader(`{"query":"query{queryXdaysSumWithDefi(input:{key:\"test key\",user_did:\"did:ont:AXdmdzbyf3WZKQzRtrNQwAR91ZxMUfhXkt\",xdays:30,assets:[{chain:\"eth\",address:\"0x45929d79a6dddaa3c8154d4f245d17d1d80dbbcc\"},{chain:\"eth\",address:\"0x93c0957c3613d778ad42e386ea8ef8b7d2e1301e\"}],encrypt:true}){data{user_did}}}","variables":{}}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	//req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkaWQiOiJkaWQ6b250OmFiY2RlIiwiZXhwIjoxNjI3MDIyNjY5fQ.ZyIyQrnz6GNotdJf5YXXBDYV-aUx4cnA1fFxabSOVqk")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func Test_Sysdataservcie2(t *testing.T) {
	initenv()
	err := InitSysDataService("did:ont:abcdefg")
	assert.Nil(t, err)

	results, err := SysDS.GetUserAssetsDetail("did:ont:AHgtXRopCzzzSBtcLhv7YydRuLb4nwCuqr")
	if err != nil {
		assert.Nil(t, err)
	}
	assert.True(t, len(results) > 0)
	for _, r := range results {
		fmt.Printf("%v\n", r)
	}
}

func Test_GetWasmCode(t *testing.T) {
	//f := "./hello.wasm"
	//f := "./snapshot_space.wasm"
	//f := "./snapshot_space_opt.wasm"
	//f := "./monaco_opt.wasm"
	//f := "./monaco2_opt.wasm"
	//f := "./monaco_activity_opt.wasm"
	//f := "./monaco_contribution_opt.wasm"
	//f := "./monaco_composite_opt.wasm"
	//f := "./monaco_nftasset_opt.wasm"
	//f := "./monaco_userasset_opt.wasm"
	//f := "./chemix_opt.wasm"
	//f := "./votingdao_activity_opt.wasm"
	//f := "./activity_opt.wasm"
	f := "./investing_dao_optimized.wasm"
	//f := "./snapshot.wasm"
	//f := "./nft.wasm"
	bts, err := ioutil.ReadFile(f)
	assert.Nil(t, err)
	hexstr := hex.EncodeToString(bts)
	fmt.Printf("%s\n", hexstr)
	addr := common.AddressFromVmCode(bts)

	fmt.Printf("addr:%s\n", addr.ToHexString())

	err = store.InitMysql(&config.DB{
		UserName: "root",
		Password: "onchain",
		//DBAddr:   "43.134.54.213:3306",
		DBAddr: "172.168.3.45:3306",
		//DBAddr: "172.168.3.246:3306",
		DbName: "oscore",
	})
	assert.Nil(t, err)
	//err = store.MySqlDB.AddWasmCode("did:ont:AM9v32SveT18bKw8odpWn3g5cbCYj4sHEd", addr.ToHexString(), hexstr, "with investing_dao_opt.wasm")
	err = store.MySqlDB.AddWasmCode("did:ont:AQs9PzfeZKTP7MfyhEruCJgAfcmo1SMvX6", addr.ToHexString(), hexstr, "with investing_dao_opt.wasm")
	assert.Nil(t, err)

}
