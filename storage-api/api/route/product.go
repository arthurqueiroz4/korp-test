package route

import (
	"storage-api/api/controller"
	"storage-api/repository"
	"storage-api/service"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func NewProductRouter(db *gorm.DB, r fiber.Router) {
	pr := repository.NewPostgresRepository(db)
	ps := service.NewProductService(pr)
	pc := controller.NewProductController(ps)
	r.Post("/products", pc.Create)
	r.Get("/products", pc.GetAll)
}
