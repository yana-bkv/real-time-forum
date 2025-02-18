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

	http.Handle("/", r)
}
