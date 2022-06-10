package main

import (
	"backend/internal/config"
	"database/sql"
)

var (
	db = config.InitDatabase()
)

func main() {
	defer func(db *sql.DB) { _ = db.Close() }(db)

	userTable := `CREATE TABLE IF NOT EXISTS users (
		id			string,
		name		string,
		email		string,
		role		string,
		password	string,
		created_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	
	_, err := db.Exec(userTable)
	if err != nil {
		panic(err)
		return
	}
}