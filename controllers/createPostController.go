package controllers

import (
	"encoding/json"
	"jwt-authentication/database"
	"jwt-authentication/models"
	"net/http"
	"time"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	timeNow := time.Now()
	authorUsername := GetUsername(w, r)

	post := models.Post{
		Title:      data["title"],
		Body:       data["body"],
		Time:       timeNow.Format("2006-01-02 15:04:05"),
		AuthorName: authorUsername,
	}

	//database is package, CreateUser is function, DB is *sql.DB, &user is *models.User
	err := database.CreatePost(database.DB, &post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the received 'data' back as JSON response
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {}

func GetPosts(w http.ResponseWriter, r *http.Request) {}

func DeletePost(w http.ResponseWriter, r *http.Request) {}
