package handlers

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/root27/go-postgres/models"
)

func GetAll(c *fiber.Ctx, db *sql.DB) error {
	rows, err := db.Query("SELECT * from users")

	if err != nil {
		log.Fatal(err)

		c.JSON(fiber.Map{
			"message": "Error getting users",
		})
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(&user.ID, &user.Name, &user.Age)

		if err != nil {
			log.Fatal(err)

			c.JSON(fiber.Map{
				"message": "Error scanning users",
			})
		}

		users = append(users, user)
	}

	c.JSON(fiber.Map{
		"users": users,
	})

	return nil

}
