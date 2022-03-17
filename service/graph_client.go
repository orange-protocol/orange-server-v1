package service

import (
	"context"
	"fmt"
	"time"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/log"

	gqlclient "github.com/machinebox/graphql"
)

const (
	TheGraphDefaultTimeout = time.Second * 300
)

type TheGraphService struct {
	gclients map[string]*gqlclient.Client
}

func NewTheGraphService(cfg *config.GraphConfig) *TheGraphService {
	c := make(map[string]*gqlclient.Client)
	c["eth"] = gqlclient.NewClient(cfg.Eth)
	c["bsc"] = gqlclient.NewClient(cfg.Bsc)
	c["polygon"] = gqlclient.NewClient(cfg.Polygon)
	return &TheGraphService{gclients: c}
}

type Token struct {
	Id     string `json:"id,omitempty"`
	MintTx string `json:"minttx,omitempty"`
}

type UserData struct {
	Tokens []*Token `json:"tokens,omitempty"`
}

type RespData struct {
	Users []*UserData `json:"users,omitempty"`
}

func (this *TheGraphService) QueryHashByAddressFromChain(address, chain string) ([]*Token, error) {
	reqStr := `
	query users($id: String){
          users ( where: {
          id: $id
  		  }) {
   		tokens {
      		id
      		mintTx
    	} 
  	   }
	}
	`

	req := gqlclient.NewRequest(reqStr)
	req.Var("id", address)
	req.Header.Set("Cache-Control", "no-cache")
	ctx, cf := context.WithTimeout(context.Background(), TheGraphDefaultTimeout)
	defer cf()
	ret := &RespData{}
	if _, present := this.gclients[chain]; !present {
		log.Errorf("chain:%s not exist", chain)
		return nil, fmt.Errorf("chain:%s not exist", chain)
	}
	client := this.gclients[chain]
	err := client.Run(ctx, req, ret)
	if err != nil {
		log.Errorf("QueryHashByAddressFromChain err:%s", err)
		return nil, err
	}
	tokens := make([]*Token, 0)
	for _, data := range ret.Users {
		tokens = append(tokens, data.Tokens...)
	}
	return tokens, nil
}
