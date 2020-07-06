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
	peud_middlware "github.com/DirtyCajunRice/PeUD/internal/middleware"
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

	router.Use(
		// allow recovery from failure
		middleware.Recoverer,
		// set timeout for long running commands
		middleware.Timeout(15*time.Second),
		// force all render contentType to JSON
		render.SetContentType(render.ContentTypeJSON),
		// use requestIDs
		middleware.RequestID,
		// use real IPs
		middleware.RealIP,
		// enable middleware logger
		peud_middlware.NewMiddlewareLogger(Env.MiddlewareLog),
		// enable default compression
		middleware.Compress(5),
		// redirect slashes to non slashes for endpoints
		middleware.RedirectSlashes,
	)

	router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Method(http.MethodGet, "/version", handlers.Handler{Env: Env, Handle: handlers.Version})
			r.Route("/users", func(r chi.Router) {
				// r.Method(http.MethodGet, "/{id:[0-9]{1,12}", handlers.Handler{Env: Env, Handle:})
				r.Route("/plex", func(r chi.Router) {
					r.Method(http.MethodGet, "/", handlers.Handler{Env: Env, Handle: handlers.ListUser})
					r.Method(http.MethodPost, "/", handlers.Handler{Env: Env, Handle: handlers.CreateUsers})
				})
			})
			fileServer(r, "/doc")
		})
	})

	config.Database.Init()
	httpAddr := fmt.Sprintf("%s:%d", config.APIServer.Address, config.APIServer.Port)
	log.Infoln("API server is now listening on", httpAddr)
	log.Error(http.ListenAndServe(httpAddr, router))
}
