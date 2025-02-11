package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

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

	http.HandleFunc("/", rest.IndexHandler) // login or auth
	http.HandleFunc("/login", rest.LoginHandler)
	http.HandleFunc("/auth", rest.CreateUserHandler)
	http.HandleFunc("/{username}/posts", rest.PostsHandler)
	http.HandleFunc("/user/createPost", rest.CreatePostHandler)
	http.HandleFunc("/user/deletePost", rest.DeletePostHandler)

	fmt.Println("Listening on port 8080")
	// Run server on port 8080

	slog.Info("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error(err.Error())
	}
}
