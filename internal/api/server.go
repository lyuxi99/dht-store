package api

import (
	"DHT/internal/chord"
	"DHT/internal/logger"
	"fmt"
	"log"
	"net"
)

// ApiServer defines a server handling api requests of clients by creating .
type ApiServer struct {
	p2pServer *chord.P2pServer // the underlying chord.P2pServer of the ApiServer
	l         net.Listener     // the net.Listener that the ApiServer is listening on
	stopped   bool             // whether the ApiServer has stopped
}

// NewApiServer creates a ApiServer with the given underlying chord.P2pServer, listening on the given address.
func NewApiServer(p2pServer *chord.P2pServer, address string) *ApiServer {
	l, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatal("ApiServer net.Listen error", err)
	}
	return &ApiServer{
		p2pServer: p2pServer,
		l:         l,
		stopped:   false,
	}
}

// Serve accepts incoming connections.
func (s *ApiServer) Serve() error {
	for {
		c, err := s.l.Accept()
		if err != nil {
			if s.stopped {
				fmt.Println("ApiServer stopped!")
				return nil
			}
			return err
		}
		logger.Logger.Infow("connection accepted", "addr", c.RemoteAddr())
		go NewConnection(s, c).threadReceiveMsg()
	}
}

// Stop stops the ApiServer.
func (s *ApiServer) Stop() {
	s.stopped = true
	s.l.Close()
}
