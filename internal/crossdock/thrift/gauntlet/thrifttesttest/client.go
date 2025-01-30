// Code generated by thriftrw-plugin-yarpc
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

package thrifttesttest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	yarpc "go.uber.org/yarpc"
	gauntlet "go.uber.org/yarpc/internal/crossdock/thrift/gauntlet"
	thrifttestclient "go.uber.org/yarpc/internal/crossdock/thrift/gauntlet/thrifttestclient"
)

// MockClient implements a gomock-compatible mock client for service
// ThriftTest.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *_MockClientRecorder
}

var _ thrifttestclient.Interface = (*MockClient)(nil)

type _MockClientRecorder struct {
	mock *MockClient
}

// Build a new mock client for service ThriftTest.
//
//	mockCtrl := gomock.NewController(t)
//	client := thrifttesttest.NewMockClient(mockCtrl)
//
// Use EXPECT() to set expectations on the mock.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &_MockClientRecorder{mock}
	return mock
}

// EXPECT returns an object that allows you to define an expectation on the
// ThriftTest mock client.
func (m *MockClient) EXPECT() *_MockClientRecorder {
	return m.recorder
}

// TestBinary responds to a TestBinary call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestBinary(gomock.Any(), ...).Return(...)
//	... := client.TestBinary(...)
func (m *MockClient) TestBinary(
	ctx context.Context,
	_Thing []byte,
	opts ...yarpc.CallOption,
) (success []byte, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestBinary", args...)
	success, _ = ret[i].([]byte)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestBinary(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestBinary", args...)
}

// TestByte responds to a TestByte call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestByte(gomock.Any(), ...).Return(...)
//	... := client.TestByte(...)
func (m *MockClient) TestByte(
	ctx context.Context,
	_Thing *int8,
	opts ...yarpc.CallOption,
) (success int8, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestByte", args...)
	success, _ = ret[i].(int8)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestByte(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestByte", args...)
}

// TestDouble responds to a TestDouble call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestDouble(gomock.Any(), ...).Return(...)
//	... := client.TestDouble(...)
func (m *MockClient) TestDouble(
	ctx context.Context,
	_Thing *float64,
	opts ...yarpc.CallOption,
) (success float64, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestDouble", args...)
	success, _ = ret[i].(float64)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestDouble(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestDouble", args...)
}

// TestEnum responds to a TestEnum call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestEnum(gomock.Any(), ...).Return(...)
//	... := client.TestEnum(...)
func (m *MockClient) TestEnum(
	ctx context.Context,
	_Thing *gauntlet.Numberz,
	opts ...yarpc.CallOption,
) (success gauntlet.Numberz, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestEnum", args...)
	success, _ = ret[i].(gauntlet.Numberz)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestEnum(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestEnum", args...)
}

// TestException responds to a TestException call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestException(gomock.Any(), ...).Return(...)
//	... := client.TestException(...)
func (m *MockClient) TestException(
	ctx context.Context,
	_Arg *string,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx, _Arg}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestException", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestException(
	ctx interface{},
	_Arg interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Arg}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestException", args...)
}

// TestI32 responds to a TestI32 call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestI32(gomock.Any(), ...).Return(...)
//	... := client.TestI32(...)
func (m *MockClient) TestI32(
	ctx context.Context,
	_Thing *int32,
	opts ...yarpc.CallOption,
) (success int32, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestI32", args...)
	success, _ = ret[i].(int32)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestI32(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestI32", args...)
}

// TestI64 responds to a TestI64 call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestI64(gomock.Any(), ...).Return(...)
//	... := client.TestI64(...)
func (m *MockClient) TestI64(
	ctx context.Context,
	_Thing *int64,
	opts ...yarpc.CallOption,
) (success int64, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestI64", args...)
	success, _ = ret[i].(int64)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestI64(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestI64", args...)
}

// TestInsanity responds to a TestInsanity call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestInsanity(gomock.Any(), ...).Return(...)
//	... := client.TestInsanity(...)
func (m *MockClient) TestInsanity(
	ctx context.Context,
	_Argument *gauntlet.Insanity,
	opts ...yarpc.CallOption,
) (success map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity, err error) {

	args := []interface{}{ctx, _Argument}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestInsanity", args...)
	success, _ = ret[i].(map[gauntlet.UserId]map[gauntlet.Numberz]*gauntlet.Insanity)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestInsanity(
	ctx interface{},
	_Argument interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Argument}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestInsanity", args...)
}

// TestList responds to a TestList call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestList(gomock.Any(), ...).Return(...)
//	... := client.TestList(...)
func (m *MockClient) TestList(
	ctx context.Context,
	_Thing []int32,
	opts ...yarpc.CallOption,
) (success []int32, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestList", args...)
	success, _ = ret[i].([]int32)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestList(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestList", args...)
}

// TestMap responds to a TestMap call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestMap(gomock.Any(), ...).Return(...)
//	... := client.TestMap(...)
func (m *MockClient) TestMap(
	ctx context.Context,
	_Thing map[int32]int32,
	opts ...yarpc.CallOption,
) (success map[int32]int32, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestMap", args...)
	success, _ = ret[i].(map[int32]int32)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestMap(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestMap", args...)
}

// TestMapMap responds to a TestMapMap call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestMapMap(gomock.Any(), ...).Return(...)
//	... := client.TestMapMap(...)
func (m *MockClient) TestMapMap(
	ctx context.Context,
	_Hello *int32,
	opts ...yarpc.CallOption,
) (success map[int32]map[int32]int32, err error) {

	args := []interface{}{ctx, _Hello}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestMapMap", args...)
	success, _ = ret[i].(map[int32]map[int32]int32)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestMapMap(
	ctx interface{},
	_Hello interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Hello}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestMapMap", args...)
}

// TestMulti responds to a TestMulti call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestMulti(gomock.Any(), ...).Return(...)
//	... := client.TestMulti(...)
func (m *MockClient) TestMulti(
	ctx context.Context,
	_Arg0 *int8,
	_Arg1 *int32,
	_Arg2 *int64,
	_Arg3 map[int16]string,
	_Arg4 *gauntlet.Numberz,
	_Arg5 *gauntlet.UserId,
	opts ...yarpc.CallOption,
) (success *gauntlet.Xtruct, err error) {

	args := []interface{}{ctx, _Arg0, _Arg1, _Arg2, _Arg3, _Arg4, _Arg5}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestMulti", args...)
	success, _ = ret[i].(*gauntlet.Xtruct)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestMulti(
	ctx interface{},
	_Arg0 interface{},
	_Arg1 interface{},
	_Arg2 interface{},
	_Arg3 interface{},
	_Arg4 interface{},
	_Arg5 interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Arg0, _Arg1, _Arg2, _Arg3, _Arg4, _Arg5}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestMulti", args...)
}

// TestMultiException responds to a TestMultiException call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestMultiException(gomock.Any(), ...).Return(...)
//	... := client.TestMultiException(...)
func (m *MockClient) TestMultiException(
	ctx context.Context,
	_Arg0 *string,
	_Arg1 *string,
	opts ...yarpc.CallOption,
) (success *gauntlet.Xtruct, err error) {

	args := []interface{}{ctx, _Arg0, _Arg1}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestMultiException", args...)
	success, _ = ret[i].(*gauntlet.Xtruct)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestMultiException(
	ctx interface{},
	_Arg0 interface{},
	_Arg1 interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Arg0, _Arg1}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestMultiException", args...)
}

// TestNest responds to a TestNest call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestNest(gomock.Any(), ...).Return(...)
//	... := client.TestNest(...)
func (m *MockClient) TestNest(
	ctx context.Context,
	_Thing *gauntlet.Xtruct2,
	opts ...yarpc.CallOption,
) (success *gauntlet.Xtruct2, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestNest", args...)
	success, _ = ret[i].(*gauntlet.Xtruct2)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestNest(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestNest", args...)
}

// TestOneway responds to a TestOneway call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestOneway(gomock.Any(), ...).Return(...)
//	... := client.TestOneway(...)
func (m *MockClient) TestOneway(
	ctx context.Context,
	_SecondsToSleep *int32,
	opts ...yarpc.CallOption,
) (ack yarpc.Ack, err error) {

	args := []interface{}{ctx, _SecondsToSleep}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestOneway", args...)
	ack, _ = ret[i].(yarpc.Ack)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestOneway(
	ctx interface{},
	_SecondsToSleep interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _SecondsToSleep}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestOneway", args...)
}

// TestSet responds to a TestSet call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestSet(gomock.Any(), ...).Return(...)
//	... := client.TestSet(...)
func (m *MockClient) TestSet(
	ctx context.Context,
	_Thing map[int32]struct{},
	opts ...yarpc.CallOption,
) (success map[int32]struct{}, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestSet", args...)
	success, _ = ret[i].(map[int32]struct{})
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestSet(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestSet", args...)
}

// TestString responds to a TestString call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestString(gomock.Any(), ...).Return(...)
//	... := client.TestString(...)
func (m *MockClient) TestString(
	ctx context.Context,
	_Thing *string,
	opts ...yarpc.CallOption,
) (success string, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestString", args...)
	success, _ = ret[i].(string)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestString(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestString", args...)
}

// TestStringMap responds to a TestStringMap call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestStringMap(gomock.Any(), ...).Return(...)
//	... := client.TestStringMap(...)
func (m *MockClient) TestStringMap(
	ctx context.Context,
	_Thing map[string]string,
	opts ...yarpc.CallOption,
) (success map[string]string, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestStringMap", args...)
	success, _ = ret[i].(map[string]string)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestStringMap(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestStringMap", args...)
}

// TestStruct responds to a TestStruct call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestStruct(gomock.Any(), ...).Return(...)
//	... := client.TestStruct(...)
func (m *MockClient) TestStruct(
	ctx context.Context,
	_Thing *gauntlet.Xtruct,
	opts ...yarpc.CallOption,
) (success *gauntlet.Xtruct, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestStruct", args...)
	success, _ = ret[i].(*gauntlet.Xtruct)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestStruct(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestStruct", args...)
}

// TestTypedef responds to a TestTypedef call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestTypedef(gomock.Any(), ...).Return(...)
//	... := client.TestTypedef(...)
func (m *MockClient) TestTypedef(
	ctx context.Context,
	_Thing *gauntlet.UserId,
	opts ...yarpc.CallOption,
) (success gauntlet.UserId, err error) {

	args := []interface{}{ctx, _Thing}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestTypedef", args...)
	success, _ = ret[i].(gauntlet.UserId)
	i++
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestTypedef(
	ctx interface{},
	_Thing interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx, _Thing}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestTypedef", args...)
}

// TestVoid responds to a TestVoid call based on the mock expectations. This
// call will fail if the mock does not expect this call. Use EXPECT to expect
// a call to this function.
//
//	client.EXPECT().TestVoid(gomock.Any(), ...).Return(...)
//	... := client.TestVoid(...)
func (m *MockClient) TestVoid(
	ctx context.Context,
	opts ...yarpc.CallOption,
) (err error) {

	args := []interface{}{ctx}
	for _, o := range opts {
		args = append(args, o)
	}
	i := 0
	ret := m.ctrl.Call(m, "TestVoid", args...)
	err, _ = ret[i].(error)
	return
}

func (mr *_MockClientRecorder) TestVoid(
	ctx interface{},
	opts ...interface{},
) *gomock.Call {
	args := append([]interface{}{ctx}, opts...)
	return mr.mock.ctrl.RecordCall(mr.mock, "TestVoid", args...)
}
