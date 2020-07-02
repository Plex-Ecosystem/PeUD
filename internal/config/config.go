package config

import "github.com/DirtyCajunRice/PeUD/internal/database"

type Config struct {
	APIServer *APIServer
	Debug     bool
	Database  *database.Database
}

type APIServer struct {
	Address string
	Port    int
}
