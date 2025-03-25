package route

import (
	"storage-api/api/controller"
	"storage-api/domain"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func NewProductRouter(db *gorm.DB, r fiber.Router) {
	var ps domain.ProductService
	pc := controller.NewProductController(ps)
	r.Post("/products", pc.Create)
}
