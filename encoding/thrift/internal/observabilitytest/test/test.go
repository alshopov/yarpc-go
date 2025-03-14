// Code generated by thriftrw v1.32.0. DO NOT EDIT.
// @generated

// Copyright (c) 2025 Uber Technologies, Inc.
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

package test

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	stream "go.uber.org/thriftrw/protocol/stream"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type ExceptionWithCode struct {
	Val string `json:"val,required"`
}

// ToWire translates a ExceptionWithCode struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//		return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//		return err
//	}
func (v *ExceptionWithCode) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Val), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ExceptionWithCode struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ExceptionWithCode struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//		return nil, err
//	}
//
//	var v ExceptionWithCode
//	if err := v.FromWire(x); err != nil {
//		return nil, err
//	}
//	return &v, nil
func (v *ExceptionWithCode) FromWire(w wire.Value) error {
	var err error

	valIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Val, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				valIsSet = true
			}
		}
	}

	if !valIsSet {
		return errors.New("field Val of ExceptionWithCode is required")
	}

	return nil
}

// Encode serializes a ExceptionWithCode struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a ExceptionWithCode struct could not be encoded.
func (v *ExceptionWithCode) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Val); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a ExceptionWithCode struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a ExceptionWithCode struct could not be generated from the wire
// representation.
func (v *ExceptionWithCode) Decode(sr stream.Reader) error {

	valIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.Val, err = sr.ReadString()
			if err != nil {
				return err
			}
			valIsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !valIsSet {
		return errors.New("field Val of ExceptionWithCode is required")
	}

	return nil
}

// String returns a readable string representation of a ExceptionWithCode
// struct.
func (v *ExceptionWithCode) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Val: %v", v.Val)
	i++

	return fmt.Sprintf("ExceptionWithCode{%v}", strings.Join(fields[:i], ", "))
}

// ErrorName is the name of this type as defined in the Thrift
// file.
func (*ExceptionWithCode) ErrorName() string {
	return "ExceptionWithCode"
}

// Equals returns true if all the fields of this ExceptionWithCode match the
// provided ExceptionWithCode.
//
// This function performs a deep comparison.
func (v *ExceptionWithCode) Equals(rhs *ExceptionWithCode) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Val == rhs.Val) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ExceptionWithCode.
func (v *ExceptionWithCode) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("val", v.Val)
	return err
}

// GetVal returns the value of Val if it is set or its
// zero value if it is unset.
func (v *ExceptionWithCode) GetVal() (o string) {
	if v != nil {
		o = v.Val
	}
	return
}

func (v *ExceptionWithCode) Error() string {
	return v.String()
}

type ExceptionWithoutCode struct {
	Val string `json:"val,required"`
}

// ToWire translates a ExceptionWithoutCode struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//		return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//		return err
//	}
func (v *ExceptionWithoutCode) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Val), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a ExceptionWithoutCode struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a ExceptionWithoutCode struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//		return nil, err
//	}
//
//	var v ExceptionWithoutCode
//	if err := v.FromWire(x); err != nil {
//		return nil, err
//	}
//	return &v, nil
func (v *ExceptionWithoutCode) FromWire(w wire.Value) error {
	var err error

	valIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Val, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				valIsSet = true
			}
		}
	}

	if !valIsSet {
		return errors.New("field Val of ExceptionWithoutCode is required")
	}

	return nil
}

// Encode serializes a ExceptionWithoutCode struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a ExceptionWithoutCode struct could not be encoded.
func (v *ExceptionWithoutCode) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Val); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a ExceptionWithoutCode struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a ExceptionWithoutCode struct could not be generated from the wire
// representation.
func (v *ExceptionWithoutCode) Decode(sr stream.Reader) error {

	valIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.Val, err = sr.ReadString()
			if err != nil {
				return err
			}
			valIsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !valIsSet {
		return errors.New("field Val of ExceptionWithoutCode is required")
	}

	return nil
}

// String returns a readable string representation of a ExceptionWithoutCode
// struct.
func (v *ExceptionWithoutCode) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Val: %v", v.Val)
	i++

	return fmt.Sprintf("ExceptionWithoutCode{%v}", strings.Join(fields[:i], ", "))
}

