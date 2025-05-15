package routes

import (
	"github.com/gorilla/mux"
	"jwt-authentication/controllers"
	"jwt-authentication/repositories"
	"jwt-authentication/services"
	"jwt-authentication/websocket"
	"net/http"
)

func Setup(r *mux.Router) {
	// Initialize repositories
	userRepo := repositories.NewUserRepository()
	postRepo := repositories.NewPostRepository()
	categoryRepo := repositories.NewCategoryRepository()
	postCategoryRepo := repositories.NewPostCategoryRepository()
	commentRepo := repositories.NewCommentRepository()
	likeRepo := repositories.NewLikeRepository()
	msgRepo := repositories.NewMessageRepository()

	// Initialize services
	authService := services.NewAuthService(userRepo)
	postService := services.NewPostService(postRepo)
	commentService := services.NewCommentService(commentRepo)
	msgService := services.NewMessageService(msgRepo)

	// Initialize controllers
	authController := controllers.NewAuthController(authService)
	postController := controllers.NewPostController(postService)
	categoryController := controllers.NewCategoryController(categoryRepo)
	postCategoryController := controllers.NewPostCategoryController(postCategoryRepo)
	commentController := controllers.NewCommentController(commentService)
	likeController := controllers.NewLikeController(likeRepo)
	msgController := controllers.NewMessageController(msgService)

	r.HandleFunc("/api/register", authController.Register).Methods("POST")
	r.HandleFunc("/api/login", authController.Login).Methods("POST")
	r.HandleFunc("/api/logout", authController.Logout).Methods("POST")

	r.HandleFunc("/api/user", authController.GetAuthUser).Methods("GET") // Authed user
	//r.HandleFunc("/api/user/{id}", controllers.GetUserById).Methods("GET")
	r.HandleFunc("/api/users", authController.GetUsers).Methods("GET")

	r.HandleFunc("/api/post", postController.Create).Methods("POST")
	r.HandleFunc("/api/post/{id}", postController.GetPost).Methods("GET")
	r.HandleFunc("/api/posts", postController.GetPosts).Methods("GET")
	r.HandleFunc("/api/post/{id}", postController.Delete).Methods("DELETE")

	r.HandleFunc("/api/category", categoryController.Create).Methods("POST")
	r.HandleFunc("/api/category/{id}", categoryController.Get).Methods("GET")
	r.HandleFunc("/api/categories", categoryController.GetAll).Methods("GET")
	r.HandleFunc("/api/category/{id}", categoryController.Delete).Methods("DELETE")

	r.HandleFunc("/api/post/{id}/category", postCategoryController.Create).Methods("POST")
	r.HandleFunc("/api/post/{id}/category/{cid}", postCategoryController.Get).Methods("GET")
	r.HandleFunc("/api/post/{id}/categories", postCategoryController.GetAll).Methods("GET")
	r.HandleFunc("/api/post/{id}/category/{cid}", postCategoryController.Delete).Methods("DELETE")

	r.HandleFunc("/api/post/{id}/comment", commentController.Create).Methods("POST")
	r.HandleFunc("/api/post/{id}/comments", commentController.GetCommentsByPostId).Methods("GET")
	r.HandleFunc("/api/post/{id}/comment/{cId}", commentController.Delete).Methods("DELETE")

	r.HandleFunc("/api/post/{id}/like", likeController.AddLikeToPost).Methods("POST")
	r.HandleFunc("/api/post/{id}/likes", likeController.GetLikesByPostId).Methods("GET")
	r.HandleFunc("/api/post/{id}/likeCount", likeController.GetUserCount).Methods("GET")
	r.HandleFunc("/api/post/{id}/like", likeController.Delete).Methods("DELETE")
	r.HandleFunc("/api/post/{id}/hasLiked", likeController.HasLiked).Methods("GET")

	r.HandleFunc("/api/messages/{user}/{peer}", msgController.Get).Methods("GET")

	// websocket
	r.HandleFunc("/ws/{user}/{peer}", websocket.ServeWs)
	go websocket.HubInstance.Run()

	http.Handle("/api", r)
}
