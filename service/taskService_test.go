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
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/store"
)

func initenv() {
	//err := config.LoadConfig("../config.json")
	//if err != nil {
	//	fmt.Println("error on load config")
	//	panic(err)
	//}

	cfg := &config.DB{
		UserName: "root",
		Password: "onchain",
		DBAddr:   "172.168.3.45:3306",
		DbName:   "oscore",
	}

	err := store.InitMysql(cfg)
	if err != nil {
		panic(err)
	}
	//cache.GlobalCache = cache2go.Cache("SYSCACHE")

}

func Test_resolveMethod(t *testing.T) {
	initenv()
	//ts := NewTaskService()

	tasks, err := store.MySqlDB.QueryTasksByUserDID("did:ont:abcde")
	assert.Nil(t, err)
	assert.Equal(t, len(tasks), 1)

	//err = ts.resovleDataProvider(tasks[0])
	//assert.Nil(t, err)
}

func Test_ap(t *testing.T) {
	initenv()
	//ts := NewTaskService()

	tasks, err := store.MySqlDB.QueryTasksByUserDID("did:ont:AV88PcsdFk2MTcPkuyPNEkpgLFiKHtCM1r")
	assert.Nil(t, err)
	fmt.Printf("%v\n", tasks)
	for _, t := range tasks {
		fmt.Printf("name:%s\n", t.ApName)
	}
	//assert.Equal(t, len(tasks), 6)
	//err = ts.resolveAlgorithmProvider(tasks[0])
	//assert.Nil(t, err)

}

func Test_DataService(t *testing.T) {

	url := "http://localhost:8081/query"
	method := "POST"
	//s := "{\"query\":\"query{queryXdaysSum(input:{key:\\\"test key\\\",user_did:\\\"did:ont:abcde\\\",xdays:30,assets:[{chain:\\\"eth\\\",address:\\\"0xabcde\\\"}]}){user_did,asset_infos{token_name,balance,xday_sum{amount,days},price}}}\",\"variables\":{}}"
	//s := "{\"query\":\"query{queryXdaysSumWithDefi(input:{key:\\\"eee81c2d-350d-4bb5-9fe1-f2ca9f9bcb10\\\",user_did:\\\"did:ont:ARNzB1pTkG61NDwxwzJfNJF8BqcZjpfNev\\\",xdays:30,assets:[{chain:\\\"eth\\\",address:\\\"0xeF8305E140ac520225DAf050e2f71d5fBcC543e7\\\"},{chain:\\\"eth\\\",address:\\\"0x45929d79a6dddaa3c8154d4f245d17d1d80dbbcc\\\"},{chain:\\\"eth\\\",address:\\\"0x93c0957c3613d778ad42e386ea8ef8b7d2e1301e\\\"}],encrypt:true}){data{user_did,defi_assets{asset_infos{token_name,balance,xday_sum{amount,days},price},defi_infos{chain,defi_name,net_balance,xday_sum{amount,days}}},sig},encrypted}}\\\",\\\"variables\\\":{}}"
	s := "{\"query\":\"query{queryNFTTotal(input:{key:\\\"a1bd9450-cabc-463e-8c1d-9c833b6c83e4\\\",user_did:\\\"did:ont:ARNzB1pTkG61NDwxwzJfNJF8BqcZjpfNev\\\",xdays:30,assets:[{chain:\\\"eth\\\",address:\\\"0xeF8305E140ac520225DAf050e2f71d5fBcC543e7\\\"},{chain:\\\"eth\\\",address:\\\"0x45929d79a6dddaa3c8154d4f245d17d1d80dbbcc\\\"},{chain:\\\"eth\\\",address:\\\"0x93c0957c3613d778ad42e386ea8ef8b7d2e1301e\\\"}],encrypt:true}){data{data{latest_transfer_days_till_now,transfer_count_for_last_year,earliest_transfer_days_till_now,owned_nft_kinds_count,,owned_nft_count,current_nft_count,current_nft_value_in_eth},sig},encrypted}}\",\"variables\":{}}"
	fmt.Printf("%s\n", s)
	//payload := strings.NewReader("{\"query\":\"query{\\n\\tqueryXdaysSum(input:{key:\\\"test key\\\",user_did:\\\"did:ont:abcde\\\",xdays:30,assets:[{chain:\\\"eth\\\",address:\\\"0xabcde\\\"}]}){\\n    user_did\\n    asset_infos{\\n      token_name\\n      balance\\n      xdays_sum{\\n        amount,\\n        days\\n      }\\n\\tprice\\n    }\\n  }\\n}\",\"variables\":{}}")
	//payload := strings.NewReader("{\"query\":\"query{queryXdaysSum(input:{key:\\\"test key\\\",user_did:\\\"did:ont:abcde\\\",xdays:30,assets:[{chain:\\\"eth\\\",address:\\\"0xabcde\\\"}]}){user_did}}\",\"varialbles\",{}}")
	payload := strings.NewReader(s)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
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

func Test_ParseParam(t *testing.T) {
	//s := `{"query":"query{queryAssets(input:{key:$API_KEY,assets:$ARRAY[{chain:$CHAIN_NAME,walletAddress:$USER_ADDRESS}],encrypted:$ENCRYPTED}){chain,address,name,balance,price,value,time,icon}}","variables":{}}`
	//winfo := []*store.UserAddressInfo{{Chain: "eth", Address: "0x12345"}, {Chain: "eth", Address: "0x67890"}}
	//fmt.Printf("res:%s\n", s)

	//res, _, err := ParseInputParam(s, "test-api", winfo, true)
	//assert.Nil(t, err)
	//fmt.Printf("res:%s\n", res)

}
func Test_ParseParam2(t *testing.T) {
	//s := `{"query":"query{queryAssets(input:{key:$API_KEY,assets:$ARRAY[$USER_ADDRESS],space_in:$ARRAY[$CHAIN_NAME],encrypted:$ENCRYPTED}){chain,address,name,balance,price,value,time,icon}}","variables":{}}`
	//winfo := []*store.UserAddressInfo{{Chain: "eth", Address: "0x12345"}, {Chain: "eth", Address: "0x67890"}}
	//fmt.Printf("res:%s\n", s)

	//r := regexp.MustCompile("\\$ARRAY\\[(.*?)\\]")
	//r := regexp.MustCompile(`$ARRAY[(.*)]`)
	//arr := r.FindAllString(s,-1)
	//arr := r.FindAllStringSubmatch(s,-1)
	//fmt.Printf("%v\n",arr)

	//res, _, err := ParseInputParam(s, "test-api", winfo, true)
	//assert.Nil(t, err)
	//fmt.Printf("res:%s\n", res)

}

func Test_match(t *testing.T) {
	//re := regexp.MustCompile("a(x*)b")
	//fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-", -1))
	//fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-", -1))
	//fmt.Printf("%q\n", re.FindAllStringSubmatch("-ab-axb-axxxxb", -1))
	//fmt.Printf("%q\n", re.FindAllStringSubmatch("-axxb-ab-", -1))

	r := regexp.MustCompile("\\$ARRAY\\[(.*)\\]")
	fmt.Printf("%q\n", r.FindAllStringSubmatch(`assets:$ARRAY[$USER_ADDRESS],space_in:$ARRAY[$CHAIN_NAME]`, -1))

}
