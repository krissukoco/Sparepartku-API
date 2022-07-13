package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/krissukoco/Sparepartku-API/api/router"
	"github.com/krissukoco/Sparepartku-API/database"
	"github.com/krissukoco/Sparepartku-API/models"
)

func main() {
	start_db := time.Now()
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("UNABLE TO CONNECT TO DATABASE")
	}
	db.AutoMigrate(&models.Person{})
	db.AutoMigrate(&models.Listing{})
	db.AutoMigrate(&models.ListingImage{})
	pgdb, _ := db.DB()
	defer pgdb.Close()
	end_db := time.Now()
	fmt.Printf("CONNECTED TO DATABASE! in %v", end_db.Sub(start_db))

	app := fiber.New()
	router.UseDefaultRouter(app)

	app.Listen(":3000")
}
