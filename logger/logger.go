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

// func NewMiddlewareLogger(logger logr.Logger) func(next http.Handler) http.Handler {
//	return middleware.RequestLogger(&MiddlewareLogger{logger})
// }
