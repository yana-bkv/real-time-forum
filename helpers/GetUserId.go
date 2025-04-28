package helpers // Return user name
import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

const SecretKey = "secret"

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

	return claims.Issuer
}
