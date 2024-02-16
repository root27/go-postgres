package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Test(c *fiber.Ctx, db *sql.DB) error {

	c.JSON(fiber.Map{
		"message": "Database connected",
	})

	return nil

}
