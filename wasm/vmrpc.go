package wasm

import (
	"encoding/json"
	"errors"
	"fmt"
	"go/token"
	"reflect"
)

// Precompute the reflect type for error. Can't use error directly
// because Typeof takes an empty interface value. This is annoying.
var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

type WasmService struct {
	methods map[string]*methodType // registered methods
}

func NewWasmService() *WasmService {
	return &WasmService{
		methods: make(map[string]*methodType),
	}
}

type serverRequest struct {
	Version uint32           `json:"version"`
	Method  string           `json:"method"`
	Params  *json.RawMessage `json:"params"`
}

type methodType struct {
	method    reflect.Type
	Func      reflect.Value
	ArgType   reflect.Type
	ReplyType reflect.Type
}

func (self *WasmService) Call(method string, argv interface{}, reply interface{}) error {
	req := struct {
		Method string      `json:"method"`
		Params interface{} `json:"params"`
	}{method, argv}

	buf, err := json.Marshal(req)
	if err != nil {
		return err
	}
	rep, err := self.HandleWasmRequest(buf)
	if err != nil {
		return err
	}

	return json.Unmarshal(rep, reply)
}

func (self *WasmService) HandleWasmRequest(req []byte) (rep []byte, err error) {
	method, argv, replyv, err := self.readRequest(req)
	if err != nil {
		return
	}
	function := method.Func
	// Invoke the method, providing a new value for the reply.
	returnValues := function.Call([]reflect.Value{argv, replyv})
	// The return value for the method is an error.
	e := returnValues[0].Interface()
	if e != nil {
		err = e.(error)
		return
	}

	return json.Marshal(replyv.Interface())
}

func (self *WasmService) readRequest(req []byte) (method *methodType, argv, replyv reflect.Value, err error) {
	var request serverRequest
	err = json.Unmarshal(req, &request)
	if err != nil {
		return
	}

	method = self.methods[request.Method]
	if method == nil {
		err = fmt.Errorf("wasm request func %s no found", request.Method)
		return
	}

	// Decode the argument value.
	argIsValue := false // if true, need to indirect before calling.
	if method.ArgType.Kind() == reflect.Ptr {
		argv = reflect.New(method.ArgType.Elem())
	} else {
		argv = reflect.New(method.ArgType)
		argIsValue = true
	}
	// argv guaranteed to be a pointer now.
	if request.Params == nil {
		err = errors.New("missing param")
		return
	}
	err = json.Unmarshal(*request.Params, argv.Interface())
	if err != nil {
		return
	}

	if argIsValue {
		argv = argv.Elem()
	}

	replyv = reflect.New(method.ReplyType.Elem())

	switch method.ReplyType.Elem().Kind() {
	case reflect.Map:
		replyv.Elem().Set(reflect.MakeMap(method.ReplyType.Elem()))
	case reflect.Slice:
		replyv.Elem().Set(reflect.MakeSlice(method.ReplyType.Elem(), 0, 0))
	}
	return
}

// fnval must be a func with signature: func (req T, reply *O) error
func (self *WasmService) Register(name string, fnval interface{}) error {
	if self.methods[name] != nil {
		return fmt.Errorf("can not register duplicated func %s", name)
	}

	fn := reflect.TypeOf(fnval)
	if fn.Kind() != reflect.Func {
		return fmt.Errorf("the type of fn %s is %q, expect func", name, fn)
	}
	// Method needs three ins: receiver, *args, *reply.
	if fn.NumIn() != 2 {
		return fmt.Errorf("func %s has %d input parameters; needs exactly 2", name, fn.NumIn())
	}
	// First arg need not be a pointer.
	argType := fn.In(0)
	if !isExportedOrBuiltinType(argType) {
		return fmt.Errorf("argument type of method %q is not exported: %q", name, argType)
	}
	// Second arg must be a pointer.
	replyType := fn.In(1)
	if replyType.Kind() != reflect.Ptr {
		return fmt.Errorf("reply type of method %q is not a pointer: %q", name, replyType)
	}
	// Reply type must be exported.
	if !isExportedOrBuiltinType(replyType) {
		return fmt.Errorf("reply type of method %q is not exported: %q", name, replyType)
	}
	// Method needs one out.
	if fn.NumOut() != 1 {
		return fmt.Errorf("method %q has %d output parameters; needs exactly one", name, fn.NumOut())
	}
	// The return type of the method must be error.
	if returnType := fn.Out(0); returnType != typeOfError {
		return fmt.Errorf("return type of method %q is %q, must be error", name, returnType)
	}

	self.methods[name] = &methodType{method: fn, Func: reflect.ValueOf(fnval), ArgType: argType, ReplyType: replyType}

	return nil
}

// Is this type exported or a builtin?
func isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return token.IsExported(t.Name()) || t.PkgPath() == ""
}
