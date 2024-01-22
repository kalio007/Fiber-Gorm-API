package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kalio007/Fiber-Gorm-API/database"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to this Awesome API")
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	app.Get("/api", welcome)
	log.Fatal(app.Listen(":3000"))
}
