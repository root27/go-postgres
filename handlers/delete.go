package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func Delete(c *fiber.Ctx, db *sql.DB) error {

	id := c.Params("id")

	_, err := db.Query("DELETE FROM users WHERE id = $1", id)

	if err != nil {
		c.JSON(fiber.Map{
			"message": "Failed to delete user",
		})
	}

	c.JSON(fiber.Map{
		"message": "User deleted successfully",
	})

	return nil

}
