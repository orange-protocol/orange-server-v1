package service

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/utils"
)

func InitTestNftClaimService() *NftClaimService {
	walletcfg := &config.ETHWallet{
		KeyStore: "../keystore",
		Password: "123456",
	}

	infos := make(map[string]config.NFTInfo)

	infos["bsc"] = config.NFTInfo{
		ContractAddress: "0x9aa240baf44c2a510ab1d2778386eaf144c77247",
		Rpc:             "https://speedy-nodes-nyc.moralis.io/6eb43157cbc67a17e7644196/bsc/testnet",
	}
	infos["eth"] = config.NFTInfo{
		ContractAddress: "0x5f3c3ea1de47a2930ba8dbe436cf2ec5382b2584",
		Rpc:             "https://speedy-nodes-nyc.moralis.io/6eb43157cbc67a17e7644196/eth/kovan",
	}

	nftcfg := &config.NFTConfig{NftInfos: infos}

	cfg := &config.SysConfig{ETHWallet: walletcfg, NFTConfig: nftcfg}

	return NewNftClaimService(cfg)
}

func TestNftClaimService_GetNFTDetail(t *testing.T) {
	s := InitTestNftClaimService()

	detail, err := s.GetNFTDetail("eth", "0x5f3c3ea1de47a2930ba8dbe436cf2ec5382b2584", 1)
	assert.Nil(t, err)
	assert.NotNil(t, detail)
	fmt.Printf("owner:%s\n", detail.OriginOwner)
	fmt.Printf("ValidTo:%d\n", detail.ValidTo)
	fmt.Printf("Score:%s\n", detail.Score)
	//fmt.Sprintf("owner:%s\n",detail.OriginOwner)
}

func TestNftClaimService_GetUserClaimHash(t *testing.T) {
	s := InitTestNftClaimService()
	bts, err := s.GetUserClaimHash("0x26356Cb66F8fd62c03F569EC3691B6F00173EB02", 2, 100)
	assert.Nil(t, err)
	fmt.Printf("hash:%s\n", hexutil.Encode(bts))

	sig, err := s.SignMsg(bts)
	//fmt.Printf("len:%d\n", len(sig))
	//fmt.Printf("sig[64] is %v\n", sig[64])
	assert.Nil(t, err)
	sighex := hexutil.Encode(sig)
	fmt.Printf("sig:%s\n", sighex)

	//f := utils.ETHVerifySig(s.ks.Accounts()[0].Address.String(),sighex,msghash)
	//assert.True(t,f)
}

func Test_abipack(t *testing.T) {
	s := InitTestNftClaimService()
	bts, err := s.GetUserClaimHash("0x26356Cb66F8fd62c03F569EC3691B6F00173EB02", 1, 100)
	assert.Nil(t, err)
	fmt.Printf("hash:%s\n", hexutil.Encode(bts))

}

func Test_sig(t *testing.T) {
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	assert.Nil(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	assert.True(t, ok)

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	data := []byte("hello")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex()) // 0x1c8aff950685c2ed4bc3174f3472287b56d9517b9c948127319a09a7a36deac8

	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	assert.Nil(t, err)
	fmt.Printf("len sig is %d\n", len(signature))
	fmt.Printf(" sig[64] is %d\n", signature[64])

	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62d509bf6049b8fb544ab06d36f916685a2eb8e57ffadde02301

	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	assert.Nil(t, err)

	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches) // true

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	assert.Nil(t, err)

	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches) // true

	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery id
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified) // true
}

func Test_sig11(t *testing.T) {
	msg := "did:etho:93c0957c3613d778ad42e386ea8ef8b7d2e1301e"
	sig := "0x2a53fe175352f537d91f8213b9d16c0cbda68dd5e52ac718f1225dcbb45df9615eff379cb24ffd82b811b45293302eb1cb527f2a7f77ae67f0fb354f95cd467a1b"
	addr := "0x93c0957c3613d778ad42e386ea8ef8b7d2e1301e"

	f := utils.ETHVerifySig(addr, sig, []byte(msg))
	assert.True(t, f)
}
