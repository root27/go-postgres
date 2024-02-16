package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/gofiber/fiber/v2"

	"github.com/root27/go-postgres/handlers"
)

func main() {

	log.Println("Database connected")

	app := fiber.New()

	app.Get("/test", handlers.Test)
	app.Post("/create", handlers.Create)

	log.Fatal(app.Listen(":3000"))

}
