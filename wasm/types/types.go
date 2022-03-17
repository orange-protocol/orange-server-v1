package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

/*
request-param: {
	"tokens": ["ONT", "ETH"], //["ALL"],
	"query_items" : {
		"balances": bool,
		"xday_sum": 10,
		"price": bool,
	}
}

// 所有的token数字转为decimial为10的定点数;
response: {
	"asset_infos" [{
		"token_name": 'ONT",
		"balance": "100000000000",
		"xday_sum": {"amount": "1000000000000000", days: 10}，
		"price": "37000000000", // $
	}, {
		"token_name": 'ETH",
		"balance": "100000000000",
		"xday_sum": {"amount": "1000000000000000", days: 10}，
		"price": "37000000000", // $
	}
}
*/

type ScoreResult struct {
	Score uint `json:"score"`
}

type RequestParam struct {
	Tokens     []string    `json:"tokens"`
	QueryItems *QueryItems `json:"query_items,omitempty"`
}

func (self RequestParam) Value() (driver.Value, error) {
	buf, err := json.Marshal(self)
	if err != nil {
		return nil, err
	}
	return string(buf), nil
}

func (self *RequestParam) Scan(src interface{}) error {
	var source []byte
	switch t := src.(type) {
	case string:
		source = []byte(t)
	case []byte:
		if len(t) == 0 {
			source = []byte("{}")
		} else {
			source = t
		}
	default:
		return errors.New("incompatible type for RequestParam")
	}

	return json.Unmarshal(source, self)
}

type QueryItems struct {
	Balance bool   `json:"balance"`
	XdaySum uint32 `json:"xday_sum"`
	Price   bool   `json:"price"`
}

type AssetInfoData struct {
	AssetInfos []*AssetInfo `json:"asset_infos"`
	UserDid    string       `json:"user_did"`
}

type AssetInfo struct {
	TokenName string   `json:"token_name"`
	Balance   string   `json:"balance"`
	XdaySum   *XdaySum `json:"xday_sum"`
	Price     string   `json:"price"`
}

type XdaySum struct {
	Days   uint   `json:"days"`
	Amount string `json:"amount"`
}
