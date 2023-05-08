package auth

import (
	"go-training/log"
	"net/mail"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func GenerateToken() {
	mailAddress := os.Getenv("email")
	validMailAddress(mailAddress)
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = mailAddress
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	checkErrors(err, "error while creating token ")
	tokenString = "token is: " + tokenString
	log.Info(tokenString)
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
