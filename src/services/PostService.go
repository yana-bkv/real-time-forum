package services

import (
	"database/sql"
	"errors"
	"jwt-authentication/models"
	"jwt-authentication/repositories"
	"time"
)

type PostServiceImpl struct {
	postRepo repositories.PostRepository
}

func NewPostService(postRepo repositories.PostRepository) *PostServiceImpl {
	return &PostServiceImpl{postRepo: postRepo}
}

func (p *PostServiceImpl) Create(authorId int, data map[string]string) (*models.Post, error) {
	timeNow := time.Now()

	post := models.Post{
		Title:    data["title"],
		Body:     data["body"],
		Category: "",
		AuthorId: authorId,
		Time:     timeNow.Format("2006-01-02 15:04:05"),
	}

	err := p.postRepo.Create(&post)
	if err != nil {
		return nil, errors.New("Unable to create post")
	}
	return &post, nil
}

func (p *PostServiceImpl) Get(postId string) (*models.Post, error) {
	post, err := p.postRepo.GetPostById(postId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Post not found")
		} else {
			return nil, errors.New("Database error")
		}
		return nil, err
	}
	return post, nil
}

func (p *PostServiceImpl) GetAll() ([]models.Post, error) {
	posts, err := p.postRepo.GetAllPosts()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Database error")
		}
		return nil, err
	}
	return posts, nil
}

func (p *PostServiceImpl) Delete(postId string) error {
	err := p.postRepo.Delete(postId)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Post not found")
		} else {
			return errors.New("Database error")
		}
		return err
	}
	return nil
}
