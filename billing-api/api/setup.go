package api

import (
	"billing-api/api/middleware"
	"billing-api/api/route"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
	"gorm.io/gorm"
)

func Setup(app *fiber.App, db *gorm.DB) {
	app.Use(cors.New())
	app.Use(recoverer.New())
	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "America/Sao_Paulo",
	}))

	app.Use(middleware.ErrorMiddleware)

	defaultRouter := app.Group("/api")

	route.NewInvoiceRouter(db, defaultRouter)
}
