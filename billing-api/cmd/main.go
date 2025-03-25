package main

import (
	"billing-api/config"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	env := config.NewEnv()
	_ = config.NewPostgresDatabase(env)
	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		_, err := c.WriteString("ok")
		return err
	})

	log.Fatal(app.Listen(env.WebServerPort))
}
