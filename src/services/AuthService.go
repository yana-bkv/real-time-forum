package services

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"jwt-authentication/models"
	"jwt-authentication/repositories"
	"net/http"
	"strconv"
	"time"
)

// SecretKey for jwt token
const SecretKey = "secret"

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{userRepo: userRepo}
}

func (a *AuthServiceImpl) Register(data map[string]string) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Username: data["username"],
		Email:    data["email"],
		Password: password,
	}

	err := a.userRepo.Create(&user)
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthServiceImpl) Login(data map[string]string) (*http.Cookie, error) {
	if data["username"] == "" && data["email"] == "" {
		fmt.Println(data)
		return nil, errors.New("username or email is required")
	}

	user, err := a.userRepo.GetUserByUsername(data["username"])
	if err != nil {
		user, err = a.userRepo.GetUserByEmail(data["email"])
		if err != nil {
			return nil, errors.New("Email not found")
		}
	}

	if user.Id == 0 {
		return nil, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		return nil, errors.New("incorrect password")
	}

	// Create token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return nil, err
	}

	// Create cookie
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	return cookie, nil
}

func (a *AuthServiceImpl) Logout() *http.Cookie {
	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}
	return cookie
}

func (a *AuthServiceImpl) GetAuthUser(cookie *http.Cookie) (*models.User, error) {
	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, errors.New("Invalid token")
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, err := a.userRepo.GetUserById(claims.Issuer)
	if err != nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (a *AuthServiceImpl) GetAllUsers() ([]models.User, error) {
	users, err := a.userRepo.GetAllUsers()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("User not found")
		}
		return nil, err
	}
	return users, nil
}
