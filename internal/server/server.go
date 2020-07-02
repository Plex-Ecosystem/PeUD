// Package server PeUD API
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /api/v1
//     Version: 0.1.0
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Nicholas St. Germain <nick@cajun.pro> https://cajun.pro
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
//     SecurityDefinitions:
//     api_key:
//          type: apiKey
//          name: API_KEY
//          in: header
//
// swagger:meta
package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/DirtyCajunRice/PeUD/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func fileServer(r chi.Router, path string) {
	workDir, _ := os.Getwd()
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler("/api/v1/"+path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(workDir, "web", "static", "redoc.html"))
	})
}

func Start(version, date *string, handlerEnv *handlers.Env) {
	log := handlerEnv.Log
	config := handlerEnv.Config
	router := chi.NewRouter()

	log.Info("Initializing Server")

	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(15 * time.Second))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	versionHandler := handlers.NewVersionHandler(version, date)
	router.Route("/api/v1", func(r chi.Router) {
		r.Method(http.MethodGet, "/version", handlers.Handler{Env: handlerEnv, Handle: versionHandler.GetVersion})
		fileServer(r, "/doc")
	})

	httpAddr := fmt.Sprintf("%s:%d", config.APIServer.Address, config.APIServer.Port)
	log.Infoln("API server is now listening on", httpAddr)
	log.Error(http.ListenAndServe(httpAddr, router))
}
