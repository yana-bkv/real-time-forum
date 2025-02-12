package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"real-time-forum/internal/database/sqlite"
	"real-time-forum/internal/handlers"
)

func main() {
	// Initialized databases
	sqlite.InitUserDB()
	sqlite.InitPostDB()
	defer sqlite.DB.Close()

	// Set logger
	handlerOpts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, handlerOpts))
	slog.SetDefault(logger)

	// Router
	router := mux.NewRouter()
	// Connect css
	fs := http.FileServer(http.Dir("web"))
	router.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the file extension
		ext := filepath.Ext(r.URL.Path)
		// Set the correct MIME type
		mimeType := mime.TypeByExtension(ext)
		if mimeType != "" {
			w.Header().Set("Content-Type", mimeType)
		}
		// Serve the file
		fs.ServeHTTP(w, r)
	})))

	router.HandleFunc("/", handlers.Index) // choose login or auth
	router.HandleFunc("/register", handlers.Register)
	router.HandleFunc("/login", handlers.Login)
	router.HandleFunc("/logout", handlers.Logout)
	router.HandleFunc("/users", handlers.ShowUsers).Methods("GET")
	router.HandleFunc("/user/{id}", handlers.ShowUser).Methods("GET")
	router.HandleFunc("/user/{id}", handlers.DeleteUser).Methods("DELETE")
	http.HandleFunc("/posts", handlers.PostsHandler)
	//http.HandleFunc("/{username}/createPost", rest.CreatePostHandler)
	//http.HandleFunc("/{username}/deletePost", rest.DeletePostHandler)

	fmt.Println("Listening on port 8060")
	// Run server on port 8080

	slog.Info("Server is running on port 8060")
	err := http.ListenAndServe(":8060", router)
	if err != nil {
		slog.Error(err.Error())
	}
}
