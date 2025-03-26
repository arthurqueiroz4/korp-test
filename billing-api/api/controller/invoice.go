package controller

import (
	"billing-api/domain"
	"billing-api/dto"
	"strconv"

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

func (ic *InvoiceController) GetAll(c fiber.Ctx) error {
	pageStr := c.Query("page", "1")
	sizeStr := c.Query("size", "10")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return err
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return err
	}
	responseDtos, err := ic.is.GetAll(page, size)
	return c.Status(fiber.StatusOK).
		JSON(responseDtos)
}
