package wasm

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/orange-protocol/orange-server-v1/wasm/types"

	lru "github.com/hashicorp/golang-lru"
	"github.com/ont-bizsuite/wagon/exec"
)

const codeCacheSize = 1024
const wasmMemLimit = 10 * 1024 * 1024
const wasmCallStackLimit = 2 * 1024

var CodeCache, _ = lru.NewARC(codeCacheSize)

type Address [32]byte

func AddressFromCode(code []byte) Address {
	return Address(sha256.Sum256(code))
}

func (self Address) ToHex() string {
	return hex.EncodeToString(self[:])
}

type ExecEnv struct {
	Service *WasmService
}

type Executor struct {
	//Input    []byte
	GasLimit uint64
	ExecStep uint64

	Env       *ExecEnv
	entryName string
}

func NewExecutor(gasLimit, stepLimit uint64, env *ExecEnv) *Executor {
	//input, _ := json.Marshal(info)
	return &Executor{
		//Input:     input,
		GasLimit:  gasLimit,
		ExecStep:  stepLimit,
		Env:       env,
		entryName: "invoke",
	}
}

func VerifyWasmModule(wasmCode []byte) error {
	_, err := ReadWasmModule(wasmCode, true)
	return err
}

func (self *Executor) Invoke(input, wasmCode []byte) (*types.ScoreResult, error) {
	addr := AddressFromCode(wasmCode)

	var compiled *exec.CompiledModule
	if CodeCache != nil {
		cached, ok := CodeCache.Get(addr)
		if ok {
			compiled = cached.(*exec.CompiledModule)
		}
	}

	var err error
	if compiled == nil {
		compiled, err = ReadWasmModule(wasmCode, true)
		if err != nil {
			return nil, err
		}
		CodeCache.Add(addr, compiled)
	}

	vm, err := exec.NewVMWithCompiled(compiled, wasmMemLimit)
	if err != nil {
		return nil, err
	}

	host := &Runtime{Input: input, Env: self.Env}
	vm.HostData = host
	vm.ExecMetrics = &exec.Gas{GasLimit: &self.GasLimit, GasPrice: 1, GasFactor: 1, ExecStep: &self.ExecStep}
	vm.CallStackDepth = uint32(wasmCallStackLimit)
	vm.RecoverPanic = true

	entry, ok := compiled.RawModule.Export.Entries[self.entryName]
	if ok == false {
		return nil, errors.New("wasm function " + self.entryName + " does not exist")
	}

	//get entry index
	index := int64(entry.Index)
	fidx := compiled.RawModule.Function.Types[int(index)]
	ftype := compiled.RawModule.Types.Entries[int(fidx)]

	//no returns of the entry function
	if len(ftype.ReturnTypes) > 0 {
		return nil, errors.New("invoke function sig error")
	}

	_, err = vm.ExecCode(index)
	if err != nil {
		return nil, errors.New("exec wasm code error: " + err.Error())
	}

	var score types.ScoreResult
	err = json.Unmarshal(host.Output, &score)
	if err != nil {
		return nil, fmt.Errorf("wasm output is invalid %s", string(host.Output))
	}

	return &score, nil
}
