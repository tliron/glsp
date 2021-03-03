package server

import (
	"crypto/tls"
	"net"
	"os"
)

func (self *Server) newNetworkListener(network string, address string) (*net.Listener, error) {
	listener, err := net.Listen(network, address)
	if err != nil {
		self.Log.Criticalf("could not bind to address %s: %v", address, err)
		return nil, err
	}

	cert := os.Getenv("TLS_CERT")
	key := os.Getenv("TLS_KEY")
	if (cert != "") && (key != "") {
		cert, err := tls.X509KeyPair([]byte(cert), []byte(key))
		if err != nil {
			return nil, err
		}
		listener = tls.NewListener(listener, &tls.Config{
			Certificates: []tls.Certificate{cert},
		})
	}

	return &listener, nil
}
