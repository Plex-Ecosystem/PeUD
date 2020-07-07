package main

import (
	"github.com/jmoiron/modl"

	v1 "github.com/DirtyCajunRice/PeUD/api/v1"
	"github.com/DirtyCajunRice/PeUD/internal/config"
	"github.com/DirtyCajunRice/PeUD/internal/database"
	"github.com/DirtyCajunRice/PeUD/internal/handlers"
	"github.com/DirtyCajunRice/PeUD/internal/server"
	"github.com/DirtyCajunRice/PeUD/logger"
)

var (
	log     = logger.New()
	version = "0.0.0"
	date    = "2020-01-01T00:00:00Z"
)

func main() {
	env := handlers.Env{
		Log:           log.WithField("package", "handlers"),
		MiddlewareLog: log.WithField("package", "peud_middleware"),
		Build: &v1.Version{
			Date:    &version,
			Version: &date,
		},
		Config: &config.Config{
			APIServer: &config.APIServer{},
			Database: &database.Database{
				Log:   log.WithField("package", "database"),
				DbMap: &modl.DbMap{},
			},
		},
	}

	server.CLI(&env)
}
