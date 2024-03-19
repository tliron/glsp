package server

import (
	"errors"
	"os"
	"strconv"
)

func (self *Server) RunNodeJs() error {
	nodeChannelFd := os.Getenv("NODE_CHANNEL_FD")
	if len(nodeChannelFd) == 0 {
		return errors.New("NODE_CHANNEL_FD not in environment")
	}
	nodeChannelFdInt, err := strconv.Atoi(nodeChannelFd)
	if err != nil {
		return err
	}
	file := os.NewFile(uintptr(nodeChannelFdInt), "/glsp/NODE_CHANNEL_FD")

	self.Log.Notice("listening for Node.js IPC connections")
	self.ServeStream(file, nil)
	return nil
}
