// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IDidDIDDocument is an auto generated low-level Go binding around an user-defined struct.
type IDidDIDDocument struct {
	Context        []string
	Id             string
	PublicKey      []IDidPublicKey
	Authentication []IDidPublicKey
	Controller     []string
	Service        []IDidService
	Updated        *big.Int
}

// IDidPublicKey is an auto generated low-level Go binding around an user-defined struct.
type IDidPublicKey struct {
	Id          string
	KeyType     string
	Controller  []string
	PubKeyData  []byte
	Deactivated bool
	IsPubKey    bool
	AuthIndex   *big.Int
}

// IDidService is an auto generated low-level Go binding around an user-defined struct.
type IDidService struct {
	ServiceId       string
	ServiceType     string
	ServiceEndpoint string
}

// IDidMetaData contains all meta data concerning the IDid contract.
var IDidMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"}],\"name\":\"AddAddr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"context\",\"type\":\"string\"}],\"name\":\"AddContext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"}],\"name\":\"AddController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"}],\"name\":\"AddKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"}],\"name\":\"AddNewAuthAddr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"}],\"name\":\"AddNewAuthKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"serviceId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"serviceType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"serviceEndpoint\",\"type\":\"string\"}],\"name\":\"AddService\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"Deactivate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"DeactivateAddr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"DeactivateAuthAddr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"DeactivateAuthKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"DeactivateKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"context\",\"type\":\"string\"}],\"name\":\"RemoveContext\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"}],\"name\":\"RemoveController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"serviceId\",\"type\":\"string\"}],\"name\":\"RemoveService\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"SetAuthAddr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"SetAuthKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"serviceId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"serviceType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"serviceEndpoint\",\"type\":\"string\"}],\"name\":\"UpdateService\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"pubKeyController\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"addAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"context\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"addContext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"addController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"newPubKey\",\"type\":\"bytes\"},{\"internalType\":\"string[]\",\"name\":\"pubKeyController\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"addKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"addNewAuthAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"controllerSigner\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"addNewAuthAddrByController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"addNewAuthKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"controllerSigner\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"addNewAuthKeyByController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceEndpoint\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"addService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"deactivateAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"deactivateAuthAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"deactivateAuthAddrByController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"deactivateAuthKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"deactivateAuthKeyByController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"deactivateID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"deactivateKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getAllAuthKey\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"keyType\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"deactivated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPubKey\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"authIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIDid.PublicKey[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getAllController\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getAllPubKey\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"keyType\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"deactivated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPubKey\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"authIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIDid.PublicKey[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getAllService\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"serviceId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceEndpoint\",\"type\":\"string\"}],\"internalType\":\"structIDid.Service[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getContext\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getDocument\",\"outputs\":[{\"components\":[{\"internalType\":\"string[]\",\"name\":\"context\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"keyType\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"deactivated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPubKey\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"authIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIDid.PublicKey[]\",\"name\":\"publicKey\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"keyType\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"deactivated\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPubKey\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"authIndex\",\"type\":\"uint256\"}],\"internalType\":\"structIDid.PublicKey[]\",\"name\":\"authentication\",\"type\":\"tuple[]\"},{\"internalType\":\"string[]\",\"name\":\"controller\",\"type\":\"string[]\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"serviceId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceEndpoint\",\"type\":\"string\"}],\"internalType\":\"structIDid.Service[]\",\"name\":\"service\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"updated\",\"type\":\"uint256\"}],\"internalType\":\"structIDid.DIDDocument\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getUpdatedTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"context\",\"type\":\"string[]\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"removeContext\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"removeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceId\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"removeService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"setAuthAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"setAuthAddrByController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"setAuthKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"setAuthKeyByController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceEndpoint\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"singer\",\"type\":\"bytes\"}],\"name\":\"updateService\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"53864503": "addAddr(string,address,string[],bytes)",
		"65cbf7c3": "addContext(string,string[],bytes)",
		"0ec0865b": "addController(string,string,bytes)",
		"52c07c56": "addKey(string,bytes,string[],bytes)",
		"8da5588c": "addNewAuthAddr(string,address,string[],bytes)",
		"b0abb199": "addNewAuthAddrByController(string,address,string[],string,bytes)",
		"002b4c82": "addNewAuthKey(string,bytes,string[],bytes)",
		"db7d1f04": "addNewAuthKeyByController(string,bytes,string[],string,bytes)",
		"7d1c754f": "addService(string,string,string,string,bytes)",
		"d558cdea": "deactivateAddr(string,address,bytes)",
		"6747e9bc": "deactivateAuthAddr(string,address,bytes)",
		"88209cf3": "deactivateAuthAddrByController(string,address,string,bytes)",
		"958607b8": "deactivateAuthKey(string,bytes,bytes)",
		"9053ad03": "deactivateAuthKeyByController(string,bytes,string,bytes)",
		"ec02a3b5": "deactivateID(string,bytes)",
		"793f66c9": "deactivateKey(string,bytes,bytes)",
		"d22eafd9": "getAllAuthKey(string)",
		"2e35d7ca": "getAllController(string)",
		"1f7c7a9d": "getAllPubKey(string)",
		"f8d8f4ad": "getAllService(string)",
		"53949d4f": "getContext(string)",
		"7ccb6a64": "getDocument(string)",
		"ba964281": "getUpdatedTime(string)",
		"8095627f": "removeContext(string,string[],bytes)",
		"244d9153": "removeController(string,string,bytes)",
		"2e97b8d4": "removeService(string,string,bytes)",
		"398a85ec": "setAuthAddr(string,address,bytes)",
		"6e5f3183": "setAuthAddrByController(string,address,string,bytes)",
		"fa1bac72": "setAuthKey(string,bytes,bytes)",
		"9033d99a": "setAuthKeyByController(string,bytes,string,bytes)",
		"667f5fa3": "updateService(string,string,string,string,bytes)",
	},
}

// IDidABI is the input ABI used to generate the binding from.
// Deprecated: Use IDidMetaData.ABI instead.
var IDidABI = IDidMetaData.ABI

// Deprecated: Use IDidMetaData.Sigs instead.
// IDidFuncSigs maps the 4-byte function signature to its string representation.
var IDidFuncSigs = IDidMetaData.Sigs

// IDid is an auto generated Go binding around an Ethereum contract.
type IDid struct {
	IDidCaller     // Read-only binding to the contract
	IDidTransactor // Write-only binding to the contract
	IDidFilterer   // Log filterer for contract events
}

// IDidCaller is an auto generated read-only Go binding around an Ethereum contract.
type IDidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IDidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IDidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IDidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IDidSession struct {
	Contract     *IDid             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IDidCallerSession struct {
	Contract *IDidCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IDidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IDidTransactorSession struct {
	Contract     *IDidTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IDidRaw is an auto generated low-level Go binding around an Ethereum contract.
type IDidRaw struct {
	Contract *IDid // Generic contract binding to access the raw methods on
}

// IDidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IDidCallerRaw struct {
	Contract *IDidCaller // Generic read-only contract binding to access the raw methods on
}

// IDidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IDidTransactorRaw struct {
	Contract *IDidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIDid creates a new instance of IDid, bound to a specific deployed contract.
func NewIDid(address common.Address, backend bind.ContractBackend) (*IDid, error) {
	contract, err := bindIDid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IDid{IDidCaller: IDidCaller{contract: contract}, IDidTransactor: IDidTransactor{contract: contract}, IDidFilterer: IDidFilterer{contract: contract}}, nil
}

// NewIDidCaller creates a new read-only instance of IDid, bound to a specific deployed contract.
func NewIDidCaller(address common.Address, caller bind.ContractCaller) (*IDidCaller, error) {
	contract, err := bindIDid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IDidCaller{contract: contract}, nil
}

// NewIDidTransactor creates a new write-only instance of IDid, bound to a specific deployed contract.
func NewIDidTransactor(address common.Address, transactor bind.ContractTransactor) (*IDidTransactor, error) {
	contract, err := bindIDid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IDidTransactor{contract: contract}, nil
}

// NewIDidFilterer creates a new log filterer instance of IDid, bound to a specific deployed contract.
func NewIDidFilterer(address common.Address, filterer bind.ContractFilterer) (*IDidFilterer, error) {
	contract, err := bindIDid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IDidFilterer{contract: contract}, nil
}

// bindIDid binds a generic wrapper to an already deployed contract.
func bindIDid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IDidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDid *IDidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDid.Contract.IDidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDid *IDidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDid.Contract.IDidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDid *IDidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDid.Contract.IDidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IDid *IDidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IDid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IDid *IDidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IDid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IDid *IDidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IDid.Contract.contract.Transact(opts, method, params...)
}

// GetAllAuthKey is a free data retrieval call binding the contract method 0xd22eafd9.
//
// Solidity: function getAllAuthKey(string did) view returns((string,string,string[],bytes,bool,bool,uint256)[])
func (_IDid *IDidCaller) GetAllAuthKey(opts *bind.CallOpts, did string) ([]IDidPublicKey, error) {
	var out []interface{}
	err := _IDid.contract.Call(opts, &out, "getAllAuthKey", did)

	if err != nil {
		return *new([]IDidPublicKey), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDidPublicKey)).(*[]IDidPublicKey)

	return out0, err

}

// GetAllAuthKey is a free data retrieval call binding the contract method 0xd22eafd9.
//
// Solidity: function getAllAuthKey(string did) view returns((string,string,string[],bytes,bool,bool,uint256)[])
func (_IDid *IDidSession) GetAllAuthKey(did string) ([]IDidPublicKey, error) {
	return _IDid.Contract.GetAllAuthKey(&_IDid.CallOpts, did)
}

// GetAllAuthKey is a free data retrieval call binding the contract method 0xd22eafd9.
//
// Solidity: function getAllAuthKey(string did) view returns((string,string,string[],bytes,bool,bool,uint256)[])
func (_IDid *IDidCallerSession) GetAllAuthKey(did string) ([]IDidPublicKey, error) {
	return _IDid.Contract.GetAllAuthKey(&_IDid.CallOpts, did)
}

