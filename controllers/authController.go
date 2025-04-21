package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"jwt-authentication/database"
	"jwt-authentication/models"
	"jwt-authentication/repositories"
	"net/http"
	"strconv"
	"time"
)

// for jwt token
const SecretKey = "secret"

func Register(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	// Decode JSON request body
	err := DecodeJson(r, w, data)
	if err != nil {
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	//database is package, CreateUser is function, DB is *sql.DB, &user is *models.User
	err = repositories.CreateUser(database.DB, &user)
	if err != nil {
		if err.Error() == "email already taken" {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	// Encode response as JSON
	err = EncodeJson(w, "Success")
	if err != nil {
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var data map[string]string

	// Decode JSON request body
	err := DecodeJson(r, w, &data)
	if err != nil {
		return
	}

	if data["username"] == "" && data["email"] == "" {
		fmt.Println(data)
		http.Error(w, "username or email is required", http.StatusUnauthorized)
		return
	}

	user, err := repositories.GetUserByUsername(database.DB, data["username"])
	if err != nil {
		user, err = repositories.GetUserByEmail(database.DB, data["email"])
		if err != nil {
			http.Error(w, "Email not found", http.StatusNotFound)
			return
		}
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

	// Encode response as JSON
	err = EncodeJson(w, "Success")
	if err != nil {
		return
	}
}

func GetAuthUser(w http.ResponseWriter, r *http.Request) {
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
	user, err := repositories.GetUserById(database.DB, claims.Issuer)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	// Encode response as JSON
	err = EncodeJson(w, user)
	if err != nil {
		return
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

// Return user name
func GetUserId(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		http.Error(w, "Unauthorized: You must be logged in to access this resource", http.StatusUnauthorized)
		return ""
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		http.Error(w, "Error creating token", http.StatusUnauthorized)
	}
	claims := token.Claims.(*jwt.StandardClaims)

	//// Put user info to user variable from database
	//// token has user id and it finds user by its id
	//user, _ := database.GetUserById(database.DB, claims.Issuer)
	//
	//if err := json.NewEncoder(w).Encode(user); err != nil {
	//	http.Error(w, "Error encoding response", http.StatusInternalServerError)
	//}

	return claims.Issuer
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Fetch post from database
	users, err := repositories.GetAllUsers(database.DB)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Post not found", http.StatusNotFound)
		} else {
			fmt.Println(err, users)
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	// Encode response as JSON
	err = EncodeJson(w, users)
	if err != nil {
		return
	}
}
