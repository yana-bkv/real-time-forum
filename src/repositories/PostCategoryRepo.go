package repositories

import (
	"jwt-authentication/database"
	"jwt-authentication/models"
)

type postCategoryRepository struct{}

func NewPostCategoryRepository() PostCategoryRepository {
	return &postCategoryRepository{}
}

func (r *postCategoryRepository) Assign(postCategory *models.PostCategory) error {
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

func (r *postCategoryRepository) Remove(postId, categoryId int) error {
	return nil
}
