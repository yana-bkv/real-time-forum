package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"jwt-authentication/database"
	"jwt-authentication/models"
	"net/http"
	"strconv"
)

func AddLike(w http.ResponseWriter, r *http.Request) {
	// ID of a post where user adds comment
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request payload for posts id", http.StatusBadRequest)
		return
	}

	authorId, err := strconv.Atoi(GetUserId(w, r)) // Get authenticated user that puts like
	if err != nil {
		fmt.Println(err, authorId)
		http.Error(w, "Invalid request payload for like", http.StatusBadRequest)
		return
	}

	like := models.Like{
		PostId:   id,
		AuthorId: authorId,
	}

	//database is package, CreateUser is function, DB is *sql.DB, &user is *models.User
	err = database.AddLike(database.DB, &like)
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

func GetLikes(w http.ResponseWriter, r *http.Request) {
	// ID of a post where user adds comment
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request payload for like id", http.StatusBadRequest)
		return
	}

	// Fetch post from database
	likes, err := database.GetLikes(database.DB, id)
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

func DeleteLike(w http.ResponseWriter, r *http.Request) {
	// Get ID from query parameters
	vars := mux.Vars(r)
	idStr := vars["lId"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request payload for like id", http.StatusBadRequest)
		return
	}

	// Fetch post from database
	err = database.DeleteLike(database.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Comment not found", http.StatusNotFound)
		} else {
			fmt.Println(err, idStr)
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
