package controllers

import (
	"jwt-authentication/repositories"
	"net/http"
)

type PostCategoryController struct {
	postCategoryRepository repositories.PostCategoryRepository
}

func NewPostCategoryController(postCategoryRepository repositories.PostCategoryRepository) *PostCategoryController {
	return &PostCategoryController{postCategoryRepository: postCategoryRepository}
}

func (c *PostCategoryController) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *PostCategoryController) Get(w http.ResponseWriter, r *http.Request) {
}
func (c *PostCategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
}
func (c *PostCategoryController) Delete(w http.ResponseWriter, r *http.Request) {
}
