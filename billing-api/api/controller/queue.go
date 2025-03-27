package controller

import (
	"billing-api/service"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type QueueController struct {
	qs *service.QueueService
}

func NewQueueController(qs *service.QueueService) *QueueController {
	return &QueueController{qs}
}

func (qc *QueueController) Enqueue(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return err
	}

	err = qc.qs.Send(uint(id))
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": "Invoice queued successfully",
	})
}
