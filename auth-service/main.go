package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	DB *Database
)

func getToken(c *gin.Context) {
	log.Println("Server: request for new token")
	req := &User{}
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

func refreshToken(c *gin.Context) {
	log.Println("Server: Request refresh")
	tokenString := c.Query("refresh_token")
	if tokenString == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// hmacSampleSecret := os.Getenv("SECRET")
		hmacSampleSecret := []byte("secc")
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		newTokens, err := genToken(claims["login"].(string))
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
			newTokens,
		)
	} else {
		log.Println("Gateway: Authorization failed: ", err.Error())
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"error": "Неудачная авторизация",
			},
		)
	}
}

func main() {
	db, err := SetUpDatabase()
	if err != nil {
		log.Panic(err)
	}
	DB = db
	r := gin.Default()
	r.POST("/", getToken)
	r.GET("/", refreshToken)
	r.Run(":8084")
}

func genToken(login string) (*Tokens, error) {
	log.Println("Server: Generating token")
	// hmacSampleSecret := os.Getenv("SECRET")
	hmacSampleSecret := []byte("secc")
	AccessTokenExp := time.Now().Add(time.Minute * 30).Unix()
	RefreshTokenExp := time.Now().Add(time.Hour * 30).Unix()
	log.Println("Server: Gen access token")
	accesToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"iss":   Iss,
		"exp":   AccessTokenExp,
		"aud":   Aud,
	})
	log.Println("Server: Gen refresh token")
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"login": login,
		"iss":   Iss,
		"exp":   RefreshTokenExp,
		"aud":   Aud,
	})

	log.Println("Server: Signing access token", accesToken, hmacSampleSecret)
	accessTokenString, err := accesToken.SignedString(hmacSampleSecret)
	if err != nil {
		return nil, err
	}
	log.Println("Server: Signing refresh token", refreshToken, hmacSampleSecret)
	refreshTokenString, err := refreshToken.SignedString(hmacSampleSecret)

	return &Tokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
