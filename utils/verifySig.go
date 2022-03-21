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
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	ethcomm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology/core/signature"
	"github.com/ontio/ontology/core/types"
	"math/big"
	"strings"

	"github.com/orange-protocol/orange-server-v1/log"
)

const (
	BTC_CHAIN  = "btc"
	ETH_CHAIN  = "eth"
	ONT_CHAIN  = "ont"
	TRON_CHAIN = "trx"
	DOT_CHAIN  = "dot"
	KLAY_CHAIN = "klay"
	NEO_CHAIN  = "neo"
)

const maxVarintBytes = 10 // maximum length of a varint

var BITCOIN_SIGNED_MESSAGE_HEADER = "Bitcoin Signed Message:\n"

func VerifyDIDSigs(chain, address, did, sig, pubkey string) bool {

	switch chain {
	case ETH_CHAIN:
		return ETHVerifySig(address, sig, []byte(did))
	case ONT_CHAIN:
		return ONTVerifySig(address, pubkey, sig, []byte(did))
	case BTC_CHAIN:
		return BTCVerifySig(address, pubkey, sig, []byte(did))
	case NEO_CHAIN:
		return NEOVerifySig(address, pubkey, sig, []byte(did))
	//case TRON_CHAIN:
	//	return TRONVerifySig(address, pubkey, sig, []byte(did))
	case KLAY_CHAIN:
		return KlayVerifySig(address, pubkey, sig, []byte(did))
	}
	return false
}

func EthSignHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func TronSignHash(data []byte) []byte {
	msg := fmt.Sprintf("\x19TRON Signed Message:\n%d%s", len(data), data)
	return crypto.Keccak256([]byte(msg))
}

func ETHVerifySig(from, sigHex string, msg []byte) bool {
	fromAddr := ethcomm.HexToAddress(from)

	sig := hexutil.MustDecode(sigHex)
	if len(sig) < 64 {
		return false
	}
	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	if sig[64] != 27 && sig[64] != 28 {
		return false
	}
	sig[64] -= 27

	pubKey, err := crypto.SigToPub(EthSignHash(msg), sig)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return strings.EqualFold(fromAddr.Hex(), recoveredAddr.Hex())
}

func ONTVerifySig(address, pubkey, sigHex string, msg []byte) bool {
	pkbytes, err := hex.DecodeString(pubkey)
	if err != nil {
		return false
	}
	pk, err := keypair.DeserializePublicKey(pkbytes)
	if err != nil {
		return false
	}
	addr := types.AddressFromPubKey(pk)
	if addr.ToBase58() != address {
		return false
	}

	sigbytes, err := hex.DecodeString(sigHex)
	if err != nil {
		return false
	}
	err = signature.Verify(pk, msg, sigbytes)
	return err == nil

}

func NEOVerifySig(address, pubkey, sigHex string, msg []byte) bool {

	//test with ontology
	pubkeydata, err := hex.DecodeString(pubkey)
	if err != nil {
		return false
	}
	pk, err := keypair.DeserializePublicKey(pubkeydata)
	if err != nil {
		return false
	}

	addr := types.AddressFromPubKey(pk)
	if !strings.EqualFold(addr.ToBase58(), address) {
		return false
	}

	sigdata, err := hex.DecodeString(sigHex)
	err = signature.Verify(pk, msg, sigdata)
	if err != nil {
		return false
	}
	return true
}

//
//func TRONVerifySig(address, publickey, sigHex string, msg []byte) bool {
//
//	sig := hexutil.MustDecode(sigHex)
//	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
//	if sig[64] != 27 && sig[64] != 28 {
//		return false
//	}
//	sig[64] -= 27
//	pubKey, err := crypto.SigToPub(TronSignHash(msg), sig)
//	if err != nil {
//		return false
//	}
//
//	addr := tron.PubkeyToAddress(*pubKey)
//	return strings.EqualFold(addr.String(), address)
//}
func KlayVerifySig(address, publickey, sigHex string, msg []byte) bool {

	return ETHVerifySig(address, sigHex, msg)
}

//todo currently , only support btc address from onto
func BTCVerifySig(address, pubkey, sigHex string, msg []byte) bool {

	sigbytes, err := hex.DecodeString(sigHex)
	if err != nil {
		log.Errorf("[BTCVerifySig]DecodeString failed:%s", err.Error())
		return false
	}

	l := len(sigbytes)
	if l != 65 {
		log.Errorf("sig length is not 65")
		return false
	}
	rbytes := sigbytes[1:33]
	sbytes := sigbytes[33:]

	R := new(big.Int).SetBytes(rbytes)
	S := new(big.Int).SetBytes(sbytes)
	sig := btcec.Signature{
		R: R,
		S: S,
	}

	pubkeybytes, err := hex.DecodeString(pubkey)
	if err != nil {
		log.Errorf("[BTCVerifySig]DecodeString failed:%s", err.Error())
		return false
	}

	//todo check address and pubkey
	//distinguish address from onto or other source(cyano)
	/*	apk,err := btcutil.NewAddressPubKey(pubkeybytes,&chaincfg.MainNetParams)
		addr := apk.EncodeAddress()
		if !strings.EqualFold(addr,address){
			return false
		}*/

	publickey, err := btcec.ParsePubKey(pubkeybytes, btcec.S256())
	if err != nil {
		log.Errorf("[BTCVerifySig]DecodeString failed:%s", err.Error())
		return false
	}

	fmttedMsg := FormatMessageForSigning(msg)

	return sig.Verify(fmttedMsg, publickey)

}

func FormatMessageForSigning(msg []byte) []byte {

	bts := bytes.NewBuffer(nil)
	bts.Write(IntToBytes(len([]byte(BITCOIN_SIGNED_MESSAGE_HEADER))))
	bts.Write([]byte(BITCOIN_SIGNED_MESSAGE_HEADER))

	bts.Write(IntToBytes(len(msg)))
	bts.Write(msg)
	tmpbs := bts.Bytes()

	s := sha256.New()
	_, err := s.Write(tmpbs)
	if err != nil {
		return nil
	}
	once := s.Sum(nil)
	s = sha256.New()
	_, err = s.Write(once)
	if err != nil {
		return nil
	}
	twice := s.Sum(nil)
	return twice
}

func IntToBytes(n int) []byte {
	return EncodeVarint(uint64(n))
}

func EncodeVarint(x uint64) []byte {
	var buf [maxVarintBytes]byte
	var n int
	for n = 0; x > 127; n++ {
		buf[n] = 0x80 | uint8(x&0x7F)
		x >>= 7
	}
	buf[n] = uint8(x)
	n++
	return buf[0:n]
}
