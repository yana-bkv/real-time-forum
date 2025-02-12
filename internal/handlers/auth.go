package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"log/slog"
	"net/http"
	"real-time-forum/internal/database/sqlite"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Get values from FORM html file when user submits
		username := r.FormValue("username")
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Add form values to database
		sqlite.CreateUser(username, email, password)

		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}

	RenderTemplate("./web/html/register.html", w, nil)
}

// Log in page
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("login")
		password := r.FormValue("password")

		user, err := sqlite.CheckLogin(username, password)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
			return
		}

		slog.Info("Successfully logged in", "User_name", user) // IF USERS PASS IS CORRECT THEN REDIRECT TO

		http.Redirect(w, r, "/users", http.StatusSeeOther)
	}

	RenderTemplate("./web/html/login.html", w, nil)
}

// Log out
func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// IF USERS PASS IS CORRECT THEN REDIRECT TO
		//http.Redirect(w, r, "/"+username+"/posts", http.StatusSeeOther)
	}
	RenderTemplate("./web/html/login.html", w, nil)
}

func generateToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		log.Fatalf("Failed to generate token: %v", err)
	}
	return base64.URLEncoding.EncodeToString(b)
}
