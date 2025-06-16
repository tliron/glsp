package protocol

import (
	"encoding/json"

	"github.com/tliron/glsp"
)

type CustomRequestHandler struct {
	Func CustomRequestFunc
	// This field should be private however it is used in both versions of the protocol
	Params json.RawMessage
}

type CustomRequestHandlers map[string]CustomRequestHandler

type CustomRequestFunc func(context *glsp.Context, params json.RawMessage) (any, error)
