package server

import (
	"strings"

	"github.com/tliron/commonlog"
)

type JSONRPCLogger struct {
	log commonlog.Logger
}

// ([jsonrpc2.Logger] interface)
func (self *JSONRPCLogger) Printf(format string, v ...any) {
	self.log.Debugf(strings.TrimSuffix(format, "\n"), v...)
}
