package config

import (
	"os"
)

type Config struct {
	HttpPort string
}

func InitConfig() *Config {
	return &Config{
		HttpPort: os.Getenv("HTTP_PORT"),
	}
}
