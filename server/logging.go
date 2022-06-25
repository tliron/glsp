package server

import (
	"strings"

	"github.com/tliron/kutil/logging"
)

type Logger struct {
	log logging.Logger
}

// jsonrpc2.Logger interface
func (self *Logger) Printf(format string, v ...any) {
	self.log.Debugf(strings.TrimSuffix(format, "\n"), v...)
}
