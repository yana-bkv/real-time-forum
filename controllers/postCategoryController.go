package controllers

import (
	"github.com/gorilla/mux"
	"jwt-authentication/models"
	"jwt-authentication/repositories"
	"net/http"
	"strconv"
)

type PostCategoryController struct {
	postCategoryRepository repositories.PostCategoryRepository
}

func NewPostCategoryController(postCategoryRepository repositories.PostCategoryRepository) *PostCategoryController {
	return &PostCategoryController{postCategoryRepository: postCategoryRepository}
}

func (c *PostCategoryController) Create(w http.ResponseWriter, r *http.Request) {
	var data map[string][]int

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Decode JSON request body
	err = DecodeJson(r, w, &data)
	if err != nil {
		return
	}

	postCategory := models.PostCategory{
		PostID:      id,
		CategoryIDs: data["category_ids"],
	}

	err = c.postCategoryRepository.Assign(&postCategory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Send the received 'data' back as JSON response
	err = EncodeJson(w, postCategory)
	if err != nil {
		return
	}
}

func (c *PostCategoryController) Get(w http.ResponseWriter, r *http.Request) {
}
func (c *PostCategoryController) GetAll(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tags, err := c.postCategoryRepository.GetTagByPostId(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, tags)
	if err != nil {
		return
	}
}

func (c *PostCategoryController) GetTagByPostId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	categoryIDs, err := c.postCategoryRepository.GetTagByPostId(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	categoryNames, err := c.postCategoryRepository.GetCategoryNamesByIDs(categoryIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Пример: вернуть просто массив названий
	names := []string{}
	for _, id := range categoryIDs {
		if text, ok := categoryNames[id]; ok {
			names = append(names, text)
		}
	}

	_ = EncodeJson(w, names)
}

func (c *PostCategoryController) Delete(w http.ResponseWriter, r *http.Request) {
}
