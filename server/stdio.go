package server

import (
	"os"
)

func (self *Server) RunStdio() error {
	self.Log.Info("reading from stdin, writing to stdout")
	self.serveStream(stdrwc{})
	self.Log.Info("stdin/stdout connection closed")
	return nil
}

type stdrwc struct{}

// io.ReadWriteCloser interface
func (stdrwc) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

// io.ReadWriteCloser interface
func (stdrwc) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

// io.ReadWriteCloser interface
func (stdrwc) Close() error {
	if err := os.Stdin.Close(); err == nil {
		return os.Stdout.Close()
	} else {
		return err
	}
}