// GetAllController is a free data retrieval call binding the contract method 0x2e35d7ca.
//
// Solidity: function getAllController(string did) view returns(string[])
func (_IDid *IDidCaller) GetAllController(opts *bind.CallOpts, did string) ([]string, error) {
	var out []interface{}
	err := _IDid.contract.Call(opts, &out, "getAllController", did)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllController is a free data retrieval call binding the contract method 0x2e35d7ca.
//
// Solidity: function getAllController(string did) view returns(string[])
func (_IDid *IDidSession) GetAllController(did string) ([]string, error) {
	return _IDid.Contract.GetAllController(&_IDid.CallOpts, did)
}

// GetAllController is a free data retrieval call binding the contract method 0x2e35d7ca.
//
// Solidity: function getAllController(string did) view returns(string[])
func (_IDid *IDidCallerSession) GetAllController(did string) ([]string, error) {
	return _IDid.Contract.GetAllController(&_IDid.CallOpts, did)
}

// GetAllPubKey is a free data retrieval call binding the contract method 0x1f7c7a9d.
//
// Solidity: function getAllPubKey(string did) view returns((string,string,string[],bytes,bool,bool,uint256)[])
func (_IDid *IDidCaller) GetAllPubKey(opts *bind.CallOpts, did string) ([]IDidPublicKey, error) {
	var out []interface{}
	err := _IDid.contract.Call(opts, &out, "getAllPubKey", did)

	if err != nil {
		return *new([]IDidPublicKey), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDidPublicKey)).(*[]IDidPublicKey)

	return out0, err

}

// GetAllPubKey is a free data retrieval call binding the contract method 0x1f7c7a9d.
//
// Solidity: function getAllPubKey(string did) view returns((string,string,string[],bytes,bool,bool,uint256)[])
func (_IDid *IDidSession) GetAllPubKey(did string) ([]IDidPublicKey, error) {
	return _IDid.Contract.GetAllPubKey(&_IDid.CallOpts, did)
}

// GetAllPubKey is a free data retrieval call binding the contract method 0x1f7c7a9d.
//
// Solidity: function getAllPubKey(string did) view returns((string,string,string[],bytes,bool,bool,uint256)[])
func (_IDid *IDidCallerSession) GetAllPubKey(did string) ([]IDidPublicKey, error) {
	return _IDid.Contract.GetAllPubKey(&_IDid.CallOpts, did)
}

// GetAllService is a free data retrieval call binding the contract method 0xf8d8f4ad.
//
// Solidity: function getAllService(string did) view returns((string,string,string)[])
func (_IDid *IDidCaller) GetAllService(opts *bind.CallOpts, did string) ([]IDidService, error) {
	var out []interface{}
	err := _IDid.contract.Call(opts, &out, "getAllService", did)

	if err != nil {
		return *new([]IDidService), err
	}

	out0 := *abi.ConvertType(out[0], new([]IDidService)).(*[]IDidService)

	return out0, err

}

// GetAllService is a free data retrieval call binding the contract method 0xf8d8f4ad.
//
// Solidity: function getAllService(string did) view returns((string,string,string)[])
func (_IDid *IDidSession) GetAllService(did string) ([]IDidService, error) {
	return _IDid.Contract.GetAllService(&_IDid.CallOpts, did)
}

// GetAllService is a free data retrieval call binding the contract method 0xf8d8f4ad.
//
// Solidity: function getAllService(string did) view returns((string,string,string)[])
func (_IDid *IDidCallerSession) GetAllService(did string) ([]IDidService, error) {
	return _IDid.Contract.GetAllService(&_IDid.CallOpts, did)
}

// GetContext is a free data retrieval call binding the contract method 0x53949d4f.
//
// Solidity: function getContext(string did) view returns(string[])
func (_IDid *IDidCaller) GetContext(opts *bind.CallOpts, did string) ([]string, error) {
	var out []interface{}
	err := _IDid.contract.Call(opts, &out, "getContext", did)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetContext is a free data retrieval call binding the contract method 0x53949d4f.
//
// Solidity: function getContext(string did) view returns(string[])
func (_IDid *IDidSession) GetContext(did string) ([]string, error) {
	return _IDid.Contract.GetContext(&_IDid.CallOpts, did)
}

// GetContext is a free data retrieval call binding the contract method 0x53949d4f.
//
// Solidity: function getContext(string did) view returns(string[])
func (_IDid *IDidCallerSession) GetContext(did string) ([]string, error) {
	return _IDid.Contract.GetContext(&_IDid.CallOpts, did)
}

// GetDocument is a free data retrieval call binding the contract method 0x7ccb6a64.
//
// Solidity: function getDocument(string did) view returns((string[],string,(string,string,string[],bytes,bool,bool,uint256)[],(string,string,string[],bytes,bool,bool,uint256)[],string[],(string,string,string)[],uint256))
func (_IDid *IDidCaller) GetDocument(opts *bind.CallOpts, did string) (IDidDIDDocument, error) {
	var out []interface{}
	err := _IDid.contract.Call(opts, &out, "getDocument", did)

	if err != nil {
		return *new(IDidDIDDocument), err
	}

	out0 := *abi.ConvertType(out[0], new(IDidDIDDocument)).(*IDidDIDDocument)

	return out0, err

}

// GetDocument is a free data retrieval call binding the contract method 0x7ccb6a64.
//
// Solidity: function getDocument(string did) view returns((string[],string,(string,string,string[],bytes,bool,bool,uint256)[],(string,string,string[],bytes,bool,bool,uint256)[],string[],(string,string,string)[],uint256))
func (_IDid *IDidSession) GetDocument(did string) (IDidDIDDocument, error) {
	return _IDid.Contract.GetDocument(&_IDid.CallOpts, did)
}

// GetDocument is a free data retrieval call binding the contract method 0x7ccb6a64.
//
// Solidity: function getDocument(string did) view returns((string[],string,(string,string,string[],bytes,bool,bool,uint256)[],(string,string,string[],bytes,bool,bool,uint256)[],string[],(string,string,string)[],uint256))
func (_IDid *IDidCallerSession) GetDocument(did string) (IDidDIDDocument, error) {
	return _IDid.Contract.GetDocument(&_IDid.CallOpts, did)
}

// GetUpdatedTime is a free data retrieval call binding the contract method 0xba964281.
//
// Solidity: function getUpdatedTime(string did) view returns(uint256)
func (_IDid *IDidCaller) GetUpdatedTime(opts *bind.CallOpts, did string) (*big.Int, error) {
	var out []interface{}
	err := _IDid.contract.Call(opts, &out, "getUpdatedTime", did)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUpdatedTime is a free data retrieval call binding the contract method 0xba964281.
//
// Solidity: function getUpdatedTime(string did) view returns(uint256)
func (_IDid *IDidSession) GetUpdatedTime(did string) (*big.Int, error) {
	return _IDid.Contract.GetUpdatedTime(&_IDid.CallOpts, did)
}

// GetUpdatedTime is a free data retrieval call binding the contract method 0xba964281.
//
// Solidity: function getUpdatedTime(string did) view returns(uint256)
func (_IDid *IDidCallerSession) GetUpdatedTime(did string) (*big.Int, error) {
	return _IDid.Contract.GetUpdatedTime(&_IDid.CallOpts, did)
}

// AddAddr is a paid mutator transaction binding the contract method 0x53864503.
//
// Solidity: function addAddr(string did, address addr, string[] pubKeyController, bytes singer) returns()
func (_IDid *IDidTransactor) AddAddr(opts *bind.TransactOpts, did string, addr common.Address, pubKeyController []string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "addAddr", did, addr, pubKeyController, singer)
}

// AddAddr is a paid mutator transaction binding the contract method 0x53864503.
//
// Solidity: function addAddr(string did, address addr, string[] pubKeyController, bytes singer) returns()
func (_IDid *IDidSession) AddAddr(did string, addr common.Address, pubKeyController []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddAddr(&_IDid.TransactOpts, did, addr, pubKeyController, singer)
}

// AddAddr is a paid mutator transaction binding the contract method 0x53864503.
//
// Solidity: function addAddr(string did, address addr, string[] pubKeyController, bytes singer) returns()
func (_IDid *IDidTransactorSession) AddAddr(did string, addr common.Address, pubKeyController []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddAddr(&_IDid.TransactOpts, did, addr, pubKeyController, singer)
}

// AddContext is a paid mutator transaction binding the contract method 0x65cbf7c3.
//
// Solidity: function addContext(string did, string[] context, bytes singer) returns()
func (_IDid *IDidTransactor) AddContext(opts *bind.TransactOpts, did string, context []string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "addContext", did, context, singer)
}

// AddContext is a paid mutator transaction binding the contract method 0x65cbf7c3.
//
// Solidity: function addContext(string did, string[] context, bytes singer) returns()
func (_IDid *IDidSession) AddContext(did string, context []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddContext(&_IDid.TransactOpts, did, context, singer)
}

// AddContext is a paid mutator transaction binding the contract method 0x65cbf7c3.
//
// Solidity: function addContext(string did, string[] context, bytes singer) returns()
func (_IDid *IDidTransactorSession) AddContext(did string, context []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddContext(&_IDid.TransactOpts, did, context, singer)
}

// AddController is a paid mutator transaction binding the contract method 0x0ec0865b.
//
// Solidity: function addController(string did, string controller, bytes singer) returns()
func (_IDid *IDidTransactor) AddController(opts *bind.TransactOpts, did string, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "addController", did, controller, singer)
}

// AddController is a paid mutator transaction binding the contract method 0x0ec0865b.
//
// Solidity: function addController(string did, string controller, bytes singer) returns()
func (_IDid *IDidSession) AddController(did string, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddController(&_IDid.TransactOpts, did, controller, singer)
}

// AddController is a paid mutator transaction binding the contract method 0x0ec0865b.
//
// Solidity: function addController(string did, string controller, bytes singer) returns()
func (_IDid *IDidTransactorSession) AddController(did string, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddController(&_IDid.TransactOpts, did, controller, singer)
}

