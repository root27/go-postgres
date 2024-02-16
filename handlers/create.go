package handlers

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/root27/go-postgres/models"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx, db *sql.DB) error {

	var user models.User

	e := c.BodyParser(&user)

	if e != nil {
		log.Fatal(e)

		c.JSON(fiber.Map{
			"message": "Error parsing request",
		})

		return e
	}

	_, err := db.Query("CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, name VARCHAR(50), age INT)")

	if err != nil {
		log.Fatal(err)

		c.JSON(fiber.Map{
			"message": "Error creating table",
		})

		return err
	}

	_, err = db.Query("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)

	if err != nil {
		log.Fatal(err)

		c.JSON(fiber.Map{
			"message": "Error creating user",
		})

		return err
	}

	c.JSON(fiber.Map{
		"message": "User created successfully",
	})

	return nil

}
