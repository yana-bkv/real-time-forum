package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"jwt-authentication/database"
	"jwt-authentication/models"
	"net/http"
	"strconv"
	"time"
)

const SecretKey = "secret"

func Register(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	//database is package, CreateUser is function, DB is *sql.DB, &user is *models.User
	err := database.CreateUser(database.DB, &user)
	if err != nil {
		if err.Error() == "email already taken" {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Example response (You can process `data` and insert it into a DB)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Send the received 'data' back as JSON response
	if err := json.NewEncoder(w).Encode("Success"); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	// Decode JSON request body
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// GET USER info BASED ON EMAIL
	user, err := database.GetUserByEmail(database.DB, data["email"])
	if err != nil {
		http.Error(w, "Email not found", http.StatusNotFound)
		return
	}

	if user.Id == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		http.Error(w, "Incorrect password", http.StatusBadRequest)
		return
	}

	// Create token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
	}

	// Create cookie
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    token,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)

	// Send the received 'data' back as JSON response
	if err := json.NewEncoder(w).Encode("Success"); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "Unauthorized: You must be logged in to access this resource", http.StatusUnauthorized)
		return
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		http.Error(w, "Error creating token", http.StatusUnauthorized)
	}
	claims := token.Claims.(*jwt.StandardClaims)

	// Put user info to user variable from database
	// token has user id and it finds user by its id
	user, err := database.GetUserById(database.DB, claims.Issuer)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)
	json.NewEncoder(w).Encode("Success logout")
}
