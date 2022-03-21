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
