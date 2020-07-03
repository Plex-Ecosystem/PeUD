package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/DirtyCajunRice/PeUD/internal/handlers"
)

func fileServer(r chi.Router, path string) {
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "docs/api"))
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler("/api/v1/"+path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"
	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(filesDir))
		fs.ServeHTTP(w, r)
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
		r.Method(http.MethodGet, "/user", handlers.Handler{Env: Env, Handle: handlers.ListUser})
		r.Method(http.MethodPost, "/user", handlers.Handler{Env: Env, Handle: handlers.CreateUsers})
		fileServer(r, "/doc")
	})

	config.Database.Init()
	httpAddr := fmt.Sprintf("%s:%d", config.APIServer.Address, config.APIServer.Port)
	log.Infoln("API server is now listening on", httpAddr)
	log.Error(http.ListenAndServe(httpAddr, router))
}
