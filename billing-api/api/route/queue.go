package route

import (
	"billing-api/api/controller"
	"billing-api/config"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func NewQueueRouter(env *config.Env, db *gorm.DB, r fiber.Router) {
	qs := config.BuildAndListenQueues(env, db)
	qc := controller.NewQueueController(qs)
	r.Post("/invoices/enqueue/:id", qc.Enqueue)
}
