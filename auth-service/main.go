package main

import (
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	DB *Database
)

func getToken(c *gin.Context) {
	log.Println("Server: request for new token")
	var req *User
	if err := c.BindJSON(req); err != nil {
		log.Println("Server: Can't parse request body:", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "Пробелмы с обработкой запроса",
			},
		)
		return
	}

	log.Println("Server: Checking login ", req.Login)
	if DB.IsAuthorized(req) {
		token, err := genToken(req.Login)
		if err != nil {
			log.Println("Server: Can't authorize this user: ", err.Error())
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"error": "Неудачная авторизация",
				},
			)
			return
		}
		c.JSON(
			http.StatusOK,
			token,
		)
	} else {
		c.JSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "Неудачная авторизация",
			},
		)
		return
	}

}

func main() {
	r := gin.Default()
	r.POST("/", getToken)
}

func genToken(login string) (*Tokens, error) {
	hmacSampleSecret := os.Getenv("SECRET")
	AccessTokenExp := time.Now().Add(time.Minute * 30).Unix()
	RefreshTokenExp := time.Now().Add(time.Hour * 30).Unix()
	accesToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"iss":   Iss,
		"exp":   AccessTokenExp,
		"aud":   Aud,
	})
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"iss":   Iss,
		"exp":   RefreshTokenExp,
		"aud":   Aud,
	})
	accessTokenString, err := accesToken.SignedString(hmacSampleSecret)
	if err != nil {
		return nil, err
	}
	refreshTokenString, err := refreshToken.SignedString(hmacSampleSecret)

	return &Tokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
		ExpiresIn:    AccessTokenExp,
	}, nil
}
