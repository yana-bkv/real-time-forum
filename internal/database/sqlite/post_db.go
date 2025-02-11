package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitPostDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./storage.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS posts (
    	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    	title Text,
    	body Text,
    	author Text
	);`

	_, err = DB.Exec(sqlStmt)
	if err != nil {
		log.Fatal("Error creating table", err, sqlStmt)
	}

}
