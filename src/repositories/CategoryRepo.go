package repositories

import (
	"errors"
	"jwt-authentication/database"
	"jwt-authentication/models"
)

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) Create(category *models.Category) error {
	sqlStmt := database.SqlCategoryDb("createCategory")
	if category.Text == "" {
		return errors.New("category text cannot be empty")
	}

	_, err := database.DB.Exec(sqlStmt, category.Text)
	if err != nil {
		return err
	}
	//// Adds id from db to json data
	//categoryID, err := result.LastInsertId()
	//if err != nil {
	//	return err
	//}
	//category.Id = int(categoryID)

	return nil
}

func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	return nil, nil
}

func (r *categoryRepository) Delete(id string) error {
	return nil
}
