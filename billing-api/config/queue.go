package config

import (
	"billing-api/queue"
	"billing-api/repository"
	"billing-api/service"

	"gorm.io/gorm"
)

func BuildAndListenQueues(env *Env, db *gorm.DB) *service.QueueService {
	ir := repository.NewInvoiceRepository(db)
	is := service.NewInvoiceService(ir)

	recv := queue.NewRabbitMQ(queue.RabbitMQConfig{
		AMPQServerUrl: env.AmqpServerUrl,
		QueueName:     env.QueueNameFromStorage,
	})

	send := queue.NewRabbitMQ(queue.RabbitMQConfig{
		AMPQServerUrl: env.AmqpServerUrl,
		QueueName:     env.QueueNameToStorage,
	})

	qs := service.NewQueueService(is, send, recv)
	go qs.Listen()

	return qs
}
