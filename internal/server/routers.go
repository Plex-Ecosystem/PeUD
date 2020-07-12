package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/DirtyCajunRice/PeUD/internal/handlers"
	peudMiddleware "github.com/DirtyCajunRice/PeUD/internal/middleware"
)

func Start(Env *handlers.Env) {
	log := Env.Log
	config := Env.Config
	router := chi.NewRouter()
	log.Info("Initializing Server")
	router.Use(
		middleware.Recoverer,
		middleware.Timeout(15*time.Second),
		render.SetContentType(render.ContentTypeJSON),
		middleware.RequestID,
		middleware.RealIP,
		peudMiddleware.NewMiddlewareLogger(Env.MiddlewareLog),
		middleware.Compress(5),
		middleware.RedirectSlashes,
	)
	router.Route("/", func(r chi.Router) {
		r.Method(http.MethodGet, "/version", handlers.Handler{Env: Env, Handle: handlers.Version})
		r.Route("/users", func(r chi.Router) {
			r.Method(http.MethodGet, "/{plex,tautulli,organizr,ombi}", handlers.Handler{Env: Env, Handle: handlers.ListUsers})
			r.Method(http.MethodGet, "/{plex,tautulli,organizr,ombi}/{id:[0-9]+}", handlers.Handler{Env: Env, Handle: handlers.GetUser})
		})
		r.Route("/sync", func(r chi.Router) {
			r.Method(http.MethodPut, "/", handlers.Handler{Env: Env, Handle: handlers.Sync})
		})
	})

	config.Database.Init()
	httpAddr := fmt.Sprintf("%s:%d", config.APIServer.Address, config.APIServer.Port)
	log.Infoln("API server is now listening on", httpAddr)
	log.Error(http.ListenAndServe(httpAddr, router))
}
