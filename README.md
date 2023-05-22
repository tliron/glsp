*This is an early release. Some features are not yet fully implemented.*

GLSP
====

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Go Reference](https://pkg.go.dev/badge/github.com/tliron/glsp.svg)](https://pkg.go.dev/github.com/tliron/kutglspil)
[![Go Report Card](https://goreportcard.com/badge/github.com/tliron/glsp)](https://goreportcard.com/report/github.com/tliron/glsp)

[Language Server Protocol](https://microsoft.github.io/language-server-protocol/) SDK for Go.

It enables you to more easily implement language servers by writing them in Go. GLSP contains:

1) all the message structures for easy serialization,
2) a handler for all client methods, and
3) a ready-to-run JSON-RPC 2.0 server supporting stdio, TCP, WebSockets, and Node.js IPC.

All you need to do, then, is provide the features for the language you want to support.

Projects using GLSP:

* [Puccini TOSCA Language Server](https://github.com/tliron/puccini-language-server)
* [zk](https://github.com/mickael-menu/zk)


References
----------

* [go-lsp](https://github.com/sourcegraph/go-lsp) is another implementation with reduced coverage of the protocol


Minimal Example
---------------

```go
package main

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
	"github.com/tliron/commonlog"

	// Must include a backend implementation
	// See CommonLog for other options: https://github.com/tliron/commonlog
	_ "github.com/tliron/commonlog/simple"
)

const lsName = "my language"

var version string = "0.0.1"
var handler protocol.Handler

func main() {
	// This increases logging verbosity (optional)
	commonlog.Configure(1, nil)

	handler = protocol.Handler{
		Initialize:  initialize,
		Initialized: initialized,
		Shutdown:    shutdown,
		SetTrace:    setTrace,
	}

	server := server.NewServer(&handler, lsName, false)

	server.RunStdio()
}

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	capabilities := handler.CreateServerCapabilities()

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    lsName,
			Version: &version,
		},
	}, nil
}

func initialized(context *glsp.Context, params *protocol.InitializedParams) error {
	return nil
}

func shutdown(context *glsp.Context) error {
	protocol.SetTraceValue(protocol.TraceValueOff)
	return nil
}

func setTrace(context *glsp.Context, params *protocol.SetTraceParams) error {
	protocol.SetTraceValue(params.Value)
	return nil
}
```
