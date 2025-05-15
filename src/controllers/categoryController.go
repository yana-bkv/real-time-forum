package controllers

import (
	"jwt-authentication/models"
	"jwt-authentication/repositories"
	"net/http"
)

type CategoryController struct {
	categoryRepository repositories.CategoryRepository
}

func NewCategoryController(categoryRepository repositories.CategoryRepository) *CategoryController {
	return &CategoryController{categoryRepository: categoryRepository}
}

func (c *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	err := DecodeJson(r, w, &data)
	if err != nil {
		return
	}

	category := models.Category{
		Text: data["text"],
	}

	err = c.categoryRepository.Create(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, category)
	if err != nil {
		return
	}
}

func (c *CategoryController) Get(w http.ResponseWriter, r *http.Request) {
}

func (c *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
}

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
}
