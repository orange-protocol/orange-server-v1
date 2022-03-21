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
	"encoding/json"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Args struct {
	A, B int
}

type Reply struct {
	C int
}

func Add(args *Args, reply *Reply) error {
	reply.C = args.A + args.B
	return nil
}

func Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func Map(i int, reply *map[int]int) error {
	(*reply)[i] = i
	return nil
}

func Slice(i int, reply *[]int) error {
	*reply = append(*reply, i)
	return nil
}

func Array(i int, reply *[1]int) error {
	(*reply)[0] = i
	return nil
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}

func initService() *WasmService {
	service := NewWasmService()
	checkerr(service.Register("add", Add))
	checkerr(service.Register("mul", Mul))
	checkerr(service.Register("map", Map))
	checkerr(service.Register("slice", Slice))
	checkerr(service.Register("array", Array))

	return service
}

func TestServerNoParams(t *testing.T) {
	ast := assert.New(t)
	service := initService()
	_, err := service.HandleWasmRequest([]byte(`{"method": "add"}`))
	ast.NotNil(err)
	_, err = service.HandleWasmRequest([]byte(`{}`))
	ast.NotNil(err)

	rep, err := service.HandleWasmRequest([]byte(`{"method": "add", "params":{"A":1, "B":2}}`))
	ast.Nil(err)
	var resp Reply
	ast.Nil(json.Unmarshal(rep, &resp))
	ast.Equal(resp.C, 3)
	rep, err = service.HandleWasmRequest([]byte(`{"method": "mul", "params":{"A":3, "B":2}}`))
	ast.Nil(err)
	ast.Nil(json.Unmarshal(rep, &resp))
	ast.Equal(resp.C, 6)

	// test builtin types
	arg := 7
	replyMap := map[int]int{}
	ast.Nil(service.Call("map", arg, &replyMap))
	ast.Equal(replyMap[arg], arg)

	var replySlice []int
	ast.Nil(service.Call("slice", arg, &replySlice))
	ast.True(reflect.DeepEqual(replySlice, []int{arg}))

	replyArray := [1]int{}
	ast.Nil(service.Call("array", arg, &replyArray))
	ast.Equal(replyArray, [1]int{arg})
}
