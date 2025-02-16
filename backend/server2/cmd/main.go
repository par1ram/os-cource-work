package main

import (
	"github.com/par1ram/server2/internal"
)

func main() {
	// Создаем и запускаем сервер
	server := internal.NewServer2()
	server.Run()
}
