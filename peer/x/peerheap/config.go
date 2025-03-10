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

package peerheap

import (
	"go.uber.org/yarpc/api/peer"
	"go.uber.org/yarpc/yarpcconfig"
)

// Spec returns a configuration specification for the least-pending peer heap
// peer chooser implementation, making it possible to select the least pending
// peer with transports that use outbound peer list configuration (like HTTP).
//
//	cfg := yarpcconfig.New()
//	cfg.MustRegisterPeerList(peerheap.Spec())
//
// This enables the least-pending peer list:
//
//	outbounds:
//	  otherservice:
//	    unary:
//	      http:
//	        url: https://host:port/rpc
//	        least-pending:
//	          peers:
//	            - 127.0.0.1:8080
//	            - 127.0.0.1:8081
func Spec() yarpcconfig.PeerListSpec {
	return yarpcconfig.PeerListSpec{
		Name: "least-pending",
		BuildPeerList: func(c struct{}, t peer.Transport, k *yarpcconfig.Kit) (peer.ChooserList, error) {
			return New(t), nil
		},
	}
}
