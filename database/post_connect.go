package database

import (
	"database/sql"
	"errors"
	"jwt-authentication/models"
)

func CreatePost(db *sql.DB, post *models.Post) error {
	sqlStmt := SqlPostDb("createPost")

	if post.Title == "" && post.Body == "" {
		return errors.New("title and body are required")
	}

	result, err := db.Exec(sqlStmt, post.Title, post.Body, post.Time, post.AuthorName)
	if err != nil {
		return errors.New("error creating post")
	}

	// Adds id from db to json data
	postID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	post.Id = uint(postID)

	return nil

}
