package controller

import (
	"storage-api/domain"

	"github.com/gofiber/fiber/v3"
)

type ProductController struct {
	ps domain.ProductService
}

func NewProductController(ps domain.ProductService) *ProductController {
	return &ProductController{ps}
}

func (pc *ProductController) Create(c fiber.Ctx) error {
	c.WriteString("ProductController#Create")
	return nil
}
