package glsp

import (
	"encoding/json"
)

type NotifyFunc func(method string, params any)
type CallFunc func(method string, params any, result any)

type Context struct {
	Method string
	Params json.RawMessage
	Notify NotifyFunc
	Call   CallFunc
}

type Handler interface {
	Handle(context *Context) (r any, validMethod bool, validParams bool, err error)
}
