package config

import "github.com/DirtyCajunRice/PeUD/internal/database"

type Config struct {
	APIServer      *APIServer
	Debug          bool
	Database       *database.Database
	Authentication *Authentication
}

type APIServer struct {
	Address string
	Port    int
}

type Authentication struct {
	PlexToken   string
	TautulliKey string
}
