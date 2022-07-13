package handler

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/krissukoco/Sparepartku-API/config"
)

func GenerateJWT(userId string) string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("!!! PROVIDE JWT_SECRET_KEY ENV !!!")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * time.Duration(config.JWT_EXPIRY_HOURS)),
	})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatal("ERROR SIGNING JWT TOKEN: ", err.Error())
	}
	return tokenString
}
