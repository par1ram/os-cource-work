package internal

import (
	"fmt"
	"net"
)

// ConnectToServer подключается к серверу
func (c *Client) ConnectToServer() net.Conn {
	serverAddr := GetServerChoice()
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		c.log.WithError(err).Fatal("Error connecting to server")
	}
	c.log.Info("Connected to server ", serverAddr)
	return conn
}

// GetServerChoice спрашивает, к какому серверу подключиться
func GetServerChoice() string {
	fmt.Println("Choose server:")
	fmt.Println("1 - Server1 (Text Processor)")
	fmt.Println("2 - Server2 (System Info)")

	var choice string
	fmt.Scanln(&choice)

	// В Docker используем имена сервисов из docker-compose.yml
	if choice == "2" {
		return "server2:8200"
	}
	return "server1:8100"
}
