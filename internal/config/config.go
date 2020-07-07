package config

import (
	"os"

	"github.com/DirtyCajunRice/PeUD/internal/database"
)

type Config struct {
	APIServer      *APIServer
	Debug          bool
	JSONLogging    bool
	Database       *database.Database
	Authentication *Authentication
}

type APIServer struct {
	Address string
	Port    int
}

type Authentication struct {
	PlexToken     string
	TautulliKey   string
	TautulliURL   string
	OrganizrToken string
	OrganizrURL   string
	OmbiURL       string
	OmbiKey       string
}

func (c *Config) LoadFromEnv() {
	c.Authentication = &Authentication{
		PlexToken:     os.Getenv("PLEX_TOKEN"),
		TautulliURL:   os.Getenv("TAUTULLI_URL"),
		TautulliKey:   os.Getenv("TAUTULLI_API_KEY"),
		OrganizrToken: os.Getenv("ORGANIZR_TOKEN"),
		OrganizrURL:   os.Getenv("ORGANIZR_URL"),
		OmbiURL:       os.Getenv("OMBI_URL"),
		OmbiKey:       os.Getenv("OMBI_API_KEY"),
	}
}
