# yarpc

[![GoDoc][doc-img]][doc]
[![GitHub release][release-img]][release]
[![Mit License][mit-img]][mit]
[![Build Status][ci-img]][ci]
[![Coverage Status][cov-img]][cov]

A message passing platform for Go that lets you:

* Write servers and clients with various encodings, including [JSON](http://www.json.org/), [Thrift](https://thrift.apache.org/), and [Protobuf](https://developers.google.com/protocol-buffers/).
* Expose servers over many transports simultaneously, including [HTTP/1.1](https://www.w3.org/Protocols/rfc2616/rfc2616.html), [gRPC](https://grpc.io/), and [TChannel](https://github.com/uber/tchannel).
* Migrate outbound calls between transports without any code changes using config.

## Installation

Add dependency to your `go.mod` file:

```
go get go.uber.org/yarpc@latest
```

Use the following import path in the code:

```go
import "go.uber.org/yarpc"
```

Please see [reference][doc] and [examples][examples-link] for more details.

## Stability

This library is `v1` and follows [SemVer](http://semver.org/) strictly.

No breaking changes will be made to exported APIs before `v2.0.0` with the
**exception of experimental packages**.

Experimental packages reside within packages named `x`, and are *not stable*. This means their
APIs can break at any time. The intention here is to validate these APIs and iterate on them
by working closely with internal customers. Once stable, their contents will be moved out of
the containing `x` package and their APIs will be locked.

[doc-img]: https://pkg.go.dev/badge/go.uber.org/yarpc.svg
[doc]: https://pkg.go.dev/go.uber.org/yarpc

[release-img]: https://img.shields.io/github/release/yarpc/yarpc-go.svg
[release]: https://github.com/yarpc/yarpc-go/releases

[mit-img]: http://img.shields.io/badge/License-MIT-blue.svg
[mit]: https://github.com/yarpc/yarpc-go/blob/master/LICENSE

[ci-img]: https://badge.buildkite.com/f7d8e675c4d5ee4f5c4e4c2e33ca03c5be9bde22b186750538.svg?branch=master
[ci]: https://buildkite.com/uberopensource/yarpc-go

[cov-img]: https://codecov.io/gh/yarpc/yarpc-go/branch/master/graph/badge.svg
[cov]: https://codecov.io/gh/yarpc/yarpc-go/branch/master

[examples-link]: https://github.com/yarpc/yarpc-go/tree/dev/internal/examples