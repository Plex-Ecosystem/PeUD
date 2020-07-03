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
// ---
// host: localhost
// basePath: /api/v1
// info:
//   version: 0.1.0
//   license:
//     name: MIT
//     url: http://opensource.org/licenses/MIT
//   contact:
//     name: Nicholas St. Germain
//     email: nick@cajun.pro
//     url: https://cajun.pro
//
// tags:
//   - name: General
//     description: Generic calls
//   - name: Users
//     description: Operations to the user database
//
// consumes:
//   - application/json
//
// produces:
//   - application/json
//
// x-tagGroups:
//  - name: General
//    tags:
//      - General
//  - name: Users
//    tags:
//      - Users
//  - name: Maintenance
//    tags:
//      - maintenance
// swagger:meta
package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/DirtyCajunRice/PeUD/internal/handlers"
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

func Start(version, date *string, Env *handlers.Env) {
	log := Env.Log
	config := Env.Config
	router := chi.NewRouter()

	log.Info("Initializing Server")

	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(15 * time.Second))
	router.Use(render.SetContentType(render.ContentTypeJSON))

	versionHandler := handlers.NewVersionHandler(version, date)
	router.Route("/api/v1", func(r chi.Router) {
		r.Method(http.MethodGet, "/version", handlers.Handler{Env: Env, Handle: versionHandler.GetVersion})
		r.Method(http.MethodGet, "/users", handlers.Handler{Env: Env, Handle: handlers.ListUsers})
		r.Method(http.MethodPost, "/users", handlers.Handler{Env: Env, Handle: handlers.CreateUsers})
		fileServer(r, "/doc")
	})

	config.Database.Init()
	httpAddr := fmt.Sprintf("%s:%d", config.APIServer.Address, config.APIServer.Port)
	log.Infoln("API server is now listening on", httpAddr)
	log.Error(http.ListenAndServe(httpAddr, router))
}
