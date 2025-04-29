package services

import (
	"jwt-authentication/models"
	"net/http"
)

type AuthService interface {
	Register(data map[string]string) error
	Login(data map[string]string) (*http.Cookie, error)
	Logout() *http.Cookie
	GetAuthUser(cookie *http.Cookie) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}

type PostService interface {
	Create(authorId int, data map[string]string) (*models.Post, error)
	Get(postId string) (*models.Post, error)
	GetAll() ([]models.Post, error)
	Delete(postId string) error
}

type CommentService interface {
	Create(authorId int, postId int, data map[string]string) (*models.Comment, error)
	GetCommentsByPost(postId int) ([]models.Comment, error)
	Delete(postId int) error
}

type LikeService interface{}

type MessageService interface{}
