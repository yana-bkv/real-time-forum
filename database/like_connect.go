package database

import (
	"database/sql"
	"errors"
	"fmt"
	"jwt-authentication/models"
)

func AddLike(db *sql.DB, like *models.Like) error {
	sqlStmt := SqlLikeDb("addLike")

	result, err := db.Exec(sqlStmt, like.PostId, like.AuthorId)
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

func GetLikes(db *sql.DB, likeByPostId int) (*[]models.Like, error) {
	sqlStmt := SqlLikeDb("getLikesByPostId")
	rows, err := db.Query(sqlStmt, likeByPostId)
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

func DeleteLike(db *sql.DB, likeByPostId int) error {
	sqlStmt := SqlLikeDb("deleteLike")

	result, err := db.Exec(sqlStmt, likeByPostId)
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
