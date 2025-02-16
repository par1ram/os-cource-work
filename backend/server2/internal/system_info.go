package internal

import (
	"fmt"
	"runtime"
	"strconv"
)

// GetThreadCount возвращает количество горутин (пример)
// Заметим, что для получения количества OS-потоков понадобится другой подход.
func GetThreadCount() int {
	return runtime.NumGoroutine()
}

// MoveWindow симулирует перемещение окна
func MoveWindow(xStr, yStr string) error {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		return fmt.Errorf("invalid x coordinate")
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		return fmt.Errorf("invalid y coordinate")
	}
	// Здесь должен быть вызов API ОС для перемещения окна.
	// Для демонстрации просто выводим координаты.
	fmt.Printf("Moving window to coordinates (%d, %d)\n", x, y)
	return nil
}
