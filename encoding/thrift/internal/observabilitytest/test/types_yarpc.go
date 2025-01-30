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

package test

import yarpcerrors "go.uber.org/yarpc/yarpcerrors"

// YARPCErrorCode returns a yarpcerrors.CodeDataLoss for ExceptionWithCode.
//
// This is derived from the rpc.code annotation on the Thrift exception.
func (e *ExceptionWithCode) YARPCErrorCode() *yarpcerrors.Code {
	code := yarpcerrors.CodeDataLoss
	return &code
}

// Name is the error name for ExceptionWithCode.
func (e *ExceptionWithCode) YARPCErrorName() string { return "ExceptionWithCode" }

// YARPCErrorCode returns nil for ExceptionWithoutCode.
//
// This is derived from the rpc.code annotation on the Thrift exception.
func (e *ExceptionWithoutCode) YARPCErrorCode() *yarpcerrors.Code {

	return nil
}

// Name is the error name for ExceptionWithoutCode.
func (e *ExceptionWithoutCode) YARPCErrorName() string { return "ExceptionWithoutCode" }
