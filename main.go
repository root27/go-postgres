package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"

	"github.com/root27/go-postgres/handlers"
)

func main() {

	db, err := sql.Open("postgres", "postgres://test:123@127.0.0.1:5432/fiber_test?sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected")

	app := fiber.New()

	app.Post("/create", func(c *fiber.Ctx) error {
		return handlers.Create(c, db)
	})

	app.Get("/getAll", func(c *fiber.Ctx) error {
		return handlers.GetAll(c, db)
	})

	app.Get("/get/:id", func(c *fiber.Ctx) error {
		return handlers.GetId(c, db)
	})

	app.Post("/update/:id", func(c *fiber.Ctx) error {
		return handlers.Update(c, db)
	})

	log.Fatal(app.Listen(":3000"))

}
