package server

func (self *Server) RunTCP(address string) error {
	listener, err := self.newNetworkListener("tcp", address)
	if err != nil {
		return err
	}
	defer (*listener).Close()

	self.Log.Infof("listening for TCP connections on %s", address)

	connectionCount := 0

	for {
		connection, err := (*listener).Accept()
		if err != nil {
			return err
		}

		connectionCount++
		connectionId := connectionCount

		self.Log.Infof("received incoming TCP connection #%d", connectionId)
		self.serveStreamAsync(connection, connectionId)
	}
}
