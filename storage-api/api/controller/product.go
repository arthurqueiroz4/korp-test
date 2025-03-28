package controller

import (
	"storage-api/domain"
	"storage-api/dto"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type ProductController struct {
	ps domain.ProductService
}

func NewProductController(ps domain.ProductService) *ProductController {
	return &ProductController{ps}
}

func (pc *ProductController) Create(c fiber.Ctx) error {
	dto := new(dto.ProductCreateDto)

	if err := c.Bind().JSON(dto); err != nil {
		return err
	}
	responseDto, err := pc.ps.Create(dto)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).
		JSON(responseDto)
}

func (pc *ProductController) GetAll(c fiber.Ctx) error {
	pageStr := c.Query("page", "0")
	sizeStr := c.Query("size", "10")
	name := c.Query("name")
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return err
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return err
	}

	responseDtos, err := pc.ps.GetAll(page, size, name)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).
		JSON(responseDtos)
}

func (pc *ProductController) Delete(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	err = pc.ps.Delete(id)
	if err != nil {
		return nil
	}
	return c.SendStatus(fiber.StatusNoContent)
}
