package internal

import (
	"fmt"
	"net"
	"strings"

	"github.com/par1ram/common"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/sirupsen/logrus"
)

// HandleConnection обрабатывает клиентское соединение для Server1.
// Теперь он использует common.ReceiveRequest для декодирования входящего JSON.
func (s *Server) HandleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		s.wg.Done()
	}()
	s.log.WithField("client", conn.RemoteAddr().String()).Info("Client connected")

	for {
		// Получаем запрос в формате JSON
		req, err := common.ReceiveRequest(conn)
		if err != nil {
			s.log.WithError(err).Error("Failed to receive request")
			break
		}
		s.log.WithFields(logrus.Fields{
			"client":  conn.RemoteAddr().String(),
			"command": req.Command,
		}).Info("Request received from client")

		// Обработка команды
		response := s.processRequest(req.Command)

		// Отправка ответа обратно клиенту в формате JSON
		if err := common.SendResponse(conn, response); err != nil {
			s.log.WithError(err).WithField("client", conn.RemoteAddr().String()).Error("Cannot send response to client")
			break
		}
	}

	s.log.WithField("client", conn.RemoteAddr().String()).Info("Client disconnected")
}

// processRequest обрабатывает команду клиента для Server1.
// Команда "memory" возвращает количество свободных байт физической памяти,
// а "mouse" возвращает количество клавиш мыши (например, 3).
func (s *Server) processRequest(command string) common.Response {
	cmd := strings.ToLower(strings.TrimSpace(command))
	switch cmd {
	case "memory":
		vmStat, err := mem.VirtualMemory()
		if err != nil {
			return common.Response{
				Status:  "error",
				Message: "Failed to retrieve memory info: " + err.Error(),
			}
		}
		return common.Response{
			Status:  "success",
			Message: fmt.Sprintf("Free memory: %d bytes", vmStat.Free),
		}
	case "mouse":
		return common.Response{
			Status:  "success",
			Message: "Mouse buttons: 3",
		}
	default:
		return common.Response{
			Status:  "error",
			Message: "Unknown command",
		}
	}
}
