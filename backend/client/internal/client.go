package internal

import (
	"net"
	"sync"

	"github.com/sirupsen/logrus"
)

type Client struct {
	log             *logrus.Logger
	conn            net.Conn
	messageChannel  chan string
	responseChannel chan string
	exitChannel     chan struct{}
	wg              sync.WaitGroup
}

func NewClient() *Client {
	return &Client{
		log:             NewLogger(),
		messageChannel:  make(chan string),
		responseChannel: make(chan string),
		exitChannel:     make(chan struct{}),
	}
}

func (c *Client) Run() {
	c.conn = c.ConnectToServer()
	defer c.conn.Close()

	// Запуск горутин
	c.startSendRoutine()
	c.startReceiveRoutine()

	// Запуск обработчика команд
	c.startInputHandler()
}

// SendRequest отправляет команду на сервер через канал
func (c *Client) SendRequest(text string) {
	c.messageChannel <- text
}
