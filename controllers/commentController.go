package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"jwt-authentication/helpers"
	"jwt-authentication/services"
	"net/http"
	"strconv"
)

type CommentController struct {
	commentService services.CommentService
}

func NewCommentController(commentService services.CommentService) *CommentController {
	return &CommentController{commentService: commentService}
}

func (c *CommentController) Create(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	// ID of a post where user adds comment
	vars := mux.Vars(r)
	idStr := vars["id"]
	postId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request payload for posts id (comment)", http.StatusBadRequest)
		return
	}

	authorId, err := strconv.Atoi(helpers.GetUserId(w, r)) // Get authorized user who creates comment
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Decode JSON request body
	err = DecodeJson(r, w, &data)
	if err != nil {
		return
	}

	comment, err := c.commentService.Create(authorId, postId, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the received 'data' back as JSON response
	err = EncodeJson(w, comment)
	if err != nil {
		return
	}
}

func (c *CommentController) GetCommentsByPostId(w http.ResponseWriter, r *http.Request) {
	// ID of a post where user adds comment
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request payload for posts id", http.StatusBadRequest)
		return
	}

	comments, err := c.commentService.GetCommentsByPost(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Comment not found", http.StatusNotFound)
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, comments)
	if err != nil {
		return
	}
}

func (c *CommentController) Delete(w http.ResponseWriter, r *http.Request) {
	// Get ID from query parameters
	vars := mux.Vars(r)
	idStr := vars["cid"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err, idStr)
		http.Error(w, "Invalid request payload for comment id", http.StatusBadRequest)
		return
	}

	err = c.commentService.Delete(id)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, "comment deleted")
	if err != nil {
		return
	}
}