// ErrorName is the name of this type as defined in the Thrift
// file.
func (*ExceptionWithoutCode) ErrorName() string {
	return "ExceptionWithoutCode"
}

// Equals returns true if all the fields of this ExceptionWithoutCode match the
// provided ExceptionWithoutCode.
//
// This function performs a deep comparison.
func (v *ExceptionWithoutCode) Equals(rhs *ExceptionWithoutCode) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Val == rhs.Val) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of ExceptionWithoutCode.
func (v *ExceptionWithoutCode) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("val", v.Val)
	return err
}

// GetVal returns the value of Val if it is set or its
// zero value if it is unset.
func (v *ExceptionWithoutCode) GetVal() (o string) {
	if v != nil {
		o = v.Val
	}
	return
}

func (v *ExceptionWithoutCode) Error() string {
	return v.String()
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "test",
	Package:  "go.uber.org/yarpc/encoding/thrift/internal/observabilitytest/test",
	FilePath: "test.thrift",
	SHA1:     "3c501036fe37f678648dd479c821bc57aa17b7d1",
	Raw:      rawIDL,
}

const rawIDL = "exception ExceptionWithCode {\n    1: required string val\n} (\n    rpc.code = \"DATA_LOSS\" // server error\n)\n\nexception ExceptionWithoutCode {\n    1: required string val\n}\n\nservice TestService  {\n    string Call(1: required string key) throws (\n      1: ExceptionWithCode exCode,\n      2: ExceptionWithoutCode exNoCode,\n    )\n}\n"

// TestService_Call_Args represents the arguments for the TestService.Call function.
//
// The arguments for Call are sent and received over the wire as this struct.
type TestService_Call_Args struct {
	Key string `json:"key,required"`
}

// ToWire translates a TestService_Call_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//		return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//		return err
//	}
func (v *TestService_Call_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Key), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a TestService_Call_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a TestService_Call_Args struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//		return nil, err
//	}
//
//	var v TestService_Call_Args
//	if err := v.FromWire(x); err != nil {
//		return nil, err
//	}
//	return &v, nil
func (v *TestService_Call_Args) FromWire(w wire.Value) error {
	var err error

	keyIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Key, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				keyIsSet = true
			}
		}
	}

	if !keyIsSet {
		return errors.New("field Key of TestService_Call_Args is required")
	}

	return nil
}

// Encode serializes a TestService_Call_Args struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a TestService_Call_Args struct could not be encoded.
func (v *TestService_Call_Args) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Key); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a TestService_Call_Args struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a TestService_Call_Args struct could not be generated from the wire
// representation.
func (v *TestService_Call_Args) Decode(sr stream.Reader) error {

	keyIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.Key, err = sr.ReadString()
			if err != nil {
				return err
			}
			keyIsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !keyIsSet {
		return errors.New("field Key of TestService_Call_Args is required")
	}

	return nil
}

