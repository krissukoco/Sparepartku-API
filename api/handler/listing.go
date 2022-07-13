package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/krissukoco/Sparepartku-API/models"
)

func GetListing(c *fiber.Ctx) error {
	return c.JSON(models.Listing{})
}

func CreateListing(c *fiber.Ctx) error {
	return c.JSON(models.Listing{})
}

func UpdateListing(c *fiber.Ctx) error {
	return c.JSON(models.Listing{})
}

func DeleteListing(c *fiber.Ctx) error {
	return c.JSON(models.Listing{})
}
