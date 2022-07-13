package router

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/krissukoco/Sparepartku-API/api/handler"
	"github.com/krissukoco/Sparepartku-API/config"
)

func returnUnauthorized(c *fiber.Ctx) error {
	return c.Status(401).JSON(handler.ErrorRes{
		Message:   "UNAUTHORIZED: Please provide valid Bearer Token",
		ErrorCode: "ERR:UNAUTHORIZED:JWT",
	})
}

func RequireAndParseJWT(c *fiber.Ctx) error {
	bearer := c.GetReqHeaders()["Authorization"]
	if bearer == "" {
		return returnUnauthorized(c)
	}
	split := strings.Split(bearer, " ")
	if len(split) != 2 {
		return returnUnauthorized(c)
	}
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(split[1], claims, func(token *jwt.Token) (interface{}, error) {
		secretKey := []byte(config.GetJWTSecretKey())
		return secretKey, nil
	})
	if err != nil {
		fmt.Println(err.Error())
		return returnUnauthorized(c)
	}
	c.Locals("user_id", claims["user_id"])
	// fmt.Printf("user_id: %s", c.Locals("user_id"))
	return c.Next()
}
