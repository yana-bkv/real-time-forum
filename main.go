package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"jwt-authentication/database"
	"jwt-authentication/routes"
	"net/http"
)

func main() {
	// connect db
	database.ConnectDB()

	// setup routing
	r := mux.NewRouter()
	routes.Setup(r)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(r)

	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", handler)
}
