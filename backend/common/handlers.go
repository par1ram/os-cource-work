package common

import (
	"encoding/json"
	"io"
)

// SendRequest - отправляет объект Request через соединение
func SendRequest(w io.Writer, req Request) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(req) // Автоматически добавляет `\n`
}

// ReceiveRequest - получает объект Request из соединения
func ReceiveRequest(r io.Reader) (Request, error) {
	var req Request
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&req) // Считывает JSON и десериализует
	return req, err
}

// SendResponse - отправляет объект Response через соединение
func SendResponse(w io.Writer, resp Response) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(resp)
}

// ReceiveResponse - получает объект Response из соединения
func ReceiveResponse(r io.Reader) (Response, error) {
	var resp Response
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&resp)
	return resp, err
}
