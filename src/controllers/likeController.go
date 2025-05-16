package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"jwt-authentication/helpers"
	"jwt-authentication/models"
	"jwt-authentication/repositories"
	"net/http"
	"strconv"
)

type LikeController struct {
	likeRepo repositories.LikeRepository
}

func NewLikeController(likeRepo repositories.LikeRepository) *LikeController {
	return &LikeController{likeRepo: likeRepo}
}

func (c *LikeController) AddLikeToPost(w http.ResponseWriter, r *http.Request) {
	// ID of a post where user adds comment
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request payload for posts id", http.StatusBadRequest)
		return
	}

	authorId, err := strconv.Atoi(helpers.GetUserId(w, r))
	if err != nil {
		http.Error(w, "Invalid request payload for post", http.StatusBadRequest)
		return
	}

	like := models.Like{
		PostId:   id,
		AuthorId: authorId,
	}

	err = c.likeRepo.CreateLike(&like)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, like)
	if err != nil {
		return
	}
}

func (c *LikeController) GetLikesByPostId(w http.ResponseWriter, r *http.Request) {
	// ID of a post where user adds comment
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request payload for like id", http.StatusBadRequest)
		return
	}

	// Fetch post from database
	likes, err := c.likeRepo.GetLikesByPostId(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Like not found", http.StatusNotFound)
		} else {
			fmt.Println(err, likes)
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, likes)
	if err != nil {
		return
	}
}

func (c *LikeController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postID := vars["id"]

	// Optional: convert to int
	id, err := strconv.Atoi(postID)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	userID := helpers.GetUserId(w, r)
	userId, _ := strconv.Atoi(userID)

	// Fetch post from database
	err = c.likeRepo.Delete(userId, id)
	if err != nil {
		if errors.Is(err, errors.New("like not found")) {
			http.Error(w, "Like not found", http.StatusNotFound)
		} else {
			fmt.Println(err, userId, id)
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, "like deleted")
	if err != nil {
		return
	}
}

func (c *LikeController) HasLiked(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	postIDStr := vars["id"]
	// Преобразование строки в число
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		http.Error(w, `Invalid post ID ${postID}`, http.StatusBadRequest)
		return
	}

	userID := helpers.GetUserId(w, r)
	userId, _ := strconv.Atoi(userID)

	count, err := c.likeRepo.HasLiked(userId, postID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Like not found", http.StatusNotFound)
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]bool{"hasLiked": count > 0}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

}

func (c *LikeController) GetUserCount(w http.ResponseWriter, r *http.Request) {
	// ID of a post where user adds comment
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request payload for like id", http.StatusBadRequest)
		return
	}

	// Fetch post from database
	count, err := c.likeRepo.GetUserCount(id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Like not found", http.StatusNotFound)
		} else {
			fmt.Println(err, count)
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, count)
	if err != nil {
		return
	}
}
