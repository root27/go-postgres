package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Update(c *fiber.Ctx, db *sql.DB) error {

	id := c.Params("id")

	var data Data

	e := c.BodyParser(&data)

	if e != nil {
		c.JSON(fiber.Map{
			"message": "Error parsing request",
		})

		return e
	}

	_, err := db.Query("UPDATE users SET name = $1, age = $2 WHERE id = $3", data.Name, data.Age, id)

	if err != nil {
		c.JSON(fiber.Map{
			"message": "Error updating user",
		})

		return err
	}

	c.JSON(fiber.Map{
		"message": "User updated successfully",
	})

	return nil

}
