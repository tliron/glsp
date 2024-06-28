package protocol

import (
	"encoding/json"

	"github.com/tliron/glsp"
)

type CustomRequestHandler struct {
	Method string
	Func   CustomRequestFunc
	params json.RawMessage
}

type CustomRequestFunc func(context *glsp.Context, params json.RawMessage) error
