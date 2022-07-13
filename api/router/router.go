package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krissukoco/Sparepartku-API/api/handler"
)

func UseDefaultRouter(app *fiber.App) {
	api := app.Group("/api/v1")

	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/signup", handler.Register)

	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	// user.Post("/", handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)

	listing := api.Group("/listing")
	listing.Get("/:id", handler.GetListing)
	listing.Post("/", handler.CreateListing)
	listing.Put("/:id", handler.UpdateListing)
	listing.Delete("/:id", handler.DeleteListing)
}
