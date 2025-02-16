package internal

import (
	"net"
	"sync"

	"github.com/sirupsen/logrus"
)

// Server - структура сервера
type Server struct {
	log      *logrus.Logger
	port     string
	listener net.Listener
	wg       sync.WaitGroup
}

// NewServer создает новый сервер
func NewServer() *Server {
	return &Server{
		log:  NewLogger(),
		port: "8100", // Можно вынести в config.go
	}
}

// Run запускает сервер
func (s *Server) Run() {
	var err error
	s.listener, err = net.Listen("tcp", ":"+s.port)
	if err != nil {
		s.log.WithError(err).Fatal("Error starting server on port ", s.port)
	}
	defer s.listener.Close()

	s.log.Info("Server1 started on port ", s.port)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			s.log.WithError(err).Error("Error accepting connection")
			continue
		}
		s.wg.Add(1)
		go s.HandleConnection(conn)
	}
}
