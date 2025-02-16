package common

// Request - структура запроса от клиента
type Request struct {
	Command string `json:"command"`
	Payload string `json:"payload,omitempty"`
}

// Response - структура ответа сервера
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
