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

type ThriftTest_TestEnum_Args struct {
	Thing *Numberz `json:"thing,omitempty"`
}

func (v *ThriftTest_TestEnum_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Thing != nil {
		w, err = v.Thing.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ThriftTest_TestEnum_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x Numberz
				x, err = _Numberz_Read(field.Value)
				v.Thing = &x
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (v *ThriftTest_TestEnum_Args) String() string {
	var fields [1]string
	i := 0
	if v.Thing != nil {
		fields[i] = fmt.Sprintf("Thing: %v", *(v.Thing))
		i++
	}
	return fmt.Sprintf("ThriftTest_TestEnum_Args{%v}", strings.Join(fields[:i], ", "))
}

func (v *ThriftTest_TestEnum_Args) MethodName() string {
	return "testEnum"
}

func (v *ThriftTest_TestEnum_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var ThriftTest_TestEnum_Helper = struct {
	Args           func(thing *Numberz) *ThriftTest_TestEnum_Args
	IsException    func(error) bool
	WrapResponse   func(Numberz, error) (*ThriftTest_TestEnum_Result, error)
	UnwrapResponse func(*ThriftTest_TestEnum_Result) (Numberz, error)
}{}

func init() {
	ThriftTest_TestEnum_Helper.Args = func(thing *Numberz) *ThriftTest_TestEnum_Args {
		return &ThriftTest_TestEnum_Args{Thing: thing}
	}
	ThriftTest_TestEnum_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	ThriftTest_TestEnum_Helper.WrapResponse = func(success Numberz, err error) (*ThriftTest_TestEnum_Result, error) {
		if err == nil {
			return &ThriftTest_TestEnum_Result{Success: &success}, nil
		}
		return nil, err
	}
	ThriftTest_TestEnum_Helper.UnwrapResponse = func(result *ThriftTest_TestEnum_Result) (success Numberz, err error) {
		if result.Success != nil {
			success = *result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type ThriftTest_TestEnum_Result struct {
	Success *Numberz `json:"success,omitempty"`
}

func (v *ThriftTest_TestEnum_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("ThriftTest_TestEnum_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *ThriftTest_TestEnum_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI32 {
				var x Numberz
				x, err = _Numberz_Read(field.Value)
				v.Success = &x
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
		return fmt.Errorf("ThriftTest_TestEnum_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *ThriftTest_TestEnum_Result) String() string {
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	return fmt.Sprintf("ThriftTest_TestEnum_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *ThriftTest_TestEnum_Result) MethodName() string {
	return "testEnum"
}

func (v *ThriftTest_TestEnum_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
