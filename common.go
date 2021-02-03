package glsp

import (
	"encoding/json"
)

type NotifyFunc func(method string, params interface{})

type Context struct {
	Method string
	Params json.RawMessage
	Notify NotifyFunc
}

type Handler interface {
	Handle(context *Context) (r interface{}, validMethod bool, validParams bool, err error)
}
