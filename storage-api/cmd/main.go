package main

import (
	"log"
	"storage-api/api"
	"storage-api/config"

	"github.com/gofiber/fiber/v3"
)

func main() {
	env := config.NewEnv()
	db := config.NewPostgresDatabase(env)
	app := fiber.New()

	api.Setup(app, db)

	app.Get("/health", func(c fiber.Ctx) error {
		_, err := c.WriteString("ok")
		return err
	})

	log.Fatal(app.Listen(env.WebServerPort))
}
