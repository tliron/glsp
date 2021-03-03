package server

import (
	"io"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	wsjsonrpc2 "github.com/sourcegraph/jsonrpc2/websocket"
	"github.com/tliron/kutil/logging"
)

// See: https://github.com/sourcegraph/go-langserver/blob/master/main.go#L179

func (self *Server) serveStream(stream io.ReadWriteCloser) {
	handler := self.newHandler()
	connectionOptions := self.newConnectionOptions()
	<-jsonrpc2.NewConn(self.Context, jsonrpc2.NewBufferedStream(stream, jsonrpc2.VSCodeObjectCodec{}), handler, connectionOptions...).DisconnectNotify()
}

func (self *Server) serveWebSocket(socket *websocket.Conn) {
	handler := self.newHandler()
	connectionOptions := self.newConnectionOptions()
	<-jsonrpc2.NewConn(self.Context, wsjsonrpc2.NewObjectStream(socket), handler, connectionOptions...).DisconnectNotify()
}

func (self *Server) serveStreamAsync(stream io.ReadWriteCloser, id int) {
	handler := self.newHandler()
	connectionOptions := self.newConnectionOptions()
	connection := jsonrpc2.NewConn(self.Context, jsonrpc2.NewBufferedStream(stream, jsonrpc2.VSCodeObjectCodec{}), handler, connectionOptions...)
	go func() {
		<-connection.DisconnectNotify()
		self.Log.Infof("connection #%d closed", id)
	}()
}

func (self *Server) newConnectionOptions() []jsonrpc2.ConnOpt {
	if self.Debug {
		logger := logging.GetLogger(self.LogBaseName + ".rpc")
		return []jsonrpc2.ConnOpt{jsonrpc2.LogMessages(&Logger{logger})}
	} else {
		return nil
	}
}
