package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"jwt-authentication/database"
	"jwt-authentication/models"
	"strconv"
)

func CreatePost(db *sql.DB, post *models.Post) error {
	sqlStmt := database.SqlPostDb("createPost")

	if post.Title == "" && post.Body == "" {
		return errors.New("title and body are required")
	}

	result, err := db.Exec(sqlStmt, post.Title, post.Body, post.Category, post.AuthorId, post.Time)
	if err != nil {
		fmt.Println(err)
		return errors.New("error creating post")
	}

	// Adds id from db to json data
	postID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	post.Id = int(postID)

	return nil

}

func GetPostById(db *sql.DB, id string) (*models.Post, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid post ID")
	}

	sqlStmt := database.SqlPostDb("getPostById")
	row := db.QueryRow(sqlStmt, intID)

	var post models.Post
	err = row.Scan(&post.Id, &post.Title, &post.Body, &post.AuthorId, &post.Time)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func GetAllPosts(db *sql.DB) ([]models.Post, error) {
	sqlStmt := database.SqlPostDb("getAllPosts")
	rows, err := db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.Id, &post.Title, &post.Body, &post.AuthorId, &post.Time)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func DeletePost(db *sql.DB, id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid post ID")
	}

	sqlStmt := database.SqlPostDb("deletePost")

	result, err := db.Exec(sqlStmt, intID)
	if err != nil {
		return errors.New("error deleting post")
	}

	// Check if any row was deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("post not found") // No post deleted
	}

	return nil
}
