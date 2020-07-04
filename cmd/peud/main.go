package main

import (
	"github.com/jmoiron/modl"

	"github.com/DirtyCajunRice/PeUD/internal/config"
	"github.com/DirtyCajunRice/PeUD/internal/database"
	"github.com/DirtyCajunRice/PeUD/internal/handlers"
	"github.com/DirtyCajunRice/PeUD/internal/server"
	"github.com/DirtyCajunRice/PeUD/logger"
)

var (
	log     = logger.New()
	version = "0.0.0"
	date    string
)

func main() {
	env := handlers.Env{
		Log: log,
		Config: &config.Config{
			APIServer: &config.APIServer{},
			Database: &database.Database{
				Log:   log.WithField("package", "database"),
				DbMap: &modl.DbMap{},
			},
		},
	}

	server.CLI(&version, &date, &env)
}
