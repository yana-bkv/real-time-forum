package routes

import (
	"github.com/gorilla/mux"
	"jwt-authentication/controllers"
	"net/http"
)

func Setup(r *mux.Router) {
	r.HandleFunc("/api/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")
	r.HandleFunc("/api/user", controllers.User).Methods("GET")
	r.HandleFunc("/api/logout", controllers.Logout).Methods("POST")

	r.HandleFunc("/api/posts", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/api/post/{id}", controllers.GetPost).Methods("GET")
	r.HandleFunc("/api/posts", controllers.GetPosts).Methods("GET")
	r.HandleFunc("/api/post/{id}", controllers.DeletePost).Methods("DELETE")

	http.Handle("/api", r)
}
