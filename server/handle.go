package server

import (
	contextpkg "context"
	"fmt"

	"github.com/sourcegraph/jsonrpc2"
	"github.com/tliron/glsp"
)

// See: https://github.com/sourcegraph/go-langserver/blob/master/langserver/handler.go#L206

func (self *Server) newHandler() jsonrpc2.Handler {
	return jsonrpc2.HandlerWithError(self.handle)
}

func (self *Server) handle(context contextpkg.Context, connection *jsonrpc2.Conn, request *jsonrpc2.Request) (any, error) {
	glspContext := glsp.Context{
		Method: request.Method,
		Notify: func(method string, params any) {
			if err := connection.Notify(context, method, params); err != nil {
				self.Log.Errorf("%s", err.Error())
			}
		},
		Call: func(method string, params any, result any) {
			if err := connection.Call(context, method, params, result); err != nil {
				self.Log.Errorf("%s", err.Error())
			}
		},
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
		// Note: jsonrpc2 will not even call this function if reqest.Params is not valid JSON,
		// so we don't need to handle jsonrpc2.CodeParseError here
		r, validMethod, validParams, err := self.Handler.Handle(&glspContext)
		if !validMethod {
			return nil, &jsonrpc2.Error{
				Code:    jsonrpc2.CodeMethodNotFound,
				Message: fmt.Sprintf("method not supported: %s", request.Method),
			}
		} else if !validParams {
			if err != nil {
				return nil, &jsonrpc2.Error{
					Code:    jsonrpc2.CodeInvalidParams,
					Message: err.Error(),
				}
			} else {
				return nil, &jsonrpc2.Error{
					Code: jsonrpc2.CodeInvalidParams,
				}
			}
		} else if err != nil {
			return nil, &jsonrpc2.Error{
				Code:    jsonrpc2.CodeInvalidRequest,
				Message: err.Error(),
			}
		} else {
			return r, nil
		}
	}
}
