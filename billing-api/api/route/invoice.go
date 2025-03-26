package route

import (
	"billing-api/api/controller"
	"billing-api/repository"
	"billing-api/service"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func NewInvoiceRouter(db *gorm.DB, r fiber.Router) {
	ir := repository.NewInvoiceRepository(db)
	is := service.NewInvoiceService(ir)
	ic := controller.NewInvoiceController(is)
	r.Post("/invoices", ic.Create)
}
