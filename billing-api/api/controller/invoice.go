package controller

import "github.com/gofiber/fiber/v3"

type InvoiceController struct{}

func NewInvoiceController() *InvoiceController {
	return &InvoiceController{}
}

func (ic *InvoiceController) Create(c fiber.Ctx) error {
	c.WriteString("InvoiceController#Create")
	return nil
}