// AddKey is a paid mutator transaction binding the contract method 0x52c07c56.
//
// Solidity: function addKey(string did, bytes newPubKey, string[] pubKeyController, bytes singer) returns()
func (_IDid *IDidTransactor) AddKey(opts *bind.TransactOpts, did string, newPubKey []byte, pubKeyController []string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "addKey", did, newPubKey, pubKeyController, singer)
}

// AddKey is a paid mutator transaction binding the contract method 0x52c07c56.
//
// Solidity: function addKey(string did, bytes newPubKey, string[] pubKeyController, bytes singer) returns()
func (_IDid *IDidSession) AddKey(did string, newPubKey []byte, pubKeyController []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddKey(&_IDid.TransactOpts, did, newPubKey, pubKeyController, singer)
}

// AddKey is a paid mutator transaction binding the contract method 0x52c07c56.
//
// Solidity: function addKey(string did, bytes newPubKey, string[] pubKeyController, bytes singer) returns()
func (_IDid *IDidTransactorSession) AddKey(did string, newPubKey []byte, pubKeyController []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddKey(&_IDid.TransactOpts, did, newPubKey, pubKeyController, singer)
}

// AddNewAuthAddr is a paid mutator transaction binding the contract method 0x8da5588c.
//
// Solidity: function addNewAuthAddr(string did, address addr, string[] controller, bytes singer) returns()
func (_IDid *IDidTransactor) AddNewAuthAddr(opts *bind.TransactOpts, did string, addr common.Address, controller []string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "addNewAuthAddr", did, addr, controller, singer)
}

// AddNewAuthAddr is a paid mutator transaction binding the contract method 0x8da5588c.
//
// Solidity: function addNewAuthAddr(string did, address addr, string[] controller, bytes singer) returns()
func (_IDid *IDidSession) AddNewAuthAddr(did string, addr common.Address, controller []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddNewAuthAddr(&_IDid.TransactOpts, did, addr, controller, singer)
}

// AddNewAuthAddr is a paid mutator transaction binding the contract method 0x8da5588c.
//
// Solidity: function addNewAuthAddr(string did, address addr, string[] controller, bytes singer) returns()
func (_IDid *IDidTransactorSession) AddNewAuthAddr(did string, addr common.Address, controller []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddNewAuthAddr(&_IDid.TransactOpts, did, addr, controller, singer)
}

// AddNewAuthAddrByController is a paid mutator transaction binding the contract method 0xb0abb199.
//
// Solidity: function addNewAuthAddrByController(string did, address addr, string[] controller, string controllerSigner, bytes singer) returns()
func (_IDid *IDidTransactor) AddNewAuthAddrByController(opts *bind.TransactOpts, did string, addr common.Address, controller []string, controllerSigner string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "addNewAuthAddrByController", did, addr, controller, controllerSigner, singer)
}

// AddNewAuthAddrByController is a paid mutator transaction binding the contract method 0xb0abb199.
//
// Solidity: function addNewAuthAddrByController(string did, address addr, string[] controller, string controllerSigner, bytes singer) returns()
func (_IDid *IDidSession) AddNewAuthAddrByController(did string, addr common.Address, controller []string, controllerSigner string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddNewAuthAddrByController(&_IDid.TransactOpts, did, addr, controller, controllerSigner, singer)
}

// AddNewAuthAddrByController is a paid mutator transaction binding the contract method 0xb0abb199.
//
// Solidity: function addNewAuthAddrByController(string did, address addr, string[] controller, string controllerSigner, bytes singer) returns()
func (_IDid *IDidTransactorSession) AddNewAuthAddrByController(did string, addr common.Address, controller []string, controllerSigner string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddNewAuthAddrByController(&_IDid.TransactOpts, did, addr, controller, controllerSigner, singer)
}

// AddNewAuthKey is a paid mutator transaction binding the contract method 0x002b4c82.
//
// Solidity: function addNewAuthKey(string did, bytes pubKey, string[] controller, bytes singer) returns()
func (_IDid *IDidTransactor) AddNewAuthKey(opts *bind.TransactOpts, did string, pubKey []byte, controller []string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "addNewAuthKey", did, pubKey, controller, singer)
}

// AddNewAuthKey is a paid mutator transaction binding the contract method 0x002b4c82.
//
// Solidity: function addNewAuthKey(string did, bytes pubKey, string[] controller, bytes singer) returns()
func (_IDid *IDidSession) AddNewAuthKey(did string, pubKey []byte, controller []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddNewAuthKey(&_IDid.TransactOpts, did, pubKey, controller, singer)
}

// AddNewAuthKey is a paid mutator transaction binding the contract method 0x002b4c82.
//
// Solidity: function addNewAuthKey(string did, bytes pubKey, string[] controller, bytes singer) returns()
func (_IDid *IDidTransactorSession) AddNewAuthKey(did string, pubKey []byte, controller []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddNewAuthKey(&_IDid.TransactOpts, did, pubKey, controller, singer)
}

// AddNewAuthKeyByController is a paid mutator transaction binding the contract method 0xdb7d1f04.
//
// Solidity: function addNewAuthKeyByController(string did, bytes pubKey, string[] controller, string controllerSigner, bytes singer) returns()
func (_IDid *IDidTransactor) AddNewAuthKeyByController(opts *bind.TransactOpts, did string, pubKey []byte, controller []string, controllerSigner string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "addNewAuthKeyByController", did, pubKey, controller, controllerSigner, singer)
}

// AddNewAuthKeyByController is a paid mutator transaction binding the contract method 0xdb7d1f04.
//
// Solidity: function addNewAuthKeyByController(string did, bytes pubKey, string[] controller, string controllerSigner, bytes singer) returns()
func (_IDid *IDidSession) AddNewAuthKeyByController(did string, pubKey []byte, controller []string, controllerSigner string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddNewAuthKeyByController(&_IDid.TransactOpts, did, pubKey, controller, controllerSigner, singer)
}

// AddNewAuthKeyByController is a paid mutator transaction binding the contract method 0xdb7d1f04.
//
// Solidity: function addNewAuthKeyByController(string did, bytes pubKey, string[] controller, string controllerSigner, bytes singer) returns()
func (_IDid *IDidTransactorSession) AddNewAuthKeyByController(did string, pubKey []byte, controller []string, controllerSigner string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddNewAuthKeyByController(&_IDid.TransactOpts, did, pubKey, controller, controllerSigner, singer)
}

// AddService is a paid mutator transaction binding the contract method 0x7d1c754f.
//
// Solidity: function addService(string did, string serviceId, string serviceType, string serviceEndpoint, bytes singer) returns()
func (_IDid *IDidTransactor) AddService(opts *bind.TransactOpts, did string, serviceId string, serviceType string, serviceEndpoint string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "addService", did, serviceId, serviceType, serviceEndpoint, singer)
}

// AddService is a paid mutator transaction binding the contract method 0x7d1c754f.
//
// Solidity: function addService(string did, string serviceId, string serviceType, string serviceEndpoint, bytes singer) returns()
func (_IDid *IDidSession) AddService(did string, serviceId string, serviceType string, serviceEndpoint string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddService(&_IDid.TransactOpts, did, serviceId, serviceType, serviceEndpoint, singer)
}

// AddService is a paid mutator transaction binding the contract method 0x7d1c754f.
//
// Solidity: function addService(string did, string serviceId, string serviceType, string serviceEndpoint, bytes singer) returns()
func (_IDid *IDidTransactorSession) AddService(did string, serviceId string, serviceType string, serviceEndpoint string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.AddService(&_IDid.TransactOpts, did, serviceId, serviceType, serviceEndpoint, singer)
}

// DeactivateAddr is a paid mutator transaction binding the contract method 0xd558cdea.
//
// Solidity: function deactivateAddr(string did, address addr, bytes singer) returns()
func (_IDid *IDidTransactor) DeactivateAddr(opts *bind.TransactOpts, did string, addr common.Address, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "deactivateAddr", did, addr, singer)
}

// DeactivateAddr is a paid mutator transaction binding the contract method 0xd558cdea.
//
// Solidity: function deactivateAddr(string did, address addr, bytes singer) returns()
func (_IDid *IDidSession) DeactivateAddr(did string, addr common.Address, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateAddr(&_IDid.TransactOpts, did, addr, singer)
}

// DeactivateAddr is a paid mutator transaction binding the contract method 0xd558cdea.
//
// Solidity: function deactivateAddr(string did, address addr, bytes singer) returns()
func (_IDid *IDidTransactorSession) DeactivateAddr(did string, addr common.Address, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateAddr(&_IDid.TransactOpts, did, addr, singer)
}

// DeactivateAuthAddr is a paid mutator transaction binding the contract method 0x6747e9bc.
//
// Solidity: function deactivateAuthAddr(string did, address addr, bytes singer) returns()
func (_IDid *IDidTransactor) DeactivateAuthAddr(opts *bind.TransactOpts, did string, addr common.Address, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "deactivateAuthAddr", did, addr, singer)
}

// DeactivateAuthAddr is a paid mutator transaction binding the contract method 0x6747e9bc.
//
// Solidity: function deactivateAuthAddr(string did, address addr, bytes singer) returns()
func (_IDid *IDidSession) DeactivateAuthAddr(did string, addr common.Address, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateAuthAddr(&_IDid.TransactOpts, did, addr, singer)
}

// DeactivateAuthAddr is a paid mutator transaction binding the contract method 0x6747e9bc.
//
// Solidity: function deactivateAuthAddr(string did, address addr, bytes singer) returns()
func (_IDid *IDidTransactorSession) DeactivateAuthAddr(did string, addr common.Address, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateAuthAddr(&_IDid.TransactOpts, did, addr, singer)
}

// DeactivateAuthAddrByController is a paid mutator transaction binding the contract method 0x88209cf3.
//
// Solidity: function deactivateAuthAddrByController(string did, address addr, string controller, bytes singer) returns()
func (_IDid *IDidTransactor) DeactivateAuthAddrByController(opts *bind.TransactOpts, did string, addr common.Address, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "deactivateAuthAddrByController", did, addr, controller, singer)
}

