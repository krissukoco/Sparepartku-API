package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krissukoco/Sparepartku-API/database"
	"github.com/krissukoco/Sparepartku-API/models"
)

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(422).JSON(ErrorRes{
			Message:   "Please provide user_id",
			ErrorCode: "ERR:USER_ID:NOT_PROVIDED",
		})
	}
	var user models.Person
	database.DB.First(&user, "id = ?", id)
	if user.ID == "" {
		return c.Status(404).JSON(ErrorRes{
			Message:   "User NOT FOUND",
			ErrorCode: "ERR:USER:NOT_FOUND",
		})
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	return c.JSON(models.Person{})
}

func DeleteUser(c *fiber.Ctx) error {
	return c.JSON(models.Person{})
}
