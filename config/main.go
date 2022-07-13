package config

import (
	"log"
	"os"
)

var (
	JWT_EXPIRY_HOURS = 24
)

func GetJWTSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("!!! PROVIDE JWT_SECRET_KEY ENV !!!")
	}
	return secretKey
}
