package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	v1 "github.com/DirtyCajunRice/PeUD/api/v1"
	"github.com/DirtyCajunRice/PeUD/internal/config"
)

type Env struct {
	Log           *logrus.Entry
	MiddlewareLog *logrus.Entry
	Config        *config.Config
	Build         *v1.Version
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

func (e *Env) toJSON(w http.ResponseWriter, r *http.Request, i interface{}, status ...int) {
	log := e.Log.WithField("function", "toJSON")
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(true)
	if err := encoder.Encode(i); err != nil {
		log.Error(err)
		http.Error(w, "Could not marshal response to JSON", http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	if len(status) > 0 && status[0] != 0 {
		w.WriteHeader(status[0])
	} else {
		w.WriteHeader(http.StatusOK)
	}

	log.Tracef("response JSON: %s", buffer.Bytes())
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Error(err)
		http.Error(w, "Could not write http response", http.StatusInternalServerError)
	}
}
