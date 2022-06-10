package config

import (
	"os"
)

type DatabaseSqliteConfig struct {
	Path	string
}

func DatabaseSqlite() DatabaseSqliteConfig {
	return DatabaseSqliteConfig{
		Path: os.Getenv("DB_PATH"),
	}
}
