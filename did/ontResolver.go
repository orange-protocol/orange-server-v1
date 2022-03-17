package did

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ontio/ontology-crypto/keypair"
	"github.com/ontio/ontology-crypto/signature"
	osdk "github.com/ontio/ontology-go-sdk"
	"github.com/ontio/ontology/common"

	"github.com/orange-protocol/orange-server-v1/config"
	"github.com/orange-protocol/orange-server-v1/log"
	"github.com/orange-protocol/orange-server-v1/utils"
)

type DidPubkey struct {
	Id           string      `json:"id"`
	Type         string      `json:"type"`
	Controller   interface{} `json:"controller"`
	PublicKeyHex string      `json:"publicKeyHex"`
}

type OntResolver struct {
	sdk  *osdk.OntologySdk
	acct *osdk.Account
}

func (o *OntResolver) ValidateSig(did string, idx int, msg string, sig string) (bool, error) {
	didpubkey, err := o.GetDidPubkeys(did, idx)
	if err != nil {
		log.Errorf("GetDidPubkeys failed:did:%s", err.Error())
		return false, err
	}

	pubkeybts, err := hex.DecodeString(didpubkey.PublicKeyHex)
	if err != nil {
		return false, err
	}
	pubkey, err := keypair.DeserializePublicKey(pubkeybts)
	if err != nil {
		return false, err
	}
	sigdata, err := hex.DecodeString(sig)
	if err != nil {
		return false, err
	}
	sign, err := signature.Deserialize(sigdata)
	if err != nil {
		return false, err
	}
	f := signature.Verify(pubkey, []byte(msg), sign)
	return f, nil
}

func (o *OntResolver) CreateCredential(did string, idx int, content interface{}, commit bool) (string, string, error) {
	contexts := []string{"context1", "context2"}
	expirationDays := config.GlobalConfig.DidConf[0].CredentialExpirationDays
	types := []string{"OscoreCredential"}
	expirationDate := time.Now().UTC().Unix() + 86400*int64(expirationDays)

	s, err := o.sdk.Credential.CreateJWTCredential(contexts, types, content, config.GlobalConfig.DidConf[0].DID, expirationDate,
		"", nil, o.acct)
	if err != nil {
		return "", "", err
	}

	if commit {
		cred, err := osdk.DeserializeJWT(s)
		if err != nil {
			return "", "", err
		}
		contractAddress, err := common.AddressFromHexString(cred.Payload.VC.CredentialStatus.Id)
		if err != nil {
			return "", "", err
		}
		txHash, err := o.sdk.Credential.CommitCredential(contractAddress, config.GlobalConfig.DidConf[0].Gasprice, config.GlobalConfig.DidConf[0].Gaslimit, cred.Payload.Jti,
			config.GlobalConfig.DidConf[0].DID, did, o.acct, o.acct)
		if err != nil {
			return "", "", err
		}
		return s, txHash.ToHexString(), nil
	} else {
		return s, "", nil
	}

}

func (o *OntResolver) RevokeCredential(did string, idx int, credstr string) (string, error) {
	cred, err := osdk.DeserializeJWT(credstr)
	if err != nil {
		return "", err
	}

	txhash, err := o.sdk.Credential.RevokeCredentialByIssuer(config.GlobalConfig.DidConf[0].Gasprice, config.GlobalConfig.DidConf[0].Gaslimit, cred.Payload.Jti,
		config.GlobalConfig.DidConf[0].DID, o.acct, o.acct)
	if err != nil {
		return "", err
	}
	return txhash.ToHexString(), nil
}

func (o *OntResolver) GetDidPubkeys(did string, idx int) (*DidPubkey, error) {
	ps, err := o.sdk.Native.OntId.GetPublicKeysJson(did)
	if err != nil {
		return nil, err
	}
	var pks []DidPubkey
	err = json.Unmarshal(ps, &pks)
	if err != nil {
		return nil, err
	}
	if len(pks) < idx {
		return nil, fmt.Errorf("no pubkey of did:%s with idx:%d", did, idx)
	}
	return &pks[idx-1], nil
}

func NewOntResolver(config *config.DidConf) (*OntResolver, error) {
	ontsdk := osdk.NewOntologySdk()
	ontsdk.NewRpcClient().SetAddress(config.URL)
	ontsdk.SetCredContractAddress(config.DIDContract)

	account, err := OpenAccount(config.Wallet, config.Password, ontsdk)
	if err != nil {
		return nil, err
	}
	return &OntResolver{sdk: ontsdk, acct: account}, nil
}

func OpenAccount(path string, pwd string, ontSdk *osdk.OntologySdk) (*osdk.Account, error) {
	wallet, err := ontSdk.OpenWallet(path)
	if err != nil {
		return nil, err
	}

	account, err := wallet.GetDefaultAccount([]byte(pwd))
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (o *OntResolver) SignData(data []byte) ([]byte, error) {
	return o.acct.Sign(data)
}

func (o *OntResolver) EncrypteDataWithDID(data []byte, did string) ([]byte, error) {
	pubkeyhex, err := o.GetDidPubkeys(did, 1)
	if err != nil {
		log.Errorf("errors on EncrypteDataWithDID:%s", err.Error())
		return nil, err
	}
	pubkey, err := hex.DecodeString(pubkeyhex.PublicKeyHex)
	if err != nil {
		log.Errorf("errors on EncrypteDataWithDID:%s", err.Error())
		return nil, err
	}
	return utils.EncryptWithDIDPubkey(data, pubkey)
}
