package wasm

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/ont-bizsuite/wagon/exec"
	"github.com/ont-bizsuite/wagon/wasm"
)

type HostModuleBuilder struct {
	sigMap map[string]*wasm.FunctionSig
	module *wasm.Module
}

func NewHostModuleBuilder() *HostModuleBuilder {
	m := wasm.NewModule()
	m.Types = &wasm.SectionTypes{}
	m.Export = &wasm.SectionExports{Entries: make(map[string]wasm.ExportEntry)}

	return &HostModuleBuilder{
		module: m,
		sigMap: make(map[string]*wasm.FunctionSig),
	}
}

func generateSignature(hostType reflect.Type) (*wasm.FunctionSig, error) {
	numIn := hostType.NumIn()
	if numIn == 0 {
		return nil, errors.New("[runtime] host function must has at least one *Process param")
	}

	proc := exec.NewProcess(nil)
	// Check that the function indeed expects a *Process argument.
	if reflect.TypeOf(proc) != hostType.In(0) {
		return nil, fmt.Errorf("[runtime]: the first argument of a host function was %s, expected *Process",
			hostType.In(0).String())
	}

	args := make([]wasm.ValueType, 0, numIn)
	for i := 1; i < numIn; i++ {
		kind := hostType.In(i).Kind()
		switch kind {
		case reflect.Uint32, reflect.Int32:
			args = append(args, wasm.ValueTypeI32)
		case reflect.Uint64, reflect.Int64:
			args = append(args, wasm.ValueTypeI64)
		default:
			return nil, fmt.Errorf("[runtime]: args %d invalid kind=%v, expect uint32/64", i, kind)
		}
	}

	var output []wasm.ValueType
	for i := 0; i < hostType.NumOut(); i++ {
		kind := hostType.Out(i).Kind()
		switch kind {
		case reflect.Uint32, reflect.Int32:
			output = append(output, wasm.ValueTypeI32)
		case reflect.Uint64, reflect.Int64:
			output = append(output, wasm.ValueTypeI64)
		default:
			return nil, fmt.Errorf("[runtime]: output %d invalid kind=%v", i, kind)
		}
	}

	return &wasm.FunctionSig{ParamTypes: args, ReturnTypes: output}, nil
}

func (self *HostModuleBuilder) Finish() *wasm.Module {
	return self.module
}

func (self *HostModuleBuilder) AppendFunc(name string, fn interface{}) error {
	hostFunc := reflect.ValueOf(fn)
	hostType := hostFunc.Type()
	if hostType.Kind() != reflect.Func {
		return errors.New("host module only allow func type")
	}

	sig, err := generateSignature(hostType)
	if err != nil {
		return err
	}

	fnSig, ok := self.sigMap[sig.String()]
	if !ok {
		fnSig = sig
		self.sigMap[sig.String()] = sig
		self.module.Types.Entries = append(self.module.Types.Entries, *sig)
	}

	self.module.FunctionIndexSpace = append(self.module.FunctionIndexSpace,
		wasm.Function{
			Sig:  fnSig,
			Host: hostFunc,
			Body: &wasm.FunctionBody{}, // create a dummy wasm body (the actual value will be taken from Host.)
		})

	self.module.Export.Entries[name] = wasm.ExportEntry{
		FieldStr: name,
		Kind:     wasm.ExternalFunction,
		Index:    uint32(len(self.module.FunctionIndexSpace)) - 1,
	}

	return nil
}
