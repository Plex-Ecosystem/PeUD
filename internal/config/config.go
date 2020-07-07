package config

import (
	"os"

	"github.com/DirtyCajunRice/PeUD/internal/database"
)

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

func (c *Config) LoadFromEnv() {
	c.Authentication = &Authentication{
		PlexToken:   os.Getenv("PLEX_TOKEN"),
		TautulliKey: os.Getenv("TAUTULLI_API_KEY"),
	}
}