// DeactivateAuthAddrByController is a paid mutator transaction binding the contract method 0x88209cf3.
//
// Solidity: function deactivateAuthAddrByController(string did, address addr, string controller, bytes singer) returns()
func (_IDid *IDidSession) DeactivateAuthAddrByController(did string, addr common.Address, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateAuthAddrByController(&_IDid.TransactOpts, did, addr, controller, singer)
}

// DeactivateAuthAddrByController is a paid mutator transaction binding the contract method 0x88209cf3.
//
// Solidity: function deactivateAuthAddrByController(string did, address addr, string controller, bytes singer) returns()
func (_IDid *IDidTransactorSession) DeactivateAuthAddrByController(did string, addr common.Address, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateAuthAddrByController(&_IDid.TransactOpts, did, addr, controller, singer)
}

// DeactivateAuthKey is a paid mutator transaction binding the contract method 0x958607b8.
//
// Solidity: function deactivateAuthKey(string did, bytes pubKey, bytes singer) returns()
func (_IDid *IDidTransactor) DeactivateAuthKey(opts *bind.TransactOpts, did string, pubKey []byte, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "deactivateAuthKey", did, pubKey, singer)
}

// DeactivateAuthKey is a paid mutator transaction binding the contract method 0x958607b8.
//
// Solidity: function deactivateAuthKey(string did, bytes pubKey, bytes singer) returns()
func (_IDid *IDidSession) DeactivateAuthKey(did string, pubKey []byte, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateAuthKey(&_IDid.TransactOpts, did, pubKey, singer)
}

// DeactivateAuthKey is a paid mutator transaction binding the contract method 0x958607b8.
//
// Solidity: function deactivateAuthKey(string did, bytes pubKey, bytes singer) returns()
func (_IDid *IDidTransactorSession) DeactivateAuthKey(did string, pubKey []byte, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateAuthKey(&_IDid.TransactOpts, did, pubKey, singer)
}

// DeactivateAuthKeyByController is a paid mutator transaction binding the contract method 0x9053ad03.
//
// Solidity: function deactivateAuthKeyByController(string did, bytes pubKey, string controller, bytes singer) returns()
func (_IDid *IDidTransactor) DeactivateAuthKeyByController(opts *bind.TransactOpts, did string, pubKey []byte, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "deactivateAuthKeyByController", did, pubKey, controller, singer)
}

// DeactivateAuthKeyByController is a paid mutator transaction binding the contract method 0x9053ad03.
//
// Solidity: function deactivateAuthKeyByController(string did, bytes pubKey, string controller, bytes singer) returns()
func (_IDid *IDidSession) DeactivateAuthKeyByController(did string, pubKey []byte, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateAuthKeyByController(&_IDid.TransactOpts, did, pubKey, controller, singer)
}

// DeactivateAuthKeyByController is a paid mutator transaction binding the contract method 0x9053ad03.
//
// Solidity: function deactivateAuthKeyByController(string did, bytes pubKey, string controller, bytes singer) returns()
func (_IDid *IDidTransactorSession) DeactivateAuthKeyByController(did string, pubKey []byte, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateAuthKeyByController(&_IDid.TransactOpts, did, pubKey, controller, singer)
}

// DeactivateID is a paid mutator transaction binding the contract method 0xec02a3b5.
//
// Solidity: function deactivateID(string did, bytes singer) returns()
func (_IDid *IDidTransactor) DeactivateID(opts *bind.TransactOpts, did string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "deactivateID", did, singer)
}

// DeactivateID is a paid mutator transaction binding the contract method 0xec02a3b5.
//
// Solidity: function deactivateID(string did, bytes singer) returns()
func (_IDid *IDidSession) DeactivateID(did string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateID(&_IDid.TransactOpts, did, singer)
}

// DeactivateID is a paid mutator transaction binding the contract method 0xec02a3b5.
//
// Solidity: function deactivateID(string did, bytes singer) returns()
func (_IDid *IDidTransactorSession) DeactivateID(did string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateID(&_IDid.TransactOpts, did, singer)
}

// DeactivateKey is a paid mutator transaction binding the contract method 0x793f66c9.
//
// Solidity: function deactivateKey(string did, bytes pubKey, bytes singer) returns()
func (_IDid *IDidTransactor) DeactivateKey(opts *bind.TransactOpts, did string, pubKey []byte, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "deactivateKey", did, pubKey, singer)
}

// DeactivateKey is a paid mutator transaction binding the contract method 0x793f66c9.
//
// Solidity: function deactivateKey(string did, bytes pubKey, bytes singer) returns()
func (_IDid *IDidSession) DeactivateKey(did string, pubKey []byte, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateKey(&_IDid.TransactOpts, did, pubKey, singer)
}

// DeactivateKey is a paid mutator transaction binding the contract method 0x793f66c9.
//
// Solidity: function deactivateKey(string did, bytes pubKey, bytes singer) returns()
func (_IDid *IDidTransactorSession) DeactivateKey(did string, pubKey []byte, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.DeactivateKey(&_IDid.TransactOpts, did, pubKey, singer)
}

// RemoveContext is a paid mutator transaction binding the contract method 0x8095627f.
//
// Solidity: function removeContext(string did, string[] context, bytes singer) returns()
func (_IDid *IDidTransactor) RemoveContext(opts *bind.TransactOpts, did string, context []string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "removeContext", did, context, singer)
}

// RemoveContext is a paid mutator transaction binding the contract method 0x8095627f.
//
// Solidity: function removeContext(string did, string[] context, bytes singer) returns()
func (_IDid *IDidSession) RemoveContext(did string, context []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.RemoveContext(&_IDid.TransactOpts, did, context, singer)
}

// RemoveContext is a paid mutator transaction binding the contract method 0x8095627f.
//
// Solidity: function removeContext(string did, string[] context, bytes singer) returns()
func (_IDid *IDidTransactorSession) RemoveContext(did string, context []string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.RemoveContext(&_IDid.TransactOpts, did, context, singer)
}

// RemoveController is a paid mutator transaction binding the contract method 0x244d9153.
//
// Solidity: function removeController(string did, string controller, bytes singer) returns()
func (_IDid *IDidTransactor) RemoveController(opts *bind.TransactOpts, did string, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "removeController", did, controller, singer)
}

// RemoveController is a paid mutator transaction binding the contract method 0x244d9153.
//
// Solidity: function removeController(string did, string controller, bytes singer) returns()
func (_IDid *IDidSession) RemoveController(did string, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.RemoveController(&_IDid.TransactOpts, did, controller, singer)
}

// RemoveController is a paid mutator transaction binding the contract method 0x244d9153.
//
// Solidity: function removeController(string did, string controller, bytes singer) returns()
func (_IDid *IDidTransactorSession) RemoveController(did string, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.RemoveController(&_IDid.TransactOpts, did, controller, singer)
}

// RemoveService is a paid mutator transaction binding the contract method 0x2e97b8d4.
//
// Solidity: function removeService(string did, string serviceId, bytes singer) returns()
func (_IDid *IDidTransactor) RemoveService(opts *bind.TransactOpts, did string, serviceId string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "removeService", did, serviceId, singer)
}

// RemoveService is a paid mutator transaction binding the contract method 0x2e97b8d4.
//
// Solidity: function removeService(string did, string serviceId, bytes singer) returns()
func (_IDid *IDidSession) RemoveService(did string, serviceId string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.RemoveService(&_IDid.TransactOpts, did, serviceId, singer)
}

// RemoveService is a paid mutator transaction binding the contract method 0x2e97b8d4.
//
// Solidity: function removeService(string did, string serviceId, bytes singer) returns()
func (_IDid *IDidTransactorSession) RemoveService(did string, serviceId string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.RemoveService(&_IDid.TransactOpts, did, serviceId, singer)
}

// SetAuthAddr is a paid mutator transaction binding the contract method 0x398a85ec.
//
// Solidity: function setAuthAddr(string did, address addr, bytes singer) returns()
func (_IDid *IDidTransactor) SetAuthAddr(opts *bind.TransactOpts, did string, addr common.Address, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "setAuthAddr", did, addr, singer)
}

// SetAuthAddr is a paid mutator transaction binding the contract method 0x398a85ec.
//
// Solidity: function setAuthAddr(string did, address addr, bytes singer) returns()
func (_IDid *IDidSession) SetAuthAddr(did string, addr common.Address, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.SetAuthAddr(&_IDid.TransactOpts, did, addr, singer)
}

// SetAuthAddr is a paid mutator transaction binding the contract method 0x398a85ec.
//
// Solidity: function setAuthAddr(string did, address addr, bytes singer) returns()
func (_IDid *IDidTransactorSession) SetAuthAddr(did string, addr common.Address, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.SetAuthAddr(&_IDid.TransactOpts, did, addr, singer)
}

// SetAuthAddrByController is a paid mutator transaction binding the contract method 0x6e5f3183.
//
// Solidity: function setAuthAddrByController(string did, address addr, string controller, bytes singer) returns()
func (_IDid *IDidTransactor) SetAuthAddrByController(opts *bind.TransactOpts, did string, addr common.Address, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "setAuthAddrByController", did, addr, controller, singer)
}

// SetAuthAddrByController is a paid mutator transaction binding the contract method 0x6e5f3183.
//
// Solidity: function setAuthAddrByController(string did, address addr, string controller, bytes singer) returns()
func (_IDid *IDidSession) SetAuthAddrByController(did string, addr common.Address, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.SetAuthAddrByController(&_IDid.TransactOpts, did, addr, controller, singer)
}

// SetAuthAddrByController is a paid mutator transaction binding the contract method 0x6e5f3183.
//
// Solidity: function setAuthAddrByController(string did, address addr, string controller, bytes singer) returns()
func (_IDid *IDidTransactorSession) SetAuthAddrByController(did string, addr common.Address, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.SetAuthAddrByController(&_IDid.TransactOpts, did, addr, controller, singer)
}

// SetAuthKey is a paid mutator transaction binding the contract method 0xfa1bac72.
//
// Solidity: function setAuthKey(string did, bytes pubKey, bytes singer) returns()
func (_IDid *IDidTransactor) SetAuthKey(opts *bind.TransactOpts, did string, pubKey []byte, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "setAuthKey", did, pubKey, singer)
}

