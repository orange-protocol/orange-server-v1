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

package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JsonSchema(t *testing.T) {
	sl := `{
  "title": "Person",
  "description": "it is a person object",
  "type": "object",
  "properties": {
    "id": {
      "description": "The unique identifier for a person",
      "type": "string"
    },
    "name": {
      "description": "The identifier for a person",
      "type": "string",
      "maxLength": 50
    }
  },
  "required": [ "id", "name" ]
}`

	doc := `{"id":"abc","name":"test"}`
	r, err := VerifyJSonSchema(sl, doc)
	assert.Nil(t, err)
	assert.True(t, r)

}

func Test_xdaysumwithDefi(t *testing.T) {
	s := `{
		"title":"xdaysSumInfoWithDefi",
		"definitions":{
			"suminfo":{
				"type":"object",
				"properties":{
					"amount":{"type":"string"},
					"days":{"type":"number"}
				}
			}
		},
		"description":"xdays assets with defi data",
		"type":"object",
		"properties":{
			"user_did":{
				"description":"user's did",
				"type":"string"
			},
			"defi_assets":{
				"description":"user total assets info with defi",
				"type":"object",
				"properties":{
					"asset_infos":{
						"description":"token assets info",
						"type":"array",
						"items":[
							{
								"type":"object",
								"properties":{
									"token_name":{"type":"string"},
									"balance":{"type":"string"},
									"xday_sum":{"#ref":"#/definitions/suminfo"},
									"price":{"type":"string"}
								}
							}
						]
					}
				}
			},
			"sig":{
				"description":"signature",
				"type":"string"
			}
		}
	}`

	doc := `{"defi_assets":{"asset_infos":[{"balance":"0.090158756028517988","price":"3544.86000000","token_name":"eth","xday_sum":{"amount":"2.70476268085553964","days":30}},{"balance":"0","price":"1.00","token_name":"usdt","xday_sum":{"amount":"0","days":30}},{"balance":"0","price":"0.99970000","token_name":"usdc","xday_sum":{"amount":"0","days":30}},{"balance":"0","price":"3544.86000000","token_name":"eth","xday_sum":{"amount":"0","days":30}},{"balance":"0","price":"1.00","token_name":"usdt","xday_sum":{"amount":"0","days":30}},{"balance":"0","price":"0.99970000","token_name":"usdc","xday_sum":{"amount":"0","days":30}}],"defi_infos":[{"chain":"eth","defi_name":"aavev2","net_balance":"0.00","xday_sum":{"amount":"0.00","days":30}},{"chain":"eth","defi_name":"aavev2","net_balance":"0.00","xday_sum":{"amount":"0.00","days":30}}]},"sig":"7f2f7284414caeb67b469d7cbc1303c2a152cf58a994b2d88409d76780b095beb9090c052c4df2b45567d72a30f4db06fe366aad2e8b999f205a86f1b99b4749","user_did":"did:ont:AGAMr5P2Ngi7SGvhKd3s5vWTWpid5uGywL"}`

	r, err := VerifyJSonSchema(s, doc)
	assert.Nil(t, err)
	assert.True(t, r)
}
