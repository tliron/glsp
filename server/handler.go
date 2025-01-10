package server

import (
	contextpkg "context"
	"fmt"
	"sync"

	"github.com/sourcegraph/jsonrpc2"

	"github.com/tliron/glsp"
)

// See: https://github.com/sourcegraph/go-langserver/blob/master/langserver/handler.go#L206
func (self *Server) newHandler() jsonrpc2.Handler {
	return newLSPHandler(
		jsonrpc2.HandlerWithError(self.handle),
	)
}

// newLSPHandler returns a handler that processes each request goes in its own
// goroutine, processing requests in a FIFO fashion besides $/cancelRequest, which are not queued.
// It allows unbounded goroutines, all stalled on the previous one.
func newLSPHandler(handler jsonrpc2.Handler) jsonrpc2.Handler {
	head := make(chan struct{})
	close(head)
	return &lspHandler{
		wrapped: handler,
		head:    head,
	}
}

type lspHandler struct {
	wrapped jsonrpc2.Handler
	head    chan struct{}
	mx      sync.Mutex
}

func (a *lspHandler) Handle(ctx contextpkg.Context, conn *jsonrpc2.Conn, request *jsonrpc2.Request) {
	// for cancel requests, allow preemption, and don't consider it part of the request queue
	if request.Method == "$/cancelRequest" {
		go a.wrapped.Handle(ctx, conn, request)
		return
	}

	a.mx.Lock()
	previous := a.head
	thisReq := make(chan struct{})
	a.head = thisReq
	a.mx.Unlock()

	go func() {
		defer close(thisReq)
		<-previous
		a.wrapped.Handle(ctx, conn, request)
	}()
}

func (self *Server) handle(context contextpkg.Context, connection *jsonrpc2.Conn, request *jsonrpc2.Request) (any, error) {
	glspContext := glsp.Context{
		Method:    request.Method,
		RequestID: request.ID,
		Notify: func(method string, params any) {
			if err := connection.Notify(context, method, params); err != nil {
				self.Log.Error(err.Error())
			}
		},
		Call: func(method string, params any, result any) {
			if err := connection.Call(context, method, params, result); err != nil {
				self.Log.Error(err.Error())
			}
		},
		Context: context,
	}

	if request.Params != nil {
		glspContext.Params = *request.Params
	}

	switch request.Method {
	case "exit":
		// We're giving the attached handler a chance to handle it first, but we'll ignore any result
		self.Handler.Handle(&glspContext)
		err := connection.Close()
		return nil, err

	default:
		// Note: jsonrpc2 will not even call this function if reqest.Params is invalid JSON,
		// so we don't need to handle jsonrpc2.CodeParseError here
		result, validMethod, validParams, err := self.Handler.Handle(&glspContext)
		if !validMethod {
			return nil, &jsonrpc2.Error{
				Code:    jsonrpc2.CodeMethodNotFound,
				Message: fmt.Sprintf("method not supported: %s", request.Method),
			}
		} else if !validParams {
			if err == nil {
				return nil, &jsonrpc2.Error{
					Code: jsonrpc2.CodeInvalidParams,
				}
			} else {
				return nil, &jsonrpc2.Error{
					Code:    jsonrpc2.CodeInvalidParams,
					Message: err.Error(),
				}
			}
		} else if err != nil {
			return nil, &jsonrpc2.Error{
				Code:    jsonrpc2.CodeInvalidRequest,
				Message: err.Error(),
			}
		} else {
			return result, nil
		}
	}
}
