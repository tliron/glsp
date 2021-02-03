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

func (self *Server) handle(context contextpkg.Context, connection *jsonrpc2.Conn, request *jsonrpc2.Request) (interface{}, error) {
	// glsp.NotifyFunc signature
	glspContext := glsp.Context{
		Method: request.Method,
		Params: *request.Params,
		Notify: func(method string, params interface{}) {
			err := connection.Notify(context, method, params)
			if err != nil {
				self.Log.Errorf("%s", err.Error())
			}
		},
	}

	switch request.Method {
	case "exit":
		connection.Close()
		return nil, nil

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
