package repositories

import (
	"errors"
	"fmt"
	"jwt-authentication/database"
	"jwt-authentication/models"
	"strconv"
)

type postRepository struct{}

func NewPostRepository() PostRepository {
	return &postRepository{}
}

func (c *postRepository) Create(post *models.Post) (int, error) {
	sqlStmt := database.SqlPostDb("createPost")

	if post.Title == "" && post.Body == "" {
		return 0, errors.New("title and body are required")
	}

	result, err := database.DB.Exec(sqlStmt, post.Title, post.Body, post.AuthorId, post.Time)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("error creating post")
	}

	// Adds id from db to json data
	postID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	post.Id = int(postID)

	return post.Id, nil
}

func (c *postRepository) GetPostById(id string) (*models.Post, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("invalid post ID")
	}

	sqlStmt := database.SqlPostDb("getPostById")
	row := database.DB.QueryRow(sqlStmt, intID)

	var post models.Post
	err = row.Scan(&post.Id, &post.Title, &post.Body, &post.AuthorId, &post.Time)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (c *postRepository) GetAllPosts() ([]models.Post, error) {
	sqlStmt := database.SqlPostDb("getAllPosts")
	rows, err := database.DB.Query(sqlStmt)
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

func (c *postRepository) Delete(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return errors.New("invalid post ID")
	}

	sqlStmt := database.SqlPostDb("deletePost")

	result, err := database.DB.Exec(sqlStmt, intID)
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

// Assigning categories to post
func (c *postRepository) Assign(postCategory *models.PostCategory) error {
	sqlStmt := database.SqlPostCategoryDb("assignCategoryToPost")

	stmt, err := database.DB.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, catID := range postCategory.CategoryIDs {
		_, err := stmt.Exec(postCategory.PostID, catID)
		if err != nil {
			return err
		}
	}
	return nil
}
