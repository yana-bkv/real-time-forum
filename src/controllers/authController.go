package controllers

import (
	"encoding/json"
	"jwt-authentication/services"
	"net/http"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	err := DecodeJson(r, w, &data)
	if err != nil {
		return
	}

	err = c.authService.Register(data)
	if err != nil {
		if err.Error() == "email or username already taken" {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = EncodeJson(w, "Success")
	if err != nil {
		return
	}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	// Decode JSON request body
	err := DecodeJson(r, w, &data)
	if err != nil {
		return
	}

	cookie, err := c.authService.Login(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	http.SetCookie(w, cookie)

	// Encode response as JSON
	err = EncodeJson(w, "Success")
	if err != nil {
		return
	}
}

func (c *AuthController) GetAuthUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "Unauthorized: You must be logged in to access this resource", http.StatusUnauthorized)
		return
	}

	user, err := c.authService.GetAuthUser(cookie)
	if err != nil {
		http.Error(w, "Unauthorized: You must be logged in to access this resource", http.StatusUnauthorized)
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, user)
	if err != nil {
		return
	}
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := c.authService.Logout()
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode("Success logout")
}

func (c *AuthController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.authService.GetAllUsers()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = EncodeJson(w, users)
	if err != nil {
		return
	}
}
