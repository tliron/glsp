package server

import (
	"io"

	"github.com/gorilla/websocket"
	"github.com/sourcegraph/jsonrpc2"
	wsjsonrpc2 "github.com/sourcegraph/jsonrpc2/websocket"
	"github.com/tliron/commonlog"
)

// See: https://github.com/sourcegraph/go-langserver/blob/master/main.go#L179

func (self *Server) getStreamConn(stream io.ReadWriteCloser) *jsonrpc2.Conn {
	handler := self.newHandler()
	connectionOptions := self.newConnectionOptions()
	return jsonrpc2.NewConn(self.Context, jsonrpc2.NewBufferedStream(stream, jsonrpc2.VSCodeObjectCodec{}), handler, connectionOptions...)
}

func (self *Server) serveStream(stream io.ReadWriteCloser) {
	<-self.getStreamConn(stream).DisconnectNotify()
}

func (self *Server) getWebSocketConn(socket *websocket.Conn) *jsonrpc2.Conn {
	handler := self.newHandler()
	connectionOptions := self.newConnectionOptions()
	return jsonrpc2.NewConn(self.Context, wsjsonrpc2.NewObjectStream(socket), handler, connectionOptions...)
}

func (self *Server) ServeWebSocket(socket *websocket.Conn) {
	<-self.getWebSocketConn(socket).DisconnectNotify()
}

func (self *Server) serveStreamAsync(stream io.ReadWriteCloser, id int) {
	go func() {
		<-self.getStreamConn(stream).DisconnectNotify()
		self.Log.Infof("connection #%d closed", id)
	}()
}

func (self *Server) newConnectionOptions() []jsonrpc2.ConnOpt {
	if self.Debug {
		logger := commonlog.GetLogger(self.LogBaseName + ".rpc")
		return []jsonrpc2.ConnOpt{jsonrpc2.LogMessages(&Logger{logger})}
	} else {
		return nil
	}
}
