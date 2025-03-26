package service

import (
	"billing-api/domain"
	"billing-api/dto"
	"billing-api/queue"
	"encoding/json"
	"log/slog"
)

type QueueService struct {
	is domain.InvoiceService

	send *queue.RabbitMQQueue
	recv *queue.RabbitMQQueue
}

func NewQueueService(is domain.InvoiceService, send *queue.RabbitMQQueue, recv *queue.RabbitMQQueue) *QueueService {
	return &QueueService{
		is:   is,
		send: send,
		recv: recv,
	}
}

func (qs *QueueService) Send(invoiceId uint) error {
	ips, err := qs.is.GetInvoiceProductsById(invoiceId)
	if err != nil {
		return err
	}
	msgInBytes, err := json.Marshal(ips)
	if err != nil {
		slog.Error("QueueService#Send", "err", err)
		return err
	}
	err = qs.send.Publish(msgInBytes)
	if err != nil {
		slog.Error("QueueService#Send", "err", err)
		return err
	}

	qs.is.UpdateStatus(invoiceId, string(domain.Processing))
	return nil
}

func (qs *QueueService) Listen() {
	messages := qs.recv.GetMessages()

	for message := range messages {
		var dtoRecv dto.InvoiceProductRecvDto
		json.Unmarshal(message.Body, &dtoRecv)
		slog.Info("QueueService#Listen", "recv", message.Body, "converted", dtoRecv)

		if dtoRecv.InvoiceId == 0 {
			continue
		}

		if dtoRecv.Status != "CLOSED" {
			// TODO: FindById and get InvoiceProducts for resend
		}
		err := qs.is.UpdateStatus(dtoRecv.InvoiceId, dtoRecv.Status)
		if err != nil {
			slog.Error("QueueService#Listen", "err", err)
		}
	}
}
