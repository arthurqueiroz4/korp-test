package controller

import (
	"billing-api/domain"
	"billing-api/dto"

	"github.com/gofiber/fiber/v3"
)

type InvoiceController struct {
	is domain.InvoiceService
}

func NewInvoiceController(is domain.InvoiceService) *InvoiceController {
	return &InvoiceController{is}
}

func (ic *InvoiceController) Create(c fiber.Ctx) error {
	dto := new(dto.InvoiceCreateDto)

	if err := c.Bind().JSON(dto); err != nil {
		return err
	}
	responseDto, err := ic.is.Create(dto)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).
		JSON(responseDto)
}
