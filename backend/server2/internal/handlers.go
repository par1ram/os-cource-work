package internal

import (
	"net"
	"strconv"
	"strings"

	"github.com/par1ram/common"
	"github.com/sirupsen/logrus"
)

// HandleConnection обрабатывает клиентское соединение для Server2, ожидая JSON-запросы.
func (s *Server2) HandleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		s.wg.Done()
	}()
	s.log.WithField("client", conn.RemoteAddr().String()).Info("Client connected")

	for {
		// Принимаем запрос как JSON
		req, err := common.ReceiveRequest(conn)
		if err != nil {
			s.log.WithError(err).Error("Failed to receive request")
			break
		}
		s.log.WithFields(logrus.Fields{
			"client":  conn.RemoteAddr().String(),
			"command": req.Command,
		}).Info("Request received from client")

		response := s.processRequest(req.Command)
		if err := common.SendResponse(conn, response); err != nil {
			s.log.WithError(err).WithField("client", conn.RemoteAddr().String()).Error("Cannot send response to client")
			break
		}
	}

	s.log.WithField("client", conn.RemoteAddr().String()).Info("Client disconnected")
}

func (s *Server2) processRequest(command string) common.Response {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return common.Response{
			Status:  "error",
			Message: "Empty command",
		}
	}
	switch strings.ToLower(parts[0]) {
	case "threads":
		return common.Response{
			Status:  "success",
			Message: strconv.Itoa(GetThreadCount()),
		}
	case "move_window":
		if len(parts) < 3 {
			return common.Response{
				Status:  "error",
				Message: "Usage: move_window x y",
			}
		}
		x, y := parts[1], parts[2]
		err := MoveWindow(x, y)
		if err != nil {
			return common.Response{
				Status:  "error",
				Message: "Failed to move window: " + err.Error(),
			}
		}
		return common.Response{
			Status:  "success",
			Message: "Window moved successfully",
		}
	default:
		return common.Response{
			Status:  "error",
			Message: "Unknown command",
		}
	}
}
