package peud_middlware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

type middlewareLogger struct {
	logger *logrus.Entry
}

type middlewareLoggerEntry struct {
	logger *logrus.Entry
}

func NewMiddlewareLogger(logger *logrus.Entry) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(&middlewareLogger{logger: logger})
}

func (c middlewareLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &middlewareLoggerEntry{c.logger}
	start := time.Now()
	var requestID string
	if reqID := r.Context().Value(middleware.RequestIDKey); reqID != nil {
		requestID = reqID.(string)
	}

	latency := time.Since(start)

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	fields := logrus.Fields{
		"proto":      r.Proto,
		"took":       latency.Nanoseconds(),
		"remote":     r.RemoteAddr,
		"scheme":     scheme,
		"uri":        fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI),
		"endpoint":   r.RequestURI,
		"method":     r.Method,
		"user_agent": r.UserAgent(),
	}
	if requestID != "" {
		fields["request-id"] = requestID
	}
	c.logger.WithFields(fields).Info("request started")

	return entry
}

func (c *middlewareLoggerEntry) Write(status, bytes int, h http.Header, elapsed time.Duration, extra interface{}) {
	c.logger = c.logger.WithFields(logrus.Fields{
		"status": status,
		"took":   elapsed,
		"bytes":  bytes,
	})
	c.logger.Info("request complete")
}

func (c *middlewareLoggerEntry) Panic(v interface{}, stack []byte) {
	c.logger = c.logger.WithFields(logrus.Fields{
		"stack": string(stack),
		"panic": fmt.Sprintf("%+v", v),
	})
}
