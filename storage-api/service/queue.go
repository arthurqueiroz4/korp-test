package service

import (
	"encoding/json"
	"log/slog"
	"storage-api/domain"
	"storage-api/dto"
	"storage-api/queue"
)

type QueueService struct {
	ps domain.ProductService

	send *queue.RabbitMQQueue
	recv *queue.RabbitMQQueue
}

func NewQueueService(ps domain.ProductService, send *queue.RabbitMQQueue, recv *queue.RabbitMQQueue) *QueueService {
	return &QueueService{
		ps:   ps,
		send: send,
		recv: recv,
	}
}

func (qs *QueueService) Listen() {
	messages := qs.recv.GetMessages()

	for message := range messages {
		var dtos []dto.InvoiceProductDto
		json.Unmarshal(message.Body, &dtos)
		slog.Info("QueueService#Listen", "recv", message.Body, "converted", dtos)

		if len(dtos) == 0 {
			continue
		}

		err := qs.ps.UpdateBalance(dtos)
		if err != nil {
			slog.Error("QueueService#Listen", "error", err)
			continue
		}

		err = qs.Send(dtos)
		if err != nil {
			continue
		}
	}
}

func (qs *QueueService) Send(dtos []dto.InvoiceProductDto) error {
	dtoToSend := dto.MessageForBilling{
		InvoiceID: dtos[0].InvoiceID,
		Status:    "CLOSED",
	}

	msg, err := json.Marshal(dtoToSend)
	if err != nil {
		slog.Error("QueueService#Send", "error", err)
		return err
	}
	err = qs.send.Publish(msg)
	if err == nil {
		slog.Info("QueueService#Send", "send", msg, "converted", dtoToSend)
	}

	return nil
}
