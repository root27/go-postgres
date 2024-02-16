package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/root27/go-postgres/database"
)

func Test(c *fiber.Ctx) error {

	db, err := database.Connect()

	if err != nil {
		log.Fatal(err)
		c.JSON(fiber.Map{
			"message": "Error connecting to database",
		})
		return err
	}

	defer db.Close()

	log.Println("Database connected")

	c.JSON(fiber.Map{
		"message": "Database connected",
	})

	return nil

}
