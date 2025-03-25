package route

import (
	"billing-api/api/controller"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func NewInvoiceRouter(db *gorm.DB, r fiber.Router) {
	ic := controller.NewInvoiceController()
	r.Post("/invoices", ic.Create)
}
