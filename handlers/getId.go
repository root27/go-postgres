package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/root27/go-postgres/models"
)

func GetId(c *fiber.Ctx, db *sql.DB) error {

	id := c.Params("id")

	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)

	var user models.User

	err := row.Scan(&user.ID, &user.Name, &user.Age)

	if err != nil {
		log.Fatal(err)

		c.JSON(fiber.Map{
			"message": "Error getting user",
		})

		return err
	}

	c.JSON(fiber.Map{
		"user": user,
	})

	return nil

}
