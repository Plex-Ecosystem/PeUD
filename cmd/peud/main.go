package main

import (
	"github.com/DirtyCajunRice/PeUD/internal/config"
	"github.com/DirtyCajunRice/PeUD/internal/handlers"
	"github.com/DirtyCajunRice/PeUD/internal/server"
	"github.com/DirtyCajunRice/PeUD/logger"
	_ "github.com/mattn/go-sqlite3"
)

var (
	log     = logger.New()
	version = "0.0.0"
	date    string
)

func main() {
	cfg := config.CreateConfig()
	handlerEnv := handlers.Env{
		Log:    log,
		Config: cfg,
	}
	server.CLI(&version, &date, &handlerEnv)
}