// SetAuthKey is a paid mutator transaction binding the contract method 0xfa1bac72.
//
// Solidity: function setAuthKey(string did, bytes pubKey, bytes singer) returns()
func (_IDid *IDidSession) SetAuthKey(did string, pubKey []byte, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.SetAuthKey(&_IDid.TransactOpts, did, pubKey, singer)
}

// SetAuthKey is a paid mutator transaction binding the contract method 0xfa1bac72.
//
// Solidity: function setAuthKey(string did, bytes pubKey, bytes singer) returns()
func (_IDid *IDidTransactorSession) SetAuthKey(did string, pubKey []byte, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.SetAuthKey(&_IDid.TransactOpts, did, pubKey, singer)
}

// SetAuthKeyByController is a paid mutator transaction binding the contract method 0x9033d99a.
//
// Solidity: function setAuthKeyByController(string did, bytes pubKey, string controller, bytes singer) returns()
func (_IDid *IDidTransactor) SetAuthKeyByController(opts *bind.TransactOpts, did string, pubKey []byte, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "setAuthKeyByController", did, pubKey, controller, singer)
}

// SetAuthKeyByController is a paid mutator transaction binding the contract method 0x9033d99a.
//
// Solidity: function setAuthKeyByController(string did, bytes pubKey, string controller, bytes singer) returns()
func (_IDid *IDidSession) SetAuthKeyByController(did string, pubKey []byte, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.SetAuthKeyByController(&_IDid.TransactOpts, did, pubKey, controller, singer)
}

// SetAuthKeyByController is a paid mutator transaction binding the contract method 0x9033d99a.
//
// Solidity: function setAuthKeyByController(string did, bytes pubKey, string controller, bytes singer) returns()
func (_IDid *IDidTransactorSession) SetAuthKeyByController(did string, pubKey []byte, controller string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.SetAuthKeyByController(&_IDid.TransactOpts, did, pubKey, controller, singer)
}

// UpdateService is a paid mutator transaction binding the contract method 0x667f5fa3.
//
// Solidity: function updateService(string did, string serviceId, string serviceType, string serviceEndpoint, bytes singer) returns()
func (_IDid *IDidTransactor) UpdateService(opts *bind.TransactOpts, did string, serviceId string, serviceType string, serviceEndpoint string, singer []byte) (*types.Transaction, error) {
	return _IDid.contract.Transact(opts, "updateService", did, serviceId, serviceType, serviceEndpoint, singer)
}

// UpdateService is a paid mutator transaction binding the contract method 0x667f5fa3.
//
// Solidity: function updateService(string did, string serviceId, string serviceType, string serviceEndpoint, bytes singer) returns()
func (_IDid *IDidSession) UpdateService(did string, serviceId string, serviceType string, serviceEndpoint string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.UpdateService(&_IDid.TransactOpts, did, serviceId, serviceType, serviceEndpoint, singer)
}

// UpdateService is a paid mutator transaction binding the contract method 0x667f5fa3.
//
// Solidity: function updateService(string did, string serviceId, string serviceType, string serviceEndpoint, bytes singer) returns()
func (_IDid *IDidTransactorSession) UpdateService(did string, serviceId string, serviceType string, serviceEndpoint string, singer []byte) (*types.Transaction, error) {
	return _IDid.Contract.UpdateService(&_IDid.TransactOpts, did, serviceId, serviceType, serviceEndpoint, singer)
}

