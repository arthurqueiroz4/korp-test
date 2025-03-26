package main

import (
	"billing-api/api"
	"billing-api/config"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	env := config.NewEnv()
	db := config.NewPostgresDatabase(env)
	config.BuildAndListenQueues(env, db)
	app := fiber.New()

	api.Setup(app, db, env)

	app.Get("/health", func(c fiber.Ctx) error {
		_, err := c.WriteString("ok")
		return err
	})

	log.Fatal(app.Listen(env.WebServerPort))
}
