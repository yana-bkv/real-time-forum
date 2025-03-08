package database

import (
	"database/sql"
	"errors"
	"fmt"
	"jwt-authentication/models"
)

func CreateComment(db *sql.DB, comment *models.Comment) error {
	sqlStmt := SqlCommentDb("createComment")

	if comment.Body == "" {
		return errors.New("body is required")
	}

	result, err := db.Exec(sqlStmt, comment.Body, comment.PostId, comment.AuthorId, comment.Time)
	if err != nil {
		fmt.Println(err)
		return errors.New("error creating comment")
	}

	// Adds id from db to json data
	commentID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	comment.Id = int(commentID)

	return nil
}

func GetCommentsByPost(db *sql.DB, id int) (*[]models.Comment, error) {
	sqlStmt := SqlCommentDb("getCommentsByPostId")
	rows, err := db.Query(sqlStmt, id)
	if err != nil {
		return nil, err
	}

	var comments []models.Comment
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.Id, &comment.Body, &comment.PostId, &comment.AuthorId, &comment.Time)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return &comments, nil
}

func DeleteComment(db *sql.DB, id int) error {
	sqlStmt := SqlCommentDb("deleteComment")

	result, err := db.Exec(sqlStmt, id)
	if err != nil {
		return errors.New("error deleting comment")
	}

	// Check if any row was deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("comment not found") // No post deleted
	}

	return nil
}
