package auth

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecretKey = "mysecrettokenkey"

func GenerateToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodES256)
	// s := "https://www.googleapis.com/oauth2/v3/certs"

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

func ValidateToken(tokenString string) (bool, error) {
	isAuth := false
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return isAuth, err
	}
	if email, ok := token.Claims.(jwt.MapClaims)["email"].(string); ok {
		if strings.HasSuffix(email, "@forbes.com") || strings.HasSuffix(email, "@gmail.com") {
			isAuth = true
		}
	}
	if !isAuth {
		fmt.Println(" user not authorized")
		return isAuth, nil
	}

	return isAuth, nil
}
