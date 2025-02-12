package handlers

import (
	"net/http"
	"real-time-forum/internal/database/sqlite"
)

// Shows list of all users
func ShowUsers(w http.ResponseWriter, r *http.Request) {
	users, err := sqlite.ShowUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	// Render the template with user data
	RenderTemplate("./web/html/users.html", w, users)
}

// Shows ONE user
func ShowUser(w http.ResponseWriter, r *http.Request) {
	//sqlite.ShowUsers()
	RenderTemplate("./web/html/users.html", w, nil)
}

// Delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//sqlite.ShowUsers(w, r)
	RenderTemplate("./web/html/users.html", w, nil)
}
