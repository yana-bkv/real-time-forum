package repositories

import (
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

func (l *likeRepository) Delete(likeByPostId int) error {
	sqlStmt := database.SqlLikeDb("deleteLike")

	result, err := database.DB.Exec(sqlStmt, likeByPostId)
	if err != nil {
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
