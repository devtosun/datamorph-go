package main

import (
	"github.com/gofiber/fiber/v2"
	data "github.com/devtosun/datamorph-go/database"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Datamorph  👋!")
    })

	data.ConnectDB()

    app.Listen(":3000")
}