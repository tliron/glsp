package server

import (
	contextpkg "context"
	"time"

	"github.com/tliron/glsp"
	"github.com/tliron/kutil/logging"
)

//
// Server
//

type Server struct {
	Handler     glsp.Handler
	LogBaseName string
	Debug       bool

	Log          logging.Logger
	Context      contextpkg.Context
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServer(handler glsp.Handler, logBaseName string, debug bool) *Server {
	return &Server{
		Handler:      handler,
		LogBaseName:  logBaseName,
		Debug:        debug,
		Log:          logging.GetLoggerf("%s.server", logBaseName),
		Context:      contextpkg.Background(),
		ReadTimeout:  75 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
}
