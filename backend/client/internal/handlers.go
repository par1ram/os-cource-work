package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/par1ram/common"
)

// startSendRoutine запускает горутину для отправки сообщений
func (c *Client) startSendRoutine() {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			select {
			case message, ok := <-c.messageChannel:
				if !ok {
					return
				}
				c.sendToServer(message)
			case <-c.exitChannel:
				return
			}
		}
	}()
}

// sendToServer отправляет сообщение на сервер
func (c *Client) sendToServer(message string) {
	req := common.Request{Command: message}
	if err := common.SendRequest(c.conn, req); err != nil {
		c.log.WithError(err).Error("Failed to send message")
	}
}

// startReceiveRoutine запускает горутину для приема сообщений
func (c *Client) startReceiveRoutine() {
	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			resp, err := common.ReceiveResponse(c.conn)
			if err != nil {
				c.log.WithError(err).Error("Failed to read response")
				return
			}
			c.responseChannel <- resp.Message
		}
	}()
}

// startInputHandler запускает обработку ввода пользователя
func (c *Client) startInputHandler() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter command (or 'exit' to quit): ")
		if !scanner.Scan() {
			c.log.Info("No more input detected. Exiting.")
			break
		}

		text := scanner.Text()
		if strings.ToLower(text) == "exit" {
			c.log.Info("Exiting")
			close(c.messageChannel)
			close(c.exitChannel)
			break
		}

		// Отправляем команду через канал
		c.SendRequest(text)

		// Получаем ответ от сервера (асинхронно)
		select {
		case response := <-c.responseChannel:
			c.log.WithField("response", response).Info("Received response")
		}
	}

	// Ожидание завершения всех горутин
	c.wg.Wait()
	c.log.Info("Program finished")
}
