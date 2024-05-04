package server

import (
	"fmt"
	"net"

	"github.com/alex-laycalvert/chattn/internal/common"
)

type Server struct {
	port    int
	clients []*net.Conn
}

func New(port int) Server {
	clients := make([]*net.Conn, 0)
	return Server{port, clients}
}

func (s *Server) Start() error {
	addr := net.TCPAddr{IP: net.IPv4(127, 0, 0, 1), Port: s.port}
	listener, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			listener.Close()
			return err
		}
		s.clients = append(s.clients, &conn)
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	buf := make([]byte, common.BUF_SIZE)
	for {
		// TODO: error handling
		bytesRead, _ := conn.Read(buf)
		if bytesRead == 0 {
			return
		}
		msg := buf[0:bytesRead]
		fmt.Println(string(msg))
	}
}
