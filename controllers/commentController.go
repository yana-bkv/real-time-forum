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

func CreateComment(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	// ID of a post where user adds comment
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request payload for posts id (comment)", http.StatusBadRequest)
		return
	}

	// Decode JSON request body
	err = DecodeJson(r, w, &data)
	if err != nil {
		return
	}

	timeNow := time.Now()
	authorId, err := strconv.Atoi(GetUserId(w, r)) // Get authorized user who creates comment
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	comment := models.Comment{
		Body:     data["body"],
		PostId:   id,
		AuthorId: authorId,
		Time:     timeNow.Format("2006-01-02 15:04:05"),
	}

	//database is package, CreateUser is function, DB is *sql.DB, &user is *models.User
	err = repositories.CreateComment(database.DB, &comment)
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

func GetCommentsByPostId(w http.ResponseWriter, r *http.Request) {
	// ID of a post where user adds comment
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid request payload for posts id", http.StatusBadRequest)
		return
	}

	// Fetch post from database
	comments, err := repositories.GetCommentsByPost(database.DB, id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Comment not found", http.StatusNotFound)
		} else {
			fmt.Println(err, comments)
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, comments)
	if err != nil {
		return
	}
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	// Get ID from query parameters
	vars := mux.Vars(r)
	idStr := vars["cid"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println(err, idStr)
		http.Error(w, "Invalid request payload for comment id", http.StatusBadRequest)
		return
	}

	// Fetch post from database
	err = repositories.DeleteComment(database.DB, id)
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
	err = EncodeJson(w, "comment deleted")
	if err != nil {
		return
	}
}
