package middleware

import (
	"billing-api/exception"
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

func ErrorMiddleware(c fiber.Ctx) error {
	err := c.Next()
	if err == nil {
		return nil
	}

	return handleErrBase(c, err)
}

func handleErrBase(c fiber.Ctx, err error) error {
	slog.Error("ErrorMiddleware", "err", err)
	if errBase, ok := err.(*exception.ErrorBase); ok {
		return c.Status(errBase.Status).
			JSON(map[string]any{
				"message": errBase.Message,
			})
	}

	return c.Status(fiber.StatusInternalServerError).
		JSON(map[string]any{
			"message": err.Error(),
		})
}
