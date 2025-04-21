package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"jwt-authentication/database"
	"jwt-authentication/models"
	"jwt-authentication/repositories"
	"net/http"
	"strconv"
	"time"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	// Decode JSON request body
	err := DecodeJson(r, w, &data)
	if err != nil {
		return
	}

	timeNow := time.Now()
	authorId, err := strconv.Atoi(GetUserId(w, r))
	if err != nil {
		http.Error(w, "Invalid request payload for post", http.StatusBadRequest)
		return
	}

	post := models.Post{
		Title:    data["title"],
		Body:     data["body"],
		Category: "",
		AuthorId: authorId,
		Time:     timeNow.Format("2006-01-02 15:04:05"),
	}

	//database is package, CreateUser is function, DB is *sql.DB, &user is *models.User
	err = repositories.CreatePost(database.DB, &post)
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

func GetPost(w http.ResponseWriter, r *http.Request) {
	// Get ID from query parameters
	vars := mux.Vars(r)
	idStr := vars["id"]

	// Fetch post from database
	post, err := repositories.GetPostById(database.DB, idStr)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			fmt.Println(err, idStr)
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, post)
	if err != nil {
		return
	}
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	// Fetch post from database
	posts, err := repositories.GetAllPosts(database.DB)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			fmt.Println(err, posts)
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, posts)
	if err != nil {
		return
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	// Get ID from query parameters
	vars := mux.Vars(r)
	idStr := vars["id"]
	fmt.Println(idStr)

	// Fetch post from database
	err := repositories.DeletePost(database.DB, idStr)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			fmt.Println(err, idStr)
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, "post deleted")
	if err != nil {
		return
	}
}
