package handlers

import (
	"net/http"
	"real-time-forum/internal/database/sqlite"
	"real-time-forum/internal/models"
	"text/template"
)

// POSTS
// Render all posts to authorized users
func PostsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := sqlite.DB.Query("SELECT id, title, body, author FROM posts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := []models.Post{}
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.Author); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	tmpl := template.Must(template.ParseFiles("./web/html/posts.html"))
	tmpl.Execute(w, posts)
}

//// Create post
//func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Method == "POST" {
//		title := r.FormValue("title") // Get the title from the form data
//		body := r.FormValue("body")
//		author := r.FormValue("author")
//		_, err := sqlite.DB.Exec("INSERT INTO posts(title,body,author) VALUES(?,?,?)", title, body, author)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError) // Return an HTTP 500 error if insertion fails
//			return
//		}
//		http.Redirect(w, r, "/user/posts", http.StatusSeeOther) // Redirect to the main page after successful creation
//	}
//}
//
//// Delete post
//func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
//	id := r.URL.Query().Get("id")                                  // Get the id from the URL query parameters
//	_, err := sqlite.DB.Exec("DELETE FROM posts WHERE id = ?", id) // Delete the todo from the database
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError) // Return an HTTP 500 error if deletion fails
//		return
//	}
//	http.Redirect(w, r, "/user/posts", http.StatusSeeOther) // Redirect to the main page after successful deletion
//
//}