// String returns a readable string representation of a TestService_Call_Args
// struct.
func (v *TestService_Call_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Key: %v", v.Key)
	i++

	return fmt.Sprintf("TestService_Call_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this TestService_Call_Args match the
// provided TestService_Call_Args.
//
// This function performs a deep comparison.
func (v *TestService_Call_Args) Equals(rhs *TestService_Call_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Key == rhs.Key) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of TestService_Call_Args.
func (v *TestService_Call_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("key", v.Key)
	return err
}

// GetKey returns the value of Key if it is set or its
// zero value if it is unset.
func (v *TestService_Call_Args) GetKey() (o string) {
	if v != nil {
		o = v.Key
	}
	return
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "Call" for this struct.
func (v *TestService_Call_Args) MethodName() string {
	return "Call"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *TestService_Call_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// TestService_Call_Helper provides functions that aid in handling the
// parameters and return values of the TestService.Call
// function.
var TestService_Call_Helper = struct {
	// Args accepts the parameters of Call in-order and returns
	// the arguments struct for the function.
	Args func(
		key string,
	) *TestService_Call_Args

	// IsException returns true if the given error can be thrown
	// by Call.
	//
	// An error can be thrown by Call only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for Call
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// Call into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by Call
	//
	//   value, err := Call(args)
	//   result, err := TestService_Call_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from Call: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(string, error) (*TestService_Call_Result, error)

	// UnwrapResponse takes the result struct for Call
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if Call threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := TestService_Call_Helper.UnwrapResponse(result)
	UnwrapResponse func(*TestService_Call_Result) (string, error)
}{}

func init() {
	TestService_Call_Helper.Args = func(
		key string,
	) *TestService_Call_Args {
		return &TestService_Call_Args{
			Key: key,
		}
	}

	TestService_Call_Helper.IsException = func(err error) bool {
		switch err.(type) {
		case *ExceptionWithCode:
			return true
		case *ExceptionWithoutCode:
			return true
		default:
			return false
		}
	}

	TestService_Call_Helper.WrapResponse = func(success string, err error) (*TestService_Call_Result, error) {
		if err == nil {
			return &TestService_Call_Result{Success: &success}, nil
		}

		switch e := err.(type) {
		case *ExceptionWithCode:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for TestService_Call_Result.ExCode")
			}
			return &TestService_Call_Result{ExCode: e}, nil
		case *ExceptionWithoutCode:
			if e == nil {
				return nil, errors.New("WrapResponse received non-nil error type with nil value for TestService_Call_Result.ExNoCode")
			}
			return &TestService_Call_Result{ExNoCode: e}, nil
		}

		return nil, err
	}
	TestService_Call_Helper.UnwrapResponse = func(result *TestService_Call_Result) (success string, err error) {
		if result.ExCode != nil {
			err = result.ExCode
			return
		}
		if result.ExNoCode != nil {
			err = result.ExNoCode
			return
		}

		if result.Success != nil {
			success = *result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// TestService_Call_Result represents the result of a TestService.Call function call.
//
// The result of a Call execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type TestService_Call_Result struct {
	// Value returned by Call after a successful execution.
	Success  *string               `json:"success,omitempty"`
	ExCode   *ExceptionWithCode    `json:"exCode,omitempty"`
	ExNoCode *ExceptionWithoutCode `json:"exNoCode,omitempty"`
}

// ToWire translates a TestService_Call_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//	x, err := v.ToWire()
//	if err != nil {
//		return err
//	}
//
//	if err := binaryProtocol.Encode(x, writer); err != nil {
//		return err
//	}
func (v *TestService_Call_Result) ToWire() (wire.Value, error) {
	var (
		fields [3]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = wire.NewValueString(*(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if v.ExCode != nil {
		w, err = v.ExCode.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	if v.ExNoCode != nil {
		w, err = v.ExNoCode.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 2, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("TestService_Call_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _ExceptionWithCode_Read(w wire.Value) (*ExceptionWithCode, error) {
	var v ExceptionWithCode
	err := v.FromWire(w)
	return &v, err
}

func _ExceptionWithoutCode_Read(w wire.Value) (*ExceptionWithoutCode, error) {
	var v ExceptionWithoutCode
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a TestService_Call_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a TestService_Call_Result struct
// from the provided intermediate representation.
//
//	x, err := binaryProtocol.Decode(reader, wire.TStruct)
//	if err != nil {
//		return nil, err
//	}
//
//	var v TestService_Call_Result
//	if err := v.FromWire(x); err != nil {
//		return nil, err
//	}
//	return &v, nil
func (v *TestService_Call_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TBinary {
				var x string
				x, err = field.Value.GetString(), error(nil)
				v.Success = &x
				if err != nil {
					return err
				}

			}
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.ExCode, err = _ExceptionWithCode_Read(field.Value)
				if err != nil {
					return err
				}

			}
		case 2:
			if field.Value.Type() == wire.TStruct {
				v.ExNoCode, err = _ExceptionWithoutCode_Read(field.Value)
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
	if v.ExCode != nil {
		count++
	}
	if v.ExNoCode != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("TestService_Call_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// Encode serializes a TestService_Call_Result struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a TestService_Call_Result struct could not be encoded.
func (v *TestService_Call_Result) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Success != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 0, Type: wire.TBinary}); err != nil {
			return err
		}
		if err := sw.WriteString(*(v.Success)); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	if v.ExCode != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TStruct}); err != nil {
			return err
		}
		if err := v.ExCode.Encode(sw); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	if v.ExNoCode != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TStruct}); err != nil {
			return err
		}
		if err := v.ExNoCode.Encode(sw); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.ExCode != nil {
		count++
	}
	if v.ExNoCode != nil {
		count++
	}

	if count != 1 {
		return fmt.Errorf("TestService_Call_Result should have exactly one field: got %v fields", count)
	}

	return sw.WriteStructEnd()
}

func _ExceptionWithCode_Decode(sr stream.Reader) (*ExceptionWithCode, error) {
	var v ExceptionWithCode
	err := v.Decode(sr)
	return &v, err
}

func _ExceptionWithoutCode_Decode(sr stream.Reader) (*ExceptionWithoutCode, error) {
	var v ExceptionWithoutCode
	err := v.Decode(sr)
	return &v, err
}

// Decode deserializes a TestService_Call_Result struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a TestService_Call_Result struct could not be generated from the wire
// representation.
func (v *TestService_Call_Result) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 0 && fh.Type == wire.TBinary:
			var x string
			x, err = sr.ReadString()
			v.Success = &x
			if err != nil {
				return err
			}

		case fh.ID == 1 && fh.Type == wire.TStruct:
			v.ExCode, err = _ExceptionWithCode_Decode(sr)
			if err != nil {
				return err
			}

		case fh.ID == 2 && fh.Type == wire.TStruct:
			v.ExNoCode, err = _ExceptionWithoutCode_Decode(sr)
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if v.ExCode != nil {
		count++
	}
	if v.ExNoCode != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("TestService_Call_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a TestService_Call_Result
// struct.
func (v *TestService_Call_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [3]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	if v.ExCode != nil {
		fields[i] = fmt.Sprintf("ExCode: %v", v.ExCode)
		i++
	}
	if v.ExNoCode != nil {
		fields[i] = fmt.Sprintf("ExNoCode: %v", v.ExNoCode)
		i++
	}

	return fmt.Sprintf("TestService_Call_Result{%v}", strings.Join(fields[:i], ", "))
}

func _String_EqualsPtr(lhs, rhs *string) bool {
	if lhs != nil && rhs != nil {

		x := *lhs
		y := *rhs
		return (x == y)
	}
	return lhs == nil && rhs == nil
}

// Equals returns true if all the fields of this TestService_Call_Result match the
// provided TestService_Call_Result.
//
// This function performs a deep comparison.
func (v *TestService_Call_Result) Equals(rhs *TestService_Call_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !_String_EqualsPtr(v.Success, rhs.Success) {
		return false
	}
	if !((v.ExCode == nil && rhs.ExCode == nil) || (v.ExCode != nil && rhs.ExCode != nil && v.ExCode.Equals(rhs.ExCode))) {
		return false
	}
	if !((v.ExNoCode == nil && rhs.ExNoCode == nil) || (v.ExNoCode != nil && rhs.ExNoCode != nil && v.ExNoCode.Equals(rhs.ExNoCode))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of TestService_Call_Result.
func (v *TestService_Call_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		enc.AddString("success", *v.Success)
	}
	if v.ExCode != nil {
		err = multierr.Append(err, enc.AddObject("exCode", v.ExCode))
	}
	if v.ExNoCode != nil {
		err = multierr.Append(err, enc.AddObject("exNoCode", v.ExNoCode))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *TestService_Call_Result) GetSuccess() (o string) {
	if v != nil && v.Success != nil {
		return *v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *TestService_Call_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// GetExCode returns the value of ExCode if it is set or its
// zero value if it is unset.
func (v *TestService_Call_Result) GetExCode() (o *ExceptionWithCode) {
	if v != nil && v.ExCode != nil {
		return v.ExCode
	}

	return
}

// IsSetExCode returns true if ExCode is not nil.
func (v *TestService_Call_Result) IsSetExCode() bool {
	return v != nil && v.ExCode != nil
}

// GetExNoCode returns the value of ExNoCode if it is set or its
// zero value if it is unset.
func (v *TestService_Call_Result) GetExNoCode() (o *ExceptionWithoutCode) {
	if v != nil && v.ExNoCode != nil {
		return v.ExNoCode
	}

	return
}

// IsSetExNoCode returns true if ExNoCode is not nil.
func (v *TestService_Call_Result) IsSetExNoCode() bool {
	return v != nil && v.ExNoCode != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "Call" for this struct.
func (v *TestService_Call_Result) MethodName() string {
	return "Call"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *TestService_Call_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
