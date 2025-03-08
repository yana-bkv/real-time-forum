package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func ConnectDB() {
	// Initialize db
	var err error
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	DB = db

	// Create user table
	sqlStmtUser := SqlUserDb("createTable")
	_, err = db.Exec(sqlStmtUser)
	if err != nil {
		log.Fatal("Error creating user table", err, sqlStmtUser)
	}

	// Create posts table
	sqlStmtPost := SqlPostDb("createTable")
	_, err = db.Exec(sqlStmtPost)
	if err != nil {
		log.Fatal("Error creating post table", err, sqlStmtPost)
	}

	// Create comments table
	sqlStmtComment := SqlCommentDb("createTable")
	_, err = db.Exec(sqlStmtComment)
	if err != nil {
		log.Fatal("Error creating comment table", err, sqlStmtComment)
	}

	// Create likes table
	sqlStmtLike := SqlLikeDb("createTable")
	_, err = db.Exec(sqlStmtLike)
	if err != nil {
		log.Fatal("Error creating like table", err, sqlStmtLike)
	}
}
