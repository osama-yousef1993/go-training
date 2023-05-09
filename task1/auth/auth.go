package auth

import (
	"fmt"
	"go-training/task1/log"
	"net/mail"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

const (
	key = "secret"
)

func GenerateToken() string {
	mailAddress := os.Getenv("email")
	validMailAddress(mailAddress)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = mailAddress
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(key))
	checkErrors(err, "error while creating token ")
	tokenString = "token is: " + tokenString
	log.Info(tokenString)
	return tokenString
}

func validMailAddress(email string) {
	_, err := mail.ParseAddress(email)
	checkErrors(err, "invalid email address")
}

func checkErrors(err error, message string) {
	if err != nil {
		log.Err(message)
		os.Exit(1)
	}
}

func validateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// Return the public key for verification
		return key, nil
	})

	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	} else {
		return false, nil
	}
}
