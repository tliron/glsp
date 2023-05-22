package server

import (
	"strings"

	"github.com/tliron/commonlog"
)

type Logger struct {
	log commonlog.Logger
}

// jsonrpc2.Logger interface
func (self *Logger) Printf(format string, v ...any) {
	self.log.Debugf(strings.TrimSuffix(format, "\n"), v...)
}
