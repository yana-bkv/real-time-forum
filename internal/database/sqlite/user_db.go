package sqlite

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitUserDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./storage.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS users (
    	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    	username Text,
    	email Text,
    	password Text
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatal("Error creating table", err, sqlStmt)
	}

}
