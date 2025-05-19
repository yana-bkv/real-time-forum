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

	return nil
}

func (r *categoryRepository) GetById(id int) (*models.Category, error) {
	sqlStmt := database.SqlCategoryDb("getCategoryById")
	row := database.DB.QueryRow(sqlStmt, id)

	var category models.Category
	err := row.Scan(&category.Id, &category.Text)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	sqlStmt := database.SqlCategoryDb("getAllCategories")
	rows, err := database.DB.Query(sqlStmt)
	if err != nil {
		return nil, err
	}

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Id, &category.Text)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (r *categoryRepository) Delete(id int) error {
	sqlStmt := database.SqlCategoryDb("deleteCategory")
	_, err := database.DB.Exec(sqlStmt, id)
	if err != nil {
		return err
	}
	return nil
}
