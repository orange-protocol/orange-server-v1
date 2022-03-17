package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/ontio/ontology-crypto/ec"
	"github.com/ontio/ontology-crypto/keypair"
	ontology_go_sdk "github.com/ontio/ontology-go-sdk"
)

func DecryptMsg(ecdsaPrivkey *ecdsa.PrivateKey, msg []byte) ([]byte, error) {
	prikey := ecies.ImportECDSA(ecdsaPrivkey)
	return prikey.Decrypt(msg, nil, nil)
}

func GetChainFromDID(did string) (string, error) {
	arr := strings.Split(did, ":")
	if len(arr) != 3 {
		return "", fmt.Errorf("did format error")
	}
	if arr[1] == "etho" {
		return "eth", nil
	}
	return arr[1], nil
}

func OpenAccount(path string, pwd string, addr string, ontSdk *ontology_go_sdk.OntologySdk) (*ontology_go_sdk.Account, error) {
	wallet, err := ontSdk.OpenWallet(path)
	if err != nil {
		return nil, err
	}

	account, err := wallet.GetAccountByAddress(addr, []byte(pwd))
	if err != nil {
		return nil, err
	}
	return account, nil
}
func PrivateKeyToEcdsaPrivkey(data []byte) (*ecdsa.PrivateKey, error) {
	c, err := keypair.GetCurve(data[1])
	if err != nil {
		return nil, err
	}
	size := (c.Params().BitSize + 7) >> 3
	if len(data) < size*2+3 {
		return nil, fmt.Errorf("deserializing private key failed: not enough length")
	}

	return ec.ConstructPrivateKey(data[2:2+size], c), nil
}
func EncryptWithDIDPubkey(msg []byte, didpubkey []byte) ([]byte, error) {
	ecdsaPubkey, err := UnmarshalPubkey(didpubkey)
	if err != nil {
		return nil, err
	}

	eciesPubkey := ecies.ImportECDSAPublic(ecdsaPubkey)

	return ecies.Encrypt(rand.Reader, eciesPubkey, msg, nil, nil)
}

func UnmarshalPubkey(data []byte) (*ecdsa.PublicKey, error) {
	pub, err := ec.DecodePublicKey(data, elliptic.P256())
	if err != nil {
		return nil, fmt.Errorf("deserializing public key failed: decode P-256 public key error")
	}

	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: pub.X, Y: pub.Y}, nil
}

func AddressInArray(addr string, allowed []string) bool {
	for _, a := range allowed {
		if strings.EqualFold(addr, a) {
			return true
		}
	}
	return false
}
