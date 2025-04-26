package repositories

import "jwt-authentication/models"

type UserRepository interface {
	Create(user *models.User) error
}

type PostRepository interface {
}

type MessageRepository interface {
}

type CommentRepository interface {
}

type LikeRepository interface {
}
