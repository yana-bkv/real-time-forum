package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"jwt-authentication/database"
	"jwt-authentication/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// connect db
	database.ConnectDB()

	frontendPort := ":3000"
	backendPort := ":8080"

	// setup routing
	backendRouter := mux.NewRouter()
	routes.Setup(backendRouter)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Start the backend server (API) with CORS middleware
	go func() {
		fmt.Println("Backend server running on http://localhost" + backendPort + "/api/user")
		log.Fatal(http.ListenAndServe(backendPort, corsHandler.Handler(backendRouter))) // Apply CORS handler here
	}()

	// Serve static files for the frontend (React, Vue, or any SPA)
	frontendHandler := http.StripPrefix("/", http.FileServer(http.Dir("./public/")))

	// Catch-all route to serve index.html for all non-static requests (SPA routing)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Check if the requested file exists, otherwise serve the SPA index.html
		_, err := os.Stat("./public" + r.URL.Path)
		if os.IsNotExist(err) {
			// If file doesn't exist, serve the SPA's index.html
			http.ServeFile(w, r, "./public/index.html")
		} else {
			// Serve static file (if it exists)
			frontendHandler.ServeHTTP(w, r)
		}
	})

	// Start the frontend server
	fmt.Println("Frontend server running on http://localhost" + frontendPort)
	log.Fatal(http.ListenAndServe(frontendPort, nil))
}
