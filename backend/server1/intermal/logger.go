package internal

import (
	"os"

	"github.com/sirupsen/logrus"
)

// NewLogger создает и настраивает новый логгер
func NewLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	return log
}
