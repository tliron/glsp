package server

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

func (self *Server) RunWebSocket(address string) error {
	mux := http.NewServeMux()
	upgrader := websocket.Upgrader{CheckOrigin: func(request *http.Request) bool { return true }}

	connectionCount := 0

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		connection, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			self.Log.Infof("error upgrading HTTP to WebSocket: %s", err.Error())
			http.Error(writer, errors.Wrap(err, "could not upgrade to WebSocket").Error(), http.StatusBadRequest)
			return
		}
		defer connection.Close()

		connectionCount++
		connectionId := connectionCount

		self.Log.Infof("received incoming WebSocket connection #%d", connectionId)
		self.ServeWebSocket(connection)
		self.Log.Infof("WebSocket connection #%d closed", connectionId)
	})

	listener, err := self.newNetworkListener("tcp", address)
	if err != nil {
		return err
	}

	server := http.Server{
		Handler:      mux,
		ReadTimeout:  self.ReadTimeout,
		WriteTimeout: self.WriteTimeout,
	}

	self.Log.Infof("listening for WebSocket connections on %s", address)
	err = server.Serve(*listener)
	return errors.Wrap(err, "WebSocket")
}
