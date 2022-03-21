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

package eth

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/orange-protocol/orange-server-v1/utils"
)

type EthResolver struct {
	client       *ethclient.Client
	privateKey   *ecies.PrivateKey
	ididContract *IDid
}

func (e EthResolver) ValidateSig(did string, idx int, msg string, sig string) (bool, error) {
	pubkey, err := e.GetPubkeyByDID(did, idx)
	if err != nil {
		return false, err
	}

	ecdsaPub, err := crypto.UnmarshalPubkey(pubkey.PubKeyData)
	if err != nil {
		return false, err
	}
	ethAddress := crypto.PubkeyToAddress(*ecdsaPub).String()
	return utils.ETHVerifySig(ethAddress, sig, []byte(msg)), nil

}

func (e EthResolver) CreateCredential(did string, idx int, content interface{}, commit bool) (string, string, error) {
	panic("implement me")
}

func (e EthResolver) RevokeCredential(did string, idx int, cred string) (string, error) {
	panic("implement me")
}

func (e EthResolver) GetPubkeyByDID(did string, idx int) (*IDidPublicKey, error) {
	pubkeys, err := e.ididContract.GetAllPubKey(nil, did)
	if err != nil {
		return nil, fmt.Errorf("GetPubkeyByDID failed:%s", err.Error())
	}

	if len(pubkeys) < idx {
		return nil, fmt.Errorf("GetPubkeyByDID idx out of bound:%d", idx)
	}

	return &pubkeys[idx-1], nil
}
func (o *EthResolver) SignData(data []byte) ([]byte, error) {
	panic("implement me")
}

func (o *EthResolver) EncrypteDataWithDID(data []byte, did string) ([]byte, error) {
	panic("implement me")
}
