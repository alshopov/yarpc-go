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

package testserviceserver

import (
	context "context"
	stream "go.uber.org/thriftrw/protocol/stream"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	test "go.uber.org/yarpc/encoding/thrift/internal/observabilitytest/test"
	yarpcerrors "go.uber.org/yarpc/yarpcerrors"
)

// Interface is the server-side interface for the TestService service.
type Interface interface {
	Call(
		ctx context.Context,
		Key string,
	) (string, error)
}

// New prepares an implementation of the TestService service for
// registration.
//
//	handler := TestServiceHandler{}
//	dispatcher.Register(testserviceserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "TestService",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "Call",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Unary,
					Unary:  thrift.UnaryHandler(h.Call),
					NoWire: call_NoWireHandler{impl},
				},
				Signature:    "Call(Key string) (string)",
				ThriftModule: test.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 1)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

type yarpcErrorNamer interface{ YARPCErrorName() string }

type yarpcErrorCoder interface{ YARPCErrorCode() *yarpcerrors.Code }

func (h handler) Call(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args test.TestService_Call_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'TestService' procedure 'Call': %w", err)
	}

	success, appErr := h.impl.Call(ctx, args.Key)

	hadError := appErr != nil
	result, err := test.TestService_Call_Helper.WrapResponse(success, appErr)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}

	return response, err
}

type call_NoWireHandler struct{ impl Interface }

func (h call_NoWireHandler) HandleNoWire(ctx context.Context, nwc *thrift.NoWireCall) (thrift.NoWireResponse, error) {
	var (
		args test.TestService_Call_Args
		rw   stream.ResponseWriter
		err  error
	)

	rw, err = nwc.RequestReader.ReadRequest(ctx, nwc.EnvelopeType, nwc.Reader, &args)
	if err != nil {
		return thrift.NoWireResponse{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode (via no wire) Thrift request for service 'TestService' procedure 'Call': %w", err)
	}

	success, appErr := h.impl.Call(ctx, args.Key)

	hadError := appErr != nil
	result, err := test.TestService_Call_Helper.WrapResponse(success, appErr)
	response := thrift.NoWireResponse{ResponseWriter: rw}
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}
	return response, err

}
