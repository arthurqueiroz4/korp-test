package config

import (
	"storage-api/queue"
	"storage-api/repository"
	"storage-api/service"

	"gorm.io/gorm"
)

func BuildAndListenQueues(env *Env, db *gorm.DB) {
	pr := repository.NewPostgresRepository(db)
	ps := service.NewProductService(pr)

	recv := queue.NewRabbitMQ(queue.RabbitMQConfig{
		AMPQServerUrl: env.AmqpServerUrl,
		QueueName:     env.QueueNameFromBilling,
	})

	send := queue.NewRabbitMQ(queue.RabbitMQConfig{
		AMPQServerUrl: env.AmqpServerUrl,
		QueueName:     env.QueueNameToBilling,
	})

	qs := service.NewQueueService(ps, send, recv)

	go qs.Listen()
}
