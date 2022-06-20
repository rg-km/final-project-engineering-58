package main

import (
	"backend/internal/config"
	"database/sql"
	"fmt"
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
		password 	string,
		created_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	
	_, err := db.Exec(userTable)
	if err != nil {
		panic(err)
		return
	}

	postTable := `CREATE TABLE IF NOT EXISTS posts (
		id			string,
		title		string,
		description	string,
		content		string,
		url_video	string,
		category_id	string,
		user_id		string,
		parent_id	string,
		created_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(postTable)
	if err != nil {
		panic(err)
		return
	}

	categoryTable := `CREATE TABLE IF NOT EXISTS categories (
		id			string,
		name		string,
		created_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at	TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(categoryTable)
	if err != nil {
		panic(err)
		return
	}

	fmt.Println("Successfully migrated")
}