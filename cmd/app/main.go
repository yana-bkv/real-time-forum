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
	"real-time-forum/internal/transport/rest"
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

	router.HandleFunc("/", rest.IndexHandler) // login or auth
	router.HandleFunc("/login", rest.LoginHandler)
	router.HandleFunc("/auth", rest.CreateUserHandler)
	//http.HandleFunc("/{username}/posts", rest.PostsHandler)
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
