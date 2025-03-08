package routes

import (
	"github.com/gorilla/mux"
	"jwt-authentication/controllers"
	"net/http"
)

func Setup(r *mux.Router) {
	r.HandleFunc("/api/register", controllers.Register).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")
	r.HandleFunc("/api/logout", controllers.Logout).Methods("POST")

	r.HandleFunc("/api/user", controllers.GetAuthUser).Methods("GET") // Authed user
	//r.HandleFunc("/api/user/{id}", controllers.GetUserById).Methods("GET")
	//r.HandleFunc("/api/users", controllers.GetAllUsers).Methods("GET")

	r.HandleFunc("/api/post", controllers.CreatePost).Methods("POST")
	r.HandleFunc("/api/post/{id}", controllers.GetPost).Methods("GET")
	r.HandleFunc("/api/posts", controllers.GetPosts).Methods("GET")
	r.HandleFunc("/api/post/{id}", controllers.DeletePost).Methods("DELETE")

	r.HandleFunc("/api/post/{id}/comment", controllers.CreateComment).Methods("POST")
	r.HandleFunc("/api/post/{id}/comments", controllers.GetCommentsByPostId).Methods("GET")
	r.HandleFunc("/api/post/{id}/comment/{cId}", controllers.DeleteComment).Methods("DELETE")

	r.HandleFunc("/api/post/{id}/like", controllers.AddLike).Methods("POST")
	r.HandleFunc("/api/post/{id}/likes", controllers.GetLikes).Methods("GET")
	r.HandleFunc("/api/post/{id}/like/{lId}", controllers.DeleteLike).Methods("DELETE")

	http.Handle("/api", r)
}
