package services

import (
	"database/sql"
	"errors"
	"jwt-authentication/models"
	"jwt-authentication/repositories"
	"time"
)

type CommentServiceImpl struct {
	commentRepo repositories.CommentRepository
}

func NewCommentService(commentRepo repositories.CommentRepository) *CommentServiceImpl {
	return &CommentServiceImpl{commentRepo: commentRepo}
}

func (c *CommentServiceImpl) Create(authorId int, postId int, data map[string]string) (*models.Comment, error) {
	timeNow := time.Now()

	comment := models.Comment{
		Body:     data["body"],
		PostId:   postId,
		AuthorId: authorId,
		Time:     timeNow.Format("2006-01-02 15:04:05"),
	}

	err := c.commentRepo.Create(&comment)
	if err != nil {
		return nil, errors.New("Internal Server Error")
	}

	return &comment, nil
}

func (c *CommentServiceImpl) GetCommentsByPost(id int) (**[]models.Comment, error) {
	// Fetch post from database
	comments, err := c.commentRepo.GetCommentsByPost(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Comment not found")
		}
		return nil, errors.New("Database error")
	}
	return &comments, nil
}

func (c *CommentServiceImpl) Delete(postId int) error {
	err := c.commentRepo.Delete(postId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Comment not found")
		}
		return errors.New("Database error")
	}
	return nil
}
