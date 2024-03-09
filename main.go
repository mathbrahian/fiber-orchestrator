package main

import (
	"fiber-orchestrator/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/orchestrator/ping", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("ping")
	})

	app.Post("/orchestrator", handlers.RegisterSportmen)

	app.Listen(":80")
}
