package controllers

import (
	"jwt-authentication/repositories"
	"net/http"
)

type PostCategoryController struct {
	postRepository repositories.CategoryRepository
}

func NewPostCategoryController(postRepository repositories.CategoryRepository) *PostCategoryController {
	return &PostCategoryController{postRepository: postRepository}
}

func (c *PostCategoryController) Create(w http.ResponseWriter, r *http.Request) {

}
