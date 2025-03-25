package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		_, err := c.WriteString("ok")
		return err
	})

	log.Fatal(app.Listen(":3000"))
}
