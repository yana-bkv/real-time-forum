package repositories

import (
	"errors"
	"fmt"
	"jwt-authentication/database"
	"jwt-authentication/models"
)

type commentRepository struct{}

func NewCommentRepository() CommentRepository {
	return &commentRepository{}
}

func (c *commentRepository) Create(comment *models.Comment) error {
	sqlStmt := database.SqlCommentDb("createComment")

	if comment.Body == "" {
		return errors.New("body is required")
	}

	result, err := database.DB.Exec(sqlStmt, comment.Body, comment.PostId, comment.AuthorId, comment.Time)
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

func (c *commentRepository) GetCommentsByPost(id int) (*[]models.Comment, error) {
	sqlStmt := database.SqlCommentDb("getCommentsByPostId")
	rows, err := database.DB.Query(sqlStmt, id)
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

func (c *commentRepository) Delete(id int) error {
	sqlStmt := database.SqlCommentDb("deleteComment")

	result, err := database.DB.Exec(sqlStmt, id)
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
