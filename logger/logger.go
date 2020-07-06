package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func New() *logrus.Logger {
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Level = logrus.DebugLevel
	return logger
}
