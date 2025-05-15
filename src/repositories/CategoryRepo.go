package repositories

import "jwt-authentication/models"

type categoryRepository struct{}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) Create(category *models.Category) error {
	return nil
}

func (r *categoryRepository) GetAllCategories() ([]models.Category, error) {
	return nil, nil
}

func (r *categoryRepository) Delete(id string) error {
	return nil
}
