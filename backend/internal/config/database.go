package config

import (
	"backend/pkg/config"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase() *sql.DB {
	cfg := config.DatabaseSqlite()

	db, err := sql.Open("sqlite3", cfg.Path)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
