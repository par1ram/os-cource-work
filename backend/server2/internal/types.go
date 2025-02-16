package internal

import "net"

// ServerInterface определяет интерфейс сервера
type ServerInterface interface {
	Run()
	HandleConnection(conn net.Conn)
}
