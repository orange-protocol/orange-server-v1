package wasm

import (
	"testing"

	"github.com/ont-bizsuite/wagon/exec"
	"github.com/stretchr/testify/assert"
)

func TestHostModuleBuilder_AppendFunc(t *testing.T) {
	check := assert.New(t)

	builder := NewHostModuleBuilder()
	check.Nil(builder.AppendFunc("", func(a *exec.Process) {}))
	check.Nil(builder.AppendFunc("", func(a *exec.Process, b int32, c int64, d uint32, e uint64) {}))
	check.Nil(builder.AppendFunc("", func(a *exec.Process, b int32, c int64) uint32 { return uint32(0) }))

	check.NotNil(builder.AppendFunc("", 10))
	check.NotNil(builder.AppendFunc("", func(a *int) {}))
	check.NotNil(builder.AppendFunc("", func(a *exec.Process, b *int) {}))
}
