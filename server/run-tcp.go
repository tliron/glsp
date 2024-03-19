package server

import (
	"github.com/tliron/commonlog"
)

func (self *Server) RunTCP(address string) error {
	listener, err := self.newNetworkListener("tcp", address)
	if err != nil {
		return err
	}

	log := commonlog.NewKeyValueLogger(self.Log, "address", address)
	defer commonlog.CallAndLogError((*listener).Close, "listener.Close", log)
	log.Notice("listening for TCP connections")

	var connectionCount uint64

	for {
		connection, err := (*listener).Accept()
		if err != nil {
			return err
		}

		connectionCount++
		connectionLog := commonlog.NewKeyValueLogger(log, "id", connectionCount)

		go self.ServeStream(connection, connectionLog)
	}
}
