package handlers

import (
	"github.com/DirtyCajunRice/PeUD/internal/config"
	"github.com/sirupsen/logrus"
	"net/http"
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
