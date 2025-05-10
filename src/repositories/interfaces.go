package repositories

import "jwt-authentication/models"

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	Create(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	GetUserById(id string) (*models.User, error)
}

type PostRepository interface {
	Create(user *models.Post) error
	GetAllPosts() ([]models.Post, error)
	GetPostById(id string) (*models.Post, error)
	Delete(id string) error
}

type MessageRepository interface {
	Save(message *models.Message) error
	GetMsg(sender, receiver string) ([]models.Message, error)
}

type CommentRepository interface {
	Create(user *models.Comment) error
	GetCommentsByPost(id int) (*[]models.Comment, error)
	Delete(id int) error
}

type LikeRepository interface {
	CreateLike(like *models.Like) error
	GetLikesByPostId(id int) (*[]models.Like, error)
	Delete(postId, userId int) error
	HasLiked(userID, likeByPostId int) (int, error)
	GetUserCount(postId int) (int, error)
}
