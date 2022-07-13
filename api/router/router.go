package router

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/krissukoco/Sparepartku-API/api/handler"
)

func UseDefaultRouter(app *fiber.App) {
	app.Use(logger.New())
	api := app.Group("/api/v1")

	// Public routes
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/signup", handler.Register)

	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	listing := api.Group("/listing")
	listing.Get("/:id", handler.GetListing)

	// Protected routes
	jwtSigningKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSigningKey == "" {
		log.Fatal("!!! PROVIDE JWT_SECRET_KEY ENV !!!")
	}
	user.Use(RequireAndParseJWT)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)

	listing.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(jwtSigningKey),
	}))
	listing.Use(RequireAndParseJWT)
	listing.Post("/", handler.CreateListing)
	listing.Put("/:id", handler.UpdateListing)
	listing.Delete("/:id", handler.DeleteListing)
}
