package tcp

import (
	"net"
	"sync"
)

type Server struct {
	Listener net.Listener
	Quit     chan interface{}
	Wg       sync.WaitGroup
}

// OpenConnection Opens a new listener tcp port and return Server interface
func OpenConnection(PORT string) (Server, error) {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		return Server{}, err
	}
	return Server{
		Listener: listener,
	}, nil
}

// Connect Used to connect to a existing open tcp connection and returns the connection interface net.Conn
func Connect(addr string) (net.Conn, error) {
	return net.Dial("tcp", addr)
}

// AcceptNewConnection Accept new tcp connections on Server
func (s *Server) AcceptNewConnection() (net.Conn, error) {
	return s.Listener.Accept()
}
