package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	fmt.Println(CreateTableUser(db))
}

func CreateTableUser(db *sql.DB) (string, error) {
	userTable := `CREATE TABLE IF NOT EXISTS users (
		id			string,
		name		string,
		email		string,
		role		string,
		password 	string,
		created_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := db.Exec(userTable)
	if err != nil {
		return "Error", err
	}

	return "Successfully migrated", nil
}
