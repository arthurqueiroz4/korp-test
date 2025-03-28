package domain

import "storage-api/dto"

type QueueConfig struct {
	Address string
	Name    string
}

type QueueService interface {
	Listen()
	Send(dtos []dto.InvoiceProductDto, status string, detail string) error
}
