package handlers

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/root27/go-postgres/database"
	"github.com/root27/go-postgres/models"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {

	var user models.User

	e := c.BodyParser(&user)

	if e != nil {
		log.Fatal(e)

		c.JSON(fiber.Map{
			"message": "Error parsing request",
		})

		return e
	}

	log.Printf("User: %v", user)

	db, er := database.Connect()

	if er != nil {
		log.Fatal(er)
		c.JSON(fiber.Map{
			"message": "Error connecting to database",
		})

		return er
	}

	log.Println("Database connected")

	defer db.Close()

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, name VARCHAR(50), age INT)")

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
		"message": "Table created successfully",
	})

	return nil

}
