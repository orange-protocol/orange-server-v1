package wasm

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/ont-bizsuite/wagon/exec"
	"github.com/ont-bizsuite/wagon/wasm"

	"github.com/orange-protocol/orange-server-v1/log"
)

const gasSHA256 = 100
const gasPerDataReq = 10000

type Runtime struct {
	vm         *exec.VM
	Input      []byte
	Output     []byte
	CallOutput []byte
	Env        *ExecEnv
}

func (self *Runtime) checkGas(gaslimit uint64) {
	gas := self.vm.ExecMetrics
	if *gas.GasLimit >= gaslimit {
		*gas.GasLimit -= gaslimit
	} else {
		panic(errors.New("[runtime] insufficient gas limit"))
	}
}

func Sha256(proc *exec.Process, src uint32, slen uint32, dst uint32) {
	self := proc.HostData().(*Runtime)
	cost := uint64((slen/1024)+1) * gasSHA256
	self.checkGas(cost)

	bs, err := ReadWasmMemory(proc, src, slen)
	if err != nil {
		panic(err)
	}

	sh := sha256.New()
	sh.Write(bs[:])
	hash := sh.Sum(nil)

	_, err = proc.WriteAt(hash[:], int64(dst))
	if err != nil {
		panic(err)
	}
}

// vmrpc is used to make vm to communicate with host env. the format is :
// req-> {"version": 0, "method": "add", "params":jsonval }
// response <-  jsonresult
func VmRpc(proc *exec.Process, src uint32, slen uint32) uint32 {
	self := proc.HostData().(*Runtime)
	self.checkGas(gasPerDataReq)

	req, err := ReadWasmMemory(proc, src, slen)
	if err != nil {
		panic(err)
	}

	reply, err := self.Env.Service.HandleWasmRequest(req)
	if err != nil {
		panic(err)
	}
	self.CallOutput = reply

	return uint32(len(reply))
}

func Ret(proc *exec.Process, ptr uint32, len uint32) {
	self := proc.HostData().(*Runtime)
	bs, err := ReadWasmMemory(proc, ptr, len)
	if err != nil {
		panic(err)
	}

	self.Output = bs
	proc.Terminate()
}

func Debug(proc *exec.Process, ptr uint32, len uint32) {
	bs, err := ReadWasmMemory(proc, ptr, len)
	if err != nil {
		//do not panic on debug
		return
	}

	log.Debugf("[runtime] debug:%s\n", bs)
}

func InputLength(proc *exec.Process) uint32 {
	self := proc.HostData().(*Runtime)
	return uint32(len(self.Input))
}

func GetInput(proc *exec.Process, dst uint32) {
	self := proc.HostData().(*Runtime)
	_, err := proc.WriteAt(self.Input, int64(dst))
	if err != nil {
		panic(err)
	}
}

func CallOutputLength(proc *exec.Process) uint32 {
	self := proc.HostData().(*Runtime)
	return uint32(len(self.CallOutput))
}

func GetCallOutput(proc *exec.Process, dst uint32) {
	self := proc.HostData().(*Runtime)
	_, err := proc.WriteAt(self.CallOutput, int64(dst))
	if err != nil {
		panic(err)
	}
}

func RaiseException(proc *exec.Process, ptr uint32, len uint32) {
	bs, err := ReadWasmMemory(proc, ptr, len)
	if err != nil {
		//do not panic on debug
		return
	}

	panic(fmt.Errorf("[runtime] raise exception:%s\n", bs))
}

func NewHostModule() *wasm.Module {
	builder := NewHostModuleBuilder()
	ensure(builder.AppendFunc("oscore_get_input", GetInput))
	ensure(builder.AppendFunc("oscore_input_length", InputLength))
	ensure(builder.AppendFunc("oscore_call_output_length", CallOutputLength))
	ensure(builder.AppendFunc("oscore_get_call_output", GetCallOutput))
	ensure(builder.AppendFunc("oscore_return", Ret))
	ensure(builder.AppendFunc("oscore_debug", Debug))
	ensure(builder.AppendFunc("oscore_panic", RaiseException))
	ensure(builder.AppendFunc("oscore_sha256", Sha256))
	ensure(builder.AppendFunc("oscore_vmrpc", VmRpc))

	return builder.Finish()
}

func ensure(err error) {
	if err != nil {
		panic(err)
	}
}
