package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"jwt-authentication/database"
	"jwt-authentication/models"
)

type likeRepository struct{}

func NewLikeRepository() LikeRepository {
	return &likeRepository{}
}

func (l *likeRepository) CreateLike(like *models.Like) error {
	sqlStmt := database.SqlLikeDb("addLike")

	result, err := database.DB.Exec(sqlStmt, like.PostId, like.AuthorId)
	if err != nil {
		fmt.Println(err)
		return errors.New("error adding like")
	}

	// Adds id from db to json data
	likeID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	like.Id = int(likeID)

	return nil
}

func (l *likeRepository) GetLikesByPostId(likeByPostId int) (*[]models.Like, error) {
	sqlStmt := database.SqlLikeDb("getLikesByPostId")
	rows, err := database.DB.Query(sqlStmt, likeByPostId)
	if err != nil {
		return nil, err
	}

	var likes []models.Like
	for rows.Next() {
		var like models.Like
		err := rows.Scan(&like.Id, &like.PostId, &like.AuthorId)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	return &likes, nil
}

func (l *likeRepository) Delete(userId, postId int) error {
	sqlStmt := database.SqlLikeDb("deleteLike")

	result, err := database.DB.Exec(sqlStmt, userId, postId)
	if err != nil {
		fmt.Println(err)
		return errors.New("error deleting like")
	}

	// Check if any row was deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("like not found") // No post deleted
	}

	return nil
}

func (l *likeRepository) HasLiked(userID, likeByPostId int) (int, error) {
	var count int
	err := database.DB.QueryRow("SELECT COUNT(*) FROM likes WHERE user_id = ? AND post_id = ?", userID, likeByPostId).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (l *likeRepository) GetUserCount(postId int) (int, error) {
	sqlStmt := database.SqlLikeDb("getUserCount")
	var likeCount int

	err := database.DB.QueryRow(sqlStmt, postId).Scan(&likeCount)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil // Если лайков нет, просто возвращаем 0
		}
		return 0, fmt.Errorf("error getting like count: %w", err)
	}

	return likeCount, nil
}