// IDidAddAddrIterator is returned from FilterAddAddr and is used to iterate over the raw logs and unpacked data for AddAddr events raised by the IDid contract.
type IDidAddAddrIterator struct {
	Event *IDidAddAddr // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidAddAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidAddAddr)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidAddAddr)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidAddAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidAddAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidAddAddr represents a AddAddr event raised by the IDid contract.
type IDidAddAddr struct {
	Did        string
	Addr       common.Address
	Controller []string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddAddr is a free log retrieval operation binding the contract event 0xa9f8cf46beb11b19df1f3f276b45158a279eff65c1fb27bb4f3b91e1a894f0dd.
//
// Solidity: event AddAddr(string did, address addr, string[] controller)
func (_IDid *IDidFilterer) FilterAddAddr(opts *bind.FilterOpts) (*IDidAddAddrIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "AddAddr")
	if err != nil {
		return nil, err
	}
	return &IDidAddAddrIterator{contract: _IDid.contract, event: "AddAddr", logs: logs, sub: sub}, nil
}

// WatchAddAddr is a free log subscription operation binding the contract event 0xa9f8cf46beb11b19df1f3f276b45158a279eff65c1fb27bb4f3b91e1a894f0dd.
//
// Solidity: event AddAddr(string did, address addr, string[] controller)
func (_IDid *IDidFilterer) WatchAddAddr(opts *bind.WatchOpts, sink chan<- *IDidAddAddr) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "AddAddr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidAddAddr)
				if err := _IDid.contract.UnpackLog(event, "AddAddr", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddAddr is a log parse operation binding the contract event 0xa9f8cf46beb11b19df1f3f276b45158a279eff65c1fb27bb4f3b91e1a894f0dd.
//
// Solidity: event AddAddr(string did, address addr, string[] controller)
func (_IDid *IDidFilterer) ParseAddAddr(log types.Log) (*IDidAddAddr, error) {
	event := new(IDidAddAddr)
	if err := _IDid.contract.UnpackLog(event, "AddAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidAddContextIterator is returned from FilterAddContext and is used to iterate over the raw logs and unpacked data for AddContext events raised by the IDid contract.
type IDidAddContextIterator struct {
	Event *IDidAddContext // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidAddContextIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidAddContext)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidAddContext)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidAddContextIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidAddContextIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidAddContext represents a AddContext event raised by the IDid contract.
type IDidAddContext struct {
	Did     string
	Context string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddContext is a free log retrieval operation binding the contract event 0xab878ced0a294d7dde0c095c50e69cd0c400119ef2d9545bda4cc63d281148c5.
//
// Solidity: event AddContext(string did, string context)
func (_IDid *IDidFilterer) FilterAddContext(opts *bind.FilterOpts) (*IDidAddContextIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "AddContext")
	if err != nil {
		return nil, err
	}
	return &IDidAddContextIterator{contract: _IDid.contract, event: "AddContext", logs: logs, sub: sub}, nil
}

// WatchAddContext is a free log subscription operation binding the contract event 0xab878ced0a294d7dde0c095c50e69cd0c400119ef2d9545bda4cc63d281148c5.
//
// Solidity: event AddContext(string did, string context)
func (_IDid *IDidFilterer) WatchAddContext(opts *bind.WatchOpts, sink chan<- *IDidAddContext) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "AddContext")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidAddContext)
				if err := _IDid.contract.UnpackLog(event, "AddContext", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddContext is a log parse operation binding the contract event 0xab878ced0a294d7dde0c095c50e69cd0c400119ef2d9545bda4cc63d281148c5.
//
// Solidity: event AddContext(string did, string context)
func (_IDid *IDidFilterer) ParseAddContext(log types.Log) (*IDidAddContext, error) {
	event := new(IDidAddContext)
	if err := _IDid.contract.UnpackLog(event, "AddContext", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidAddControllerIterator is returned from FilterAddController and is used to iterate over the raw logs and unpacked data for AddController events raised by the IDid contract.
type IDidAddControllerIterator struct {
	Event *IDidAddController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidAddControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidAddController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidAddController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidAddControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidAddControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidAddController represents a AddController event raised by the IDid contract.
type IDidAddController struct {
	Did        string
	Controller string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddController is a free log retrieval operation binding the contract event 0x158049e5f97a0923906423a422819a72f859b99834ab1471ec73c7e5d8c2221d.
//
// Solidity: event AddController(string did, string controller)
func (_IDid *IDidFilterer) FilterAddController(opts *bind.FilterOpts) (*IDidAddControllerIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "AddController")
	if err != nil {
		return nil, err
	}
	return &IDidAddControllerIterator{contract: _IDid.contract, event: "AddController", logs: logs, sub: sub}, nil
}

// WatchAddController is a free log subscription operation binding the contract event 0x158049e5f97a0923906423a422819a72f859b99834ab1471ec73c7e5d8c2221d.
//
// Solidity: event AddController(string did, string controller)
func (_IDid *IDidFilterer) WatchAddController(opts *bind.WatchOpts, sink chan<- *IDidAddController) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "AddController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidAddController)
				if err := _IDid.contract.UnpackLog(event, "AddController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddController is a log parse operation binding the contract event 0x158049e5f97a0923906423a422819a72f859b99834ab1471ec73c7e5d8c2221d.
//
// Solidity: event AddController(string did, string controller)
func (_IDid *IDidFilterer) ParseAddController(log types.Log) (*IDidAddController, error) {
	event := new(IDidAddController)
	if err := _IDid.contract.UnpackLog(event, "AddController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidAddKeyIterator is returned from FilterAddKey and is used to iterate over the raw logs and unpacked data for AddKey events raised by the IDid contract.
type IDidAddKeyIterator struct {
	Event *IDidAddKey // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidAddKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidAddKey)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidAddKey)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidAddKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidAddKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidAddKey represents a AddKey event raised by the IDid contract.
type IDidAddKey struct {
	Did        string
	PubKey     []byte
	Controller []string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddKey is a free log retrieval operation binding the contract event 0x0f649d2ebae9f0c1584c133d91aaae79f84b2580c33e7c2480a4bd0e27a29109.
//
// Solidity: event AddKey(string did, bytes pubKey, string[] controller)
func (_IDid *IDidFilterer) FilterAddKey(opts *bind.FilterOpts) (*IDidAddKeyIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "AddKey")
	if err != nil {
		return nil, err
	}
	return &IDidAddKeyIterator{contract: _IDid.contract, event: "AddKey", logs: logs, sub: sub}, nil
}

// WatchAddKey is a free log subscription operation binding the contract event 0x0f649d2ebae9f0c1584c133d91aaae79f84b2580c33e7c2480a4bd0e27a29109.
//
// Solidity: event AddKey(string did, bytes pubKey, string[] controller)
func (_IDid *IDidFilterer) WatchAddKey(opts *bind.WatchOpts, sink chan<- *IDidAddKey) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "AddKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidAddKey)
				if err := _IDid.contract.UnpackLog(event, "AddKey", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddKey is a log parse operation binding the contract event 0x0f649d2ebae9f0c1584c133d91aaae79f84b2580c33e7c2480a4bd0e27a29109.
//
// Solidity: event AddKey(string did, bytes pubKey, string[] controller)
func (_IDid *IDidFilterer) ParseAddKey(log types.Log) (*IDidAddKey, error) {
	event := new(IDidAddKey)
	if err := _IDid.contract.UnpackLog(event, "AddKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidAddNewAuthAddrIterator is returned from FilterAddNewAuthAddr and is used to iterate over the raw logs and unpacked data for AddNewAuthAddr events raised by the IDid contract.
type IDidAddNewAuthAddrIterator struct {
	Event *IDidAddNewAuthAddr // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidAddNewAuthAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidAddNewAuthAddr)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidAddNewAuthAddr)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidAddNewAuthAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidAddNewAuthAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidAddNewAuthAddr represents a AddNewAuthAddr event raised by the IDid contract.
type IDidAddNewAuthAddr struct {
	Did        string
	Addr       common.Address
	Controller []string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddNewAuthAddr is a free log retrieval operation binding the contract event 0xe2bd3ac8e89e6fc034167d63de32d1247aa49f6802f2264875d0a00fd645da54.
//
// Solidity: event AddNewAuthAddr(string did, address addr, string[] controller)
func (_IDid *IDidFilterer) FilterAddNewAuthAddr(opts *bind.FilterOpts) (*IDidAddNewAuthAddrIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "AddNewAuthAddr")
	if err != nil {
		return nil, err
	}
	return &IDidAddNewAuthAddrIterator{contract: _IDid.contract, event: "AddNewAuthAddr", logs: logs, sub: sub}, nil
}

// WatchAddNewAuthAddr is a free log subscription operation binding the contract event 0xe2bd3ac8e89e6fc034167d63de32d1247aa49f6802f2264875d0a00fd645da54.
//
// Solidity: event AddNewAuthAddr(string did, address addr, string[] controller)
func (_IDid *IDidFilterer) WatchAddNewAuthAddr(opts *bind.WatchOpts, sink chan<- *IDidAddNewAuthAddr) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "AddNewAuthAddr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidAddNewAuthAddr)
				if err := _IDid.contract.UnpackLog(event, "AddNewAuthAddr", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddNewAuthAddr is a log parse operation binding the contract event 0xe2bd3ac8e89e6fc034167d63de32d1247aa49f6802f2264875d0a00fd645da54.
//
// Solidity: event AddNewAuthAddr(string did, address addr, string[] controller)
func (_IDid *IDidFilterer) ParseAddNewAuthAddr(log types.Log) (*IDidAddNewAuthAddr, error) {
	event := new(IDidAddNewAuthAddr)
	if err := _IDid.contract.UnpackLog(event, "AddNewAuthAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidAddNewAuthKeyIterator is returned from FilterAddNewAuthKey and is used to iterate over the raw logs and unpacked data for AddNewAuthKey events raised by the IDid contract.
type IDidAddNewAuthKeyIterator struct {
	Event *IDidAddNewAuthKey // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidAddNewAuthKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidAddNewAuthKey)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidAddNewAuthKey)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidAddNewAuthKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidAddNewAuthKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidAddNewAuthKey represents a AddNewAuthKey event raised by the IDid contract.
type IDidAddNewAuthKey struct {
	Did        string
	PubKey     []byte
	Controller []string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddNewAuthKey is a free log retrieval operation binding the contract event 0xb22d64a17a74dc192d537930ccfefdcc25c0e279b7595081c4f42e7613f01cf8.
//
// Solidity: event AddNewAuthKey(string did, bytes pubKey, string[] controller)
func (_IDid *IDidFilterer) FilterAddNewAuthKey(opts *bind.FilterOpts) (*IDidAddNewAuthKeyIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "AddNewAuthKey")
	if err != nil {
		return nil, err
	}
	return &IDidAddNewAuthKeyIterator{contract: _IDid.contract, event: "AddNewAuthKey", logs: logs, sub: sub}, nil
}

// WatchAddNewAuthKey is a free log subscription operation binding the contract event 0xb22d64a17a74dc192d537930ccfefdcc25c0e279b7595081c4f42e7613f01cf8.
//
// Solidity: event AddNewAuthKey(string did, bytes pubKey, string[] controller)
func (_IDid *IDidFilterer) WatchAddNewAuthKey(opts *bind.WatchOpts, sink chan<- *IDidAddNewAuthKey) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "AddNewAuthKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidAddNewAuthKey)
				if err := _IDid.contract.UnpackLog(event, "AddNewAuthKey", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddNewAuthKey is a log parse operation binding the contract event 0xb22d64a17a74dc192d537930ccfefdcc25c0e279b7595081c4f42e7613f01cf8.
//
// Solidity: event AddNewAuthKey(string did, bytes pubKey, string[] controller)
func (_IDid *IDidFilterer) ParseAddNewAuthKey(log types.Log) (*IDidAddNewAuthKey, error) {
	event := new(IDidAddNewAuthKey)
	if err := _IDid.contract.UnpackLog(event, "AddNewAuthKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidAddServiceIterator is returned from FilterAddService and is used to iterate over the raw logs and unpacked data for AddService events raised by the IDid contract.
type IDidAddServiceIterator struct {
	Event *IDidAddService // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidAddServiceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidAddService)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidAddService)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidAddServiceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidAddServiceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidAddService represents a AddService event raised by the IDid contract.
type IDidAddService struct {
	Did             string
	ServiceId       string
	ServiceType     string
	ServiceEndpoint string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAddService is a free log retrieval operation binding the contract event 0x8e2e47c852a8c05bbd1bc96df47294a54af338a51f99884924137ab38278860f.
//
// Solidity: event AddService(string did, string serviceId, string serviceType, string serviceEndpoint)
func (_IDid *IDidFilterer) FilterAddService(opts *bind.FilterOpts) (*IDidAddServiceIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "AddService")
	if err != nil {
		return nil, err
	}
	return &IDidAddServiceIterator{contract: _IDid.contract, event: "AddService", logs: logs, sub: sub}, nil
}

// WatchAddService is a free log subscription operation binding the contract event 0x8e2e47c852a8c05bbd1bc96df47294a54af338a51f99884924137ab38278860f.
//
// Solidity: event AddService(string did, string serviceId, string serviceType, string serviceEndpoint)
func (_IDid *IDidFilterer) WatchAddService(opts *bind.WatchOpts, sink chan<- *IDidAddService) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "AddService")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidAddService)
				if err := _IDid.contract.UnpackLog(event, "AddService", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddService is a log parse operation binding the contract event 0x8e2e47c852a8c05bbd1bc96df47294a54af338a51f99884924137ab38278860f.
//
// Solidity: event AddService(string did, string serviceId, string serviceType, string serviceEndpoint)
func (_IDid *IDidFilterer) ParseAddService(log types.Log) (*IDidAddService, error) {
	event := new(IDidAddService)
	if err := _IDid.contract.UnpackLog(event, "AddService", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidDeactivateIterator is returned from FilterDeactivate and is used to iterate over the raw logs and unpacked data for Deactivate events raised by the IDid contract.
type IDidDeactivateIterator struct {
	Event *IDidDeactivate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidDeactivateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidDeactivate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidDeactivate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidDeactivateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidDeactivateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidDeactivate represents a Deactivate event raised by the IDid contract.
type IDidDeactivate struct {
	Did string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeactivate is a free log retrieval operation binding the contract event 0xe0d447dee36c113fb0b5edf3262a72a06d967bd571d810cfaa2b7df89749f7ca.
//
// Solidity: event Deactivate(string did)
func (_IDid *IDidFilterer) FilterDeactivate(opts *bind.FilterOpts) (*IDidDeactivateIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "Deactivate")
	if err != nil {
		return nil, err
	}
	return &IDidDeactivateIterator{contract: _IDid.contract, event: "Deactivate", logs: logs, sub: sub}, nil
}

// WatchDeactivate is a free log subscription operation binding the contract event 0xe0d447dee36c113fb0b5edf3262a72a06d967bd571d810cfaa2b7df89749f7ca.
//
// Solidity: event Deactivate(string did)
func (_IDid *IDidFilterer) WatchDeactivate(opts *bind.WatchOpts, sink chan<- *IDidDeactivate) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "Deactivate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidDeactivate)
				if err := _IDid.contract.UnpackLog(event, "Deactivate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivate is a log parse operation binding the contract event 0xe0d447dee36c113fb0b5edf3262a72a06d967bd571d810cfaa2b7df89749f7ca.
//
// Solidity: event Deactivate(string did)
func (_IDid *IDidFilterer) ParseDeactivate(log types.Log) (*IDidDeactivate, error) {
	event := new(IDidDeactivate)
	if err := _IDid.contract.UnpackLog(event, "Deactivate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidDeactivateAddrIterator is returned from FilterDeactivateAddr and is used to iterate over the raw logs and unpacked data for DeactivateAddr events raised by the IDid contract.
type IDidDeactivateAddrIterator struct {
	Event *IDidDeactivateAddr // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidDeactivateAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidDeactivateAddr)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidDeactivateAddr)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidDeactivateAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidDeactivateAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidDeactivateAddr represents a DeactivateAddr event raised by the IDid contract.
type IDidDeactivateAddr struct {
	Did  string
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeactivateAddr is a free log retrieval operation binding the contract event 0x3364de2ecbd53e405faf989044be284732af4a4d2f29f9db04df4bcf1d753b51.
//
// Solidity: event DeactivateAddr(string did, address addr)
func (_IDid *IDidFilterer) FilterDeactivateAddr(opts *bind.FilterOpts) (*IDidDeactivateAddrIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "DeactivateAddr")
	if err != nil {
		return nil, err
	}
	return &IDidDeactivateAddrIterator{contract: _IDid.contract, event: "DeactivateAddr", logs: logs, sub: sub}, nil
}

// WatchDeactivateAddr is a free log subscription operation binding the contract event 0x3364de2ecbd53e405faf989044be284732af4a4d2f29f9db04df4bcf1d753b51.
//
// Solidity: event DeactivateAddr(string did, address addr)
func (_IDid *IDidFilterer) WatchDeactivateAddr(opts *bind.WatchOpts, sink chan<- *IDidDeactivateAddr) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "DeactivateAddr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidDeactivateAddr)
				if err := _IDid.contract.UnpackLog(event, "DeactivateAddr", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivateAddr is a log parse operation binding the contract event 0x3364de2ecbd53e405faf989044be284732af4a4d2f29f9db04df4bcf1d753b51.
//
// Solidity: event DeactivateAddr(string did, address addr)
func (_IDid *IDidFilterer) ParseDeactivateAddr(log types.Log) (*IDidDeactivateAddr, error) {
	event := new(IDidDeactivateAddr)
	if err := _IDid.contract.UnpackLog(event, "DeactivateAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidDeactivateAuthAddrIterator is returned from FilterDeactivateAuthAddr and is used to iterate over the raw logs and unpacked data for DeactivateAuthAddr events raised by the IDid contract.
type IDidDeactivateAuthAddrIterator struct {
	Event *IDidDeactivateAuthAddr // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidDeactivateAuthAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidDeactivateAuthAddr)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidDeactivateAuthAddr)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidDeactivateAuthAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidDeactivateAuthAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidDeactivateAuthAddr represents a DeactivateAuthAddr event raised by the IDid contract.
type IDidDeactivateAuthAddr struct {
	Did  string
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeactivateAuthAddr is a free log retrieval operation binding the contract event 0x5b4ad0c78947966b48caff8e1256800d9ecc297743953b6ba78776d0e81d327f.
//
// Solidity: event DeactivateAuthAddr(string did, address addr)
func (_IDid *IDidFilterer) FilterDeactivateAuthAddr(opts *bind.FilterOpts) (*IDidDeactivateAuthAddrIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "DeactivateAuthAddr")
	if err != nil {
		return nil, err
	}
	return &IDidDeactivateAuthAddrIterator{contract: _IDid.contract, event: "DeactivateAuthAddr", logs: logs, sub: sub}, nil
}

// WatchDeactivateAuthAddr is a free log subscription operation binding the contract event 0x5b4ad0c78947966b48caff8e1256800d9ecc297743953b6ba78776d0e81d327f.
//
// Solidity: event DeactivateAuthAddr(string did, address addr)
func (_IDid *IDidFilterer) WatchDeactivateAuthAddr(opts *bind.WatchOpts, sink chan<- *IDidDeactivateAuthAddr) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "DeactivateAuthAddr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidDeactivateAuthAddr)
				if err := _IDid.contract.UnpackLog(event, "DeactivateAuthAddr", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivateAuthAddr is a log parse operation binding the contract event 0x5b4ad0c78947966b48caff8e1256800d9ecc297743953b6ba78776d0e81d327f.
//
// Solidity: event DeactivateAuthAddr(string did, address addr)
func (_IDid *IDidFilterer) ParseDeactivateAuthAddr(log types.Log) (*IDidDeactivateAuthAddr, error) {
	event := new(IDidDeactivateAuthAddr)
	if err := _IDid.contract.UnpackLog(event, "DeactivateAuthAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidDeactivateAuthKeyIterator is returned from FilterDeactivateAuthKey and is used to iterate over the raw logs and unpacked data for DeactivateAuthKey events raised by the IDid contract.
type IDidDeactivateAuthKeyIterator struct {
	Event *IDidDeactivateAuthKey // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidDeactivateAuthKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidDeactivateAuthKey)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidDeactivateAuthKey)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidDeactivateAuthKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidDeactivateAuthKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidDeactivateAuthKey represents a DeactivateAuthKey event raised by the IDid contract.
type IDidDeactivateAuthKey struct {
	Did    string
	PubKey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeactivateAuthKey is a free log retrieval operation binding the contract event 0x2b5997c0533b1b213e90a8955d60e56efed753c769a096d6e30f68b8fe955d06.
//
// Solidity: event DeactivateAuthKey(string did, bytes pubKey)
func (_IDid *IDidFilterer) FilterDeactivateAuthKey(opts *bind.FilterOpts) (*IDidDeactivateAuthKeyIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "DeactivateAuthKey")
	if err != nil {
		return nil, err
	}
	return &IDidDeactivateAuthKeyIterator{contract: _IDid.contract, event: "DeactivateAuthKey", logs: logs, sub: sub}, nil
}

// WatchDeactivateAuthKey is a free log subscription operation binding the contract event 0x2b5997c0533b1b213e90a8955d60e56efed753c769a096d6e30f68b8fe955d06.
//
// Solidity: event DeactivateAuthKey(string did, bytes pubKey)
func (_IDid *IDidFilterer) WatchDeactivateAuthKey(opts *bind.WatchOpts, sink chan<- *IDidDeactivateAuthKey) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "DeactivateAuthKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidDeactivateAuthKey)
				if err := _IDid.contract.UnpackLog(event, "DeactivateAuthKey", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivateAuthKey is a log parse operation binding the contract event 0x2b5997c0533b1b213e90a8955d60e56efed753c769a096d6e30f68b8fe955d06.
//
// Solidity: event DeactivateAuthKey(string did, bytes pubKey)
func (_IDid *IDidFilterer) ParseDeactivateAuthKey(log types.Log) (*IDidDeactivateAuthKey, error) {
	event := new(IDidDeactivateAuthKey)
	if err := _IDid.contract.UnpackLog(event, "DeactivateAuthKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidDeactivateKeyIterator is returned from FilterDeactivateKey and is used to iterate over the raw logs and unpacked data for DeactivateKey events raised by the IDid contract.
type IDidDeactivateKeyIterator struct {
	Event *IDidDeactivateKey // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidDeactivateKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidDeactivateKey)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidDeactivateKey)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidDeactivateKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidDeactivateKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidDeactivateKey represents a DeactivateKey event raised by the IDid contract.
type IDidDeactivateKey struct {
	Did    string
	PubKey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeactivateKey is a free log retrieval operation binding the contract event 0xc2613a8fbe2f84ee12626ff6451fe839595edd92aa6d8fd95e62e587b42a0e15.
//
// Solidity: event DeactivateKey(string did, bytes pubKey)
func (_IDid *IDidFilterer) FilterDeactivateKey(opts *bind.FilterOpts) (*IDidDeactivateKeyIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "DeactivateKey")
	if err != nil {
		return nil, err
	}
	return &IDidDeactivateKeyIterator{contract: _IDid.contract, event: "DeactivateKey", logs: logs, sub: sub}, nil
}

// WatchDeactivateKey is a free log subscription operation binding the contract event 0xc2613a8fbe2f84ee12626ff6451fe839595edd92aa6d8fd95e62e587b42a0e15.
//
// Solidity: event DeactivateKey(string did, bytes pubKey)
func (_IDid *IDidFilterer) WatchDeactivateKey(opts *bind.WatchOpts, sink chan<- *IDidDeactivateKey) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "DeactivateKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidDeactivateKey)
				if err := _IDid.contract.UnpackLog(event, "DeactivateKey", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeactivateKey is a log parse operation binding the contract event 0xc2613a8fbe2f84ee12626ff6451fe839595edd92aa6d8fd95e62e587b42a0e15.
//
// Solidity: event DeactivateKey(string did, bytes pubKey)
func (_IDid *IDidFilterer) ParseDeactivateKey(log types.Log) (*IDidDeactivateKey, error) {
	event := new(IDidDeactivateKey)
	if err := _IDid.contract.UnpackLog(event, "DeactivateKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidRemoveContextIterator is returned from FilterRemoveContext and is used to iterate over the raw logs and unpacked data for RemoveContext events raised by the IDid contract.
type IDidRemoveContextIterator struct {
	Event *IDidRemoveContext // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidRemoveContextIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidRemoveContext)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidRemoveContext)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidRemoveContextIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidRemoveContextIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidRemoveContext represents a RemoveContext event raised by the IDid contract.
type IDidRemoveContext struct {
	Did     string
	Context string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveContext is a free log retrieval operation binding the contract event 0x58055c159e081c2caccfa4852d9a4285566a460de7ba910919b0ff75271d877f.
//
// Solidity: event RemoveContext(string did, string context)
func (_IDid *IDidFilterer) FilterRemoveContext(opts *bind.FilterOpts) (*IDidRemoveContextIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "RemoveContext")
	if err != nil {
		return nil, err
	}
	return &IDidRemoveContextIterator{contract: _IDid.contract, event: "RemoveContext", logs: logs, sub: sub}, nil
}

// WatchRemoveContext is a free log subscription operation binding the contract event 0x58055c159e081c2caccfa4852d9a4285566a460de7ba910919b0ff75271d877f.
//
// Solidity: event RemoveContext(string did, string context)
func (_IDid *IDidFilterer) WatchRemoveContext(opts *bind.WatchOpts, sink chan<- *IDidRemoveContext) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "RemoveContext")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidRemoveContext)
				if err := _IDid.contract.UnpackLog(event, "RemoveContext", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveContext is a log parse operation binding the contract event 0x58055c159e081c2caccfa4852d9a4285566a460de7ba910919b0ff75271d877f.
//
// Solidity: event RemoveContext(string did, string context)
func (_IDid *IDidFilterer) ParseRemoveContext(log types.Log) (*IDidRemoveContext, error) {
	event := new(IDidRemoveContext)
	if err := _IDid.contract.UnpackLog(event, "RemoveContext", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidRemoveControllerIterator is returned from FilterRemoveController and is used to iterate over the raw logs and unpacked data for RemoveController events raised by the IDid contract.
type IDidRemoveControllerIterator struct {
	Event *IDidRemoveController // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidRemoveControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidRemoveController)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidRemoveController)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidRemoveControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidRemoveControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidRemoveController represents a RemoveController event raised by the IDid contract.
type IDidRemoveController struct {
	Did        string
	Controller string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRemoveController is a free log retrieval operation binding the contract event 0x60b2df265d895cbd29c118ef04311b5ccf506bbc063740f4ed3eb7a60448afe7.
//
// Solidity: event RemoveController(string did, string controller)
func (_IDid *IDidFilterer) FilterRemoveController(opts *bind.FilterOpts) (*IDidRemoveControllerIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "RemoveController")
	if err != nil {
		return nil, err
	}
	return &IDidRemoveControllerIterator{contract: _IDid.contract, event: "RemoveController", logs: logs, sub: sub}, nil
}

// WatchRemoveController is a free log subscription operation binding the contract event 0x60b2df265d895cbd29c118ef04311b5ccf506bbc063740f4ed3eb7a60448afe7.
//
// Solidity: event RemoveController(string did, string controller)
func (_IDid *IDidFilterer) WatchRemoveController(opts *bind.WatchOpts, sink chan<- *IDidRemoveController) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "RemoveController")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidRemoveController)
				if err := _IDid.contract.UnpackLog(event, "RemoveController", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveController is a log parse operation binding the contract event 0x60b2df265d895cbd29c118ef04311b5ccf506bbc063740f4ed3eb7a60448afe7.
//
// Solidity: event RemoveController(string did, string controller)
func (_IDid *IDidFilterer) ParseRemoveController(log types.Log) (*IDidRemoveController, error) {
	event := new(IDidRemoveController)
	if err := _IDid.contract.UnpackLog(event, "RemoveController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidRemoveServiceIterator is returned from FilterRemoveService and is used to iterate over the raw logs and unpacked data for RemoveService events raised by the IDid contract.
type IDidRemoveServiceIterator struct {
	Event *IDidRemoveService // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidRemoveServiceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidRemoveService)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidRemoveService)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidRemoveServiceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidRemoveServiceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidRemoveService represents a RemoveService event raised by the IDid contract.
type IDidRemoveService struct {
	Did       string
	ServiceId string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemoveService is a free log retrieval operation binding the contract event 0xab8833e166d8cbd16f823a1e7c05653d3f78d9c3498ebdf0b2393c16e3fe9487.
//
// Solidity: event RemoveService(string did, string serviceId)
func (_IDid *IDidFilterer) FilterRemoveService(opts *bind.FilterOpts) (*IDidRemoveServiceIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "RemoveService")
	if err != nil {
		return nil, err
	}
	return &IDidRemoveServiceIterator{contract: _IDid.contract, event: "RemoveService", logs: logs, sub: sub}, nil
}

// WatchRemoveService is a free log subscription operation binding the contract event 0xab8833e166d8cbd16f823a1e7c05653d3f78d9c3498ebdf0b2393c16e3fe9487.
//
// Solidity: event RemoveService(string did, string serviceId)
func (_IDid *IDidFilterer) WatchRemoveService(opts *bind.WatchOpts, sink chan<- *IDidRemoveService) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "RemoveService")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidRemoveService)
				if err := _IDid.contract.UnpackLog(event, "RemoveService", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveService is a log parse operation binding the contract event 0xab8833e166d8cbd16f823a1e7c05653d3f78d9c3498ebdf0b2393c16e3fe9487.
//
// Solidity: event RemoveService(string did, string serviceId)
func (_IDid *IDidFilterer) ParseRemoveService(log types.Log) (*IDidRemoveService, error) {
	event := new(IDidRemoveService)
	if err := _IDid.contract.UnpackLog(event, "RemoveService", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidSetAuthAddrIterator is returned from FilterSetAuthAddr and is used to iterate over the raw logs and unpacked data for SetAuthAddr events raised by the IDid contract.
type IDidSetAuthAddrIterator struct {
	Event *IDidSetAuthAddr // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidSetAuthAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidSetAuthAddr)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidSetAuthAddr)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidSetAuthAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidSetAuthAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidSetAuthAddr represents a SetAuthAddr event raised by the IDid contract.
type IDidSetAuthAddr struct {
	Did  string
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterSetAuthAddr is a free log retrieval operation binding the contract event 0x605a8562b3e375915ae30d3f6a56e766cc11be0a6c592b68704f048f5ac9d867.
//
// Solidity: event SetAuthAddr(string did, address addr)
func (_IDid *IDidFilterer) FilterSetAuthAddr(opts *bind.FilterOpts) (*IDidSetAuthAddrIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "SetAuthAddr")
	if err != nil {
		return nil, err
	}
	return &IDidSetAuthAddrIterator{contract: _IDid.contract, event: "SetAuthAddr", logs: logs, sub: sub}, nil
}

// WatchSetAuthAddr is a free log subscription operation binding the contract event 0x605a8562b3e375915ae30d3f6a56e766cc11be0a6c592b68704f048f5ac9d867.
//
// Solidity: event SetAuthAddr(string did, address addr)
func (_IDid *IDidFilterer) WatchSetAuthAddr(opts *bind.WatchOpts, sink chan<- *IDidSetAuthAddr) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "SetAuthAddr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidSetAuthAddr)
				if err := _IDid.contract.UnpackLog(event, "SetAuthAddr", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAuthAddr is a log parse operation binding the contract event 0x605a8562b3e375915ae30d3f6a56e766cc11be0a6c592b68704f048f5ac9d867.
//
// Solidity: event SetAuthAddr(string did, address addr)
func (_IDid *IDidFilterer) ParseSetAuthAddr(log types.Log) (*IDidSetAuthAddr, error) {
	event := new(IDidSetAuthAddr)
	if err := _IDid.contract.UnpackLog(event, "SetAuthAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidSetAuthKeyIterator is returned from FilterSetAuthKey and is used to iterate over the raw logs and unpacked data for SetAuthKey events raised by the IDid contract.
type IDidSetAuthKeyIterator struct {
	Event *IDidSetAuthKey // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidSetAuthKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidSetAuthKey)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidSetAuthKey)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidSetAuthKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidSetAuthKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidSetAuthKey represents a SetAuthKey event raised by the IDid contract.
type IDidSetAuthKey struct {
	Did    string
	PubKey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetAuthKey is a free log retrieval operation binding the contract event 0x40f36fc5010e18c1e871f706915909d1fcf1215d949a2c4d72350380162b1bd5.
//
// Solidity: event SetAuthKey(string did, bytes pubKey)
func (_IDid *IDidFilterer) FilterSetAuthKey(opts *bind.FilterOpts) (*IDidSetAuthKeyIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "SetAuthKey")
	if err != nil {
		return nil, err
	}
	return &IDidSetAuthKeyIterator{contract: _IDid.contract, event: "SetAuthKey", logs: logs, sub: sub}, nil
}

// WatchSetAuthKey is a free log subscription operation binding the contract event 0x40f36fc5010e18c1e871f706915909d1fcf1215d949a2c4d72350380162b1bd5.
//
// Solidity: event SetAuthKey(string did, bytes pubKey)
func (_IDid *IDidFilterer) WatchSetAuthKey(opts *bind.WatchOpts, sink chan<- *IDidSetAuthKey) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "SetAuthKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidSetAuthKey)
				if err := _IDid.contract.UnpackLog(event, "SetAuthKey", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetAuthKey is a log parse operation binding the contract event 0x40f36fc5010e18c1e871f706915909d1fcf1215d949a2c4d72350380162b1bd5.
//
// Solidity: event SetAuthKey(string did, bytes pubKey)
func (_IDid *IDidFilterer) ParseSetAuthKey(log types.Log) (*IDidSetAuthKey, error) {
	event := new(IDidSetAuthKey)
	if err := _IDid.contract.UnpackLog(event, "SetAuthKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IDidUpdateServiceIterator is returned from FilterUpdateService and is used to iterate over the raw logs and unpacked data for UpdateService events raised by the IDid contract.
type IDidUpdateServiceIterator struct {
	Event *IDidUpdateService // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IDidUpdateServiceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IDidUpdateService)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IDidUpdateService)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IDidUpdateServiceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IDidUpdateServiceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IDidUpdateService represents a UpdateService event raised by the IDid contract.
type IDidUpdateService struct {
	Did             string
	ServiceId       string
	ServiceType     string
	ServiceEndpoint string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateService is a free log retrieval operation binding the contract event 0xce02a40ad64d07c09428e43775774b9c09a2fb4050cb2133feffdaf82e7c201f.
//
// Solidity: event UpdateService(string did, string serviceId, string serviceType, string serviceEndpoint)
func (_IDid *IDidFilterer) FilterUpdateService(opts *bind.FilterOpts) (*IDidUpdateServiceIterator, error) {

	logs, sub, err := _IDid.contract.FilterLogs(opts, "UpdateService")
	if err != nil {
		return nil, err
	}
	return &IDidUpdateServiceIterator{contract: _IDid.contract, event: "UpdateService", logs: logs, sub: sub}, nil
}

// WatchUpdateService is a free log subscription operation binding the contract event 0xce02a40ad64d07c09428e43775774b9c09a2fb4050cb2133feffdaf82e7c201f.
//
// Solidity: event UpdateService(string did, string serviceId, string serviceType, string serviceEndpoint)
func (_IDid *IDidFilterer) WatchUpdateService(opts *bind.WatchOpts, sink chan<- *IDidUpdateService) (event.Subscription, error) {

	logs, sub, err := _IDid.contract.WatchLogs(opts, "UpdateService")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IDidUpdateService)
				if err := _IDid.contract.UnpackLog(event, "UpdateService", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateService is a log parse operation binding the contract event 0xce02a40ad64d07c09428e43775774b9c09a2fb4050cb2133feffdaf82e7c201f.
//
// Solidity: event UpdateService(string did, string serviceId, string serviceType, string serviceEndpoint)
func (_IDid *IDidFilterer) ParseUpdateService(log types.Log) (*IDidUpdateService, error) {
	event := new(IDidUpdateService)
	if err := _IDid.contract.UnpackLog(event, "UpdateService", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
