package handlers

import (
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

		user, err := sqlite.LogUser(username, password)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, "Failed to fetch user", http.StatusInternalServerError)
			return
		}

		slog.Info("Successfully logged in", "input_username", username, "db_username", user) // IF USERS PASS IS CORRECT THEN REDIRECT TO
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
