package handlers

import (
	"bytes"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/DirtyCajunRice/PeUD/internal/config"
)

type Env struct {
	Log    *logrus.Logger
	Config *config.Config
}

type Handler struct {
	*Env
	Handle func(e *Env, w http.ResponseWriter, r *http.Request)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handle(h.Env, w, r)
}

func arrayOrObject(data []byte) (isArray, isObject bool) {
	x := bytes.TrimLeft(data, " \t\r\n")
	isArray = len(x) > 0 && x[0] == '['
	isObject = len(x) > 0 && x[0] == '{'
	return
}
