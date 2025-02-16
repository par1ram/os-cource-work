package internal

import (
	"net"
)

// ClientInterface определяет интерфейс клиента
type ClientInterface interface {
	Run()
	ConnectToServer() net.Conn
	SendRequest(text string)
}
