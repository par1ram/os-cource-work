package internal

import (
	"net"
	"sync"

	"github.com/sirupsen/logrus"
)

// Server2 - структура сервера
type Server2 struct {
	log      *logrus.Logger
	port     string
	listener net.Listener
	wg       sync.WaitGroup
}

// NewServer2 создает новый сервер
func NewServer2() *Server2 {
	return &Server2{
		log:  NewLogger(),
		port: "8200", // Можно вынести в config.go
	}
}

// Run запускает сервер
func (s *Server2) Run() {
	var err error
	s.listener, err = net.Listen("tcp", ":"+s.port)
	if err != nil {
		s.log.WithError(err).Fatal("Error starting server on port ", s.port)
	}
	defer s.listener.Close()

	s.log.Info("Server2 started on port ", s.port)

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

// Close закрывает сервер и ждет завершения всех горутин
func (s *Server2) Close() {
	s.listener.Close()
	s.wg.Wait()
	s.log.Info("Server2 shutdown completed")
}
