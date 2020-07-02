package config

type Config struct {
	APIServer *APIServer
	Debug     bool
}

type APIServer struct {
	Address string
	Port    int
}

func CreateConfig() *Config {
	return &Config{
		APIServer: &APIServer{},
	}
}
