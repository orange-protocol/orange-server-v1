package wasm

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/ont-bizsuite/wagon/exec"
	"github.com/ont-bizsuite/wagon/validate"
	"github.com/ont-bizsuite/wagon/wasm"
)

func ReadWasmMemory(proc *exec.Process, ptr uint32, len uint32) ([]byte, error) {
	if uint64(proc.MemSize()) < uint64(ptr)+uint64(len) {
		return nil, errors.New("contract create len is greater than memory size")
	}
	keybytes := make([]byte, len)
	_, err := proc.ReadAt(keybytes, int64(ptr))
	if err != nil {
		return nil, err
	}

	return keybytes, nil
}

func checkOntoWasm(m *wasm.Module) error {
	if m.Start != nil {
		return errors.New("[validate] start section is not allowed")
	}

	if m.Export == nil {
		return errors.New("[validate] No export in wasm")
	}

	entry, ok := m.Export.Entries["invoke"]
	if ok == false {
		return errors.New("[validate] invoke entry function does not export")
	}

	if entry.Kind != wasm.ExternalFunction {
		return errors.New("[validate] can only export invoke function entry")
	}

	//get entry index
	index := int64(entry.Index)
	//get function index
	fidx := m.Function.Types[int(index)]
	//get  function type
	ftype := m.Types.Entries[int(fidx)]

	if len(ftype.ReturnTypes) > 0 {
		return errors.New("[validate] ExecCode error! Invoke function return sig error")
	}
	if len(ftype.ParamTypes) > 0 {
		return errors.New("[validate] ExecCode error! Invoke function param sig error")
	}

	return nil
}

func ReadWasmModule(Code []byte, verify bool) (*exec.CompiledModule, error) {
	m, err := wasm.ReadModule(bytes.NewReader(Code), func(name string) (*wasm.Module, error) {
		switch name {
		case "env":
			return NewHostModule(), nil
		}
		return nil, fmt.Errorf("module %q unknown", name)
	})
	if err != nil {
		return nil, err
	}

	if verify {
		err = checkOntoWasm(m)
		if err != nil {
			return nil, err
		}

		err = validate.VerifyModule(m)
		if err != nil {
			return nil, err
		}

		//err = validate.VerifyWasmCodeFromRust(Code)
		//if err != nil {
		//	return nil, err
		//}
	}

	compiled, err := exec.CompileModule(m)
	if err != nil {
		return nil, err
	}

	return compiled, nil
}
