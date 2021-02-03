package server

import (
	contextpkg "context"
	"time"

	"github.com/op/go-logging"
	"github.com/tliron/glsp"
)

//
// Server
//

type Server struct {
	Handler glsp.Handler
	Log     *logging.Logger
	Debug   bool

	Context      contextpkg.Context
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServer(handler glsp.Handler, logName string, debug bool) *Server {
	return &Server{
		Handler:      handler,
		Log:          logging.MustGetLogger(logName),
		Debug:        debug,
		Context:      contextpkg.Background(),
		ReadTimeout:  75 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
}
