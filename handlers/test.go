package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/root27/go-postgres/database"
)

func Test(c *fiber.Ctx) error {

	db := database.Connect()

	defer db.Close()

	log.Println("Database connected")

	c.JSON(fiber.Map{
		"message": "Database connected",
	})

	return nil

}
