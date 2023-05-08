package auth

import (
	"fmt"
	"os"
	"regexp"

	jwt "github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	log "github.com/osama-yousef1993/go-training/log"
)

const (
	emailFormat = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
)

type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.LogErr("error loading environment file " + err.Error())
	}
	secretKey := os.Getenv("TOKEN_SECRET")
	claims := CustomClaims{
		email,
		jwt.StandardClaims{
			Issuer: "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	ts, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	log.LogInfo("Successfully generated the token")
	return ts, nil
}

func CheckToken(ts string) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.LogErr("error loading environment file " + err.Error())
	}
	secretKey := os.Getenv("TOKEN_SECRET")
	parsedToken, err := jwt.Parse(ts, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return fmt.Errorf("error parsing token: %s", err.Error())
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		email := claims["email"].(string)

		if !regexp.MustCompile(emailFormat).MatchString(email) {
			return fmt.Errorf("email '%s' format isn't correct", email)
		}
		log.LogInfo("check completed!")
		return nil
	} else {
		return fmt.Errorf("invalid token")
	}
}
