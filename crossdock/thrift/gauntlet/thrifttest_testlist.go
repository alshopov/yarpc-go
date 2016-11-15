// Code generated by thriftrw v1.0.0
// @generated

// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package gauntlet

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type ThriftTest_TestList_Args struct {
	Thing []int32 `json:"thing"`
}

func (v *ThriftTest_TestList_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Thing != nil {
		w, err = wire.NewValueList(_List_I32_ValueList(v.Thing)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ThriftTest_TestList_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TList {
				v.Thing, err = _List_I32_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *ThriftTest_TestList_Args) String() string {
	var fields [1]string
	i := 0
	if v.Thing != nil {
		fields[i] = fmt.Sprintf("Thing: %v", v.Thing)
		i++
	}
	return fmt.Sprintf("ThriftTest_TestList_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *ThriftTest_TestList_Args) MethodName() string {
	return "testList"
}

func (v *ThriftTest_TestList_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var ThriftTest_TestList_Helper = struct {
	Args           func(thing []int32) *ThriftTest_TestList_Args
	IsException    func(error) bool
	WrapResponse   func([]int32, error) (*ThriftTest_TestList_Result, error)
	UnwrapResponse func(*ThriftTest_TestList_Result) ([]int32, error)
}{}

func init() {
	ThriftTest_TestList_Helper.Args = func(thing []int32) *ThriftTest_TestList_Args {
		return &ThriftTest_TestList_Args{Thing: thing}
	}
	ThriftTest_TestList_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	ThriftTest_TestList_Helper.WrapResponse = func(success []int32, err error) (*ThriftTest_TestList_Result, error) {
		if err == nil {
			return &ThriftTest_TestList_Result{Success: success}, nil
		}
		return nil, err
	}
	ThriftTest_TestList_Helper.UnwrapResponse = func(result *ThriftTest_TestList_Result) (success []int32, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type ThriftTest_TestList_Result struct {
	Success []int32 `json:"success"`
}

func (v *ThriftTest_TestList_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueList(_List_I32_ValueList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("ThriftTest_TestList_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ThriftTest_TestList_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TList {
				v.Success, err = _List_I32_Read(field.Value.GetList())
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("ThriftTest_TestList_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *ThriftTest_TestList_Result) String() string {
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("ThriftTest_TestList_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *ThriftTest_TestList_Result) MethodName() string {
	return "testList"
}

func (v *ThriftTest_TestList_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
