package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"jwt-authentication/helpers"
	"jwt-authentication/models"
	"jwt-authentication/services"
	"net/http"
	"strconv"
)

type PostController struct {
	postService services.PostService
}

func NewPostController(postService services.PostService) *PostController {
	return &PostController{postService: postService}
}

func (c *PostController) Create(w http.ResponseWriter, r *http.Request) {
	var req models.CreatePostRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	authorId, err := strconv.Atoi(helpers.GetUserId(w, r))
	if err != nil {
		http.Error(w, "Invalid request payload for post", http.StatusBadRequest)
	}

	post, err := c.postService.Create(authorId, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = EncodeJson(w, post)
	if err != nil {
		return
	}
}

func (c *PostController) GetPost(w http.ResponseWriter, r *http.Request) {
	// Get ID from query parameters
	vars := mux.Vars(r)
	idStr := vars["id"]

	post, err := c.postService.Get(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, post)
	if err != nil {
		return
	}
}

func (c *PostController) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := c.postService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Encode response as JSON
	err = EncodeJson(w, posts)
	if err != nil {
		return
	}
}

func (c *PostController) Delete(w http.ResponseWriter, r *http.Request) {
	// Get ID from query parameters
	vars := mux.Vars(r)
	idStr := vars["id"]

	err := c.postService.Delete(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, "post deleted")
	if err != nil {
		return
	}
}
