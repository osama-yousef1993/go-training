package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type AuthMiddleware struct{}
type Generate struct {
	Token string `json:""`
}

func GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error :: %s", err.Error())
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := to.SignedString([]byte(jwtSecretKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error :: %s", err.Error())
	}
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (auth *AuthMiddleware) IsAuthorized(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isAuth := false
		au := r.Header.Get("Authorization")
		if au == " " {
			http.Error(w, "you must pass the token", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(au, " ")
		if len(tokenParts) < 2 {
			http.Error(w, "you must pass the token", http.StatusUnauthorized)
			return
		}
		token, err := ValidateToken(tokenParts[1])
		if err != nil {
			http.Error(w, "you must pass the token", http.StatusUnauthorized)
			return
		}
		if email, ok := token.Claims.(jwt.MapClaims)["email"].(string); ok {
			if strings.HasSuffix(email, "@hotmail.com") || strings.HasSuffix(email, "@gmail.com") {
				isAuth = true
			}
		}
		if !isAuth {
			fmt.Println(" user not authorized")
			http.Error(w, "you must pass the token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
