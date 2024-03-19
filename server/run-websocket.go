package server

import (
	"net/http"
	"sync/atomic"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/tliron/commonlog"
)

func (self *Server) RunWebSocket(address string) error {
	mux := http.NewServeMux()
	upgrader := websocket.Upgrader{CheckOrigin: func(request *http.Request) bool { return true }}

	var connectionCount uint64

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		connection, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			self.Log.Warningf("error upgrading HTTP to web socket: %s", err.Error())
			http.Error(writer, errors.Wrap(err, "could not upgrade to web socket").Error(), http.StatusBadRequest)
			return
		}

		log := commonlog.NewKeyValueLogger(self.Log, "id", atomic.AddUint64(&connectionCount, 1))
		defer commonlog.CallAndLogError(connection.Close, "connection.Close", log)
		self.ServeWebSocket(connection, log)
	})

	listener, err := self.newNetworkListener("tcp", address)
	if err != nil {
		return err
	}

	server := http.Server{
		Handler:      http.TimeoutHandler(mux, self.Timeout, ""),
		ReadTimeout:  self.ReadTimeout,
		WriteTimeout: self.WriteTimeout,
	}

	self.Log.Notice("listening for web socket connections", "address", address)
	err = server.Serve(*listener)
	return errors.Wrap(err, "WebSocket")
}
