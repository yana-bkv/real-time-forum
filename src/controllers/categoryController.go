package controllers

import (
	"github.com/gorilla/mux"
	"jwt-authentication/models"
	"jwt-authentication/repositories"
	"net/http"
	"strconv"
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
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	category, err := c.categoryRepository.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = EncodeJson(w, category)
	if err != nil {
		return
	}
}

func (c *CategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	categories, err := c.categoryRepository.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = EncodeJson(w, categories)
	if err != nil {
		return
	}
}

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = c.categoryRepository.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = EncodeJson(w, c)
	if err != nil {
		return
	}
}
