package main
import (
	"github.com/golang-jwt/jwt"
	"fmt"
	"time"
	"os"
	"strings"
)
var mySigningKey = []byte("keyfromenv")

type mainClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

var claims = mainClaims{
	"hello@outlook.com",
	jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer: "test",
	},
}

func genjwt() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
	return ss, err

	
}


func valjwt(tok string, key string) {
	at(time.Unix(0, 0), func() {
		token, err := jwt.ParseWithClaims(tok, &mainClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		})

		email_check :="@outlook.com"



		if claims, ok := token.Claims.(*mainClaims); ok && token.Valid && strings.Contains(claims.Email, email_check) {
			fmt.Printf("%v %v %v", claims.Email, claims.StandardClaims.ExpiresAt, claims.StandardClaims.Issuer)
		} else {
			if !strings.Contains(claims.Email, email_check) {
				fmt.Println("You are not authorized ")
			} else {
				fmt.Println(err)
			}
		}
	})


}


func at(t time.Time, f func()) {
	jwt.TimeFunc = func() time.Time {
		return t 
	}
	f()
	jwt.TimeFunc = time.Now
}

func main() {
    var key string = os.Getenv("KEY") 
	t, _ := genjwt()
	fmt.Println()
    valjwt(t , key)
}

