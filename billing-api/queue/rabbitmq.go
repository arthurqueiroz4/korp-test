package queue

import (
	"log"
	"log/slog"

	"github.com/streadway/amqp"
)

type (
	Messages <-chan amqp.Delivery
)

type RabbitMQQueue struct {
	ch        *amqp.Channel
	conn      *amqp.Connection
	queueName string
}

type RabbitMQConfig struct {
	AMPQServerUrl string
	QueueName     string
}

func NewRabbitMQ(config RabbitMQConfig) *RabbitMQQueue {
	connectRabbitMQ, err := amqp.Dial(config.AMPQServerUrl)
	if err != nil {
		panic(err)
	}

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}

	_, err = channelRabbitMQ.QueueDeclare(
		config.QueueName, // queue name
		true,             // durable
		false,            // auto delete
		false,            // exclusive
		false,            // no wait
		nil,              // arguments
	)
	if err != nil {
		panic(err)
	}
	slog.Info("Queue created", "name", config.QueueName)

	return &RabbitMQQueue{
		ch:        channelRabbitMQ,
		conn:      connectRabbitMQ,
		queueName: config.QueueName,
	}
}

func (q *RabbitMQQueue) Publish(msg []byte) error {
	message := amqp.Publishing{
		ContentType: "application/json",
		Body:        msg,
	}

	if err := q.ch.Publish(
		"",          // exchange
		q.queueName, // queue name
		false,       // mandatory
		false,       // immediate
		message,     // message to publish
	); err != nil {
		return err
	}
	return nil
}

func (q *RabbitMQQueue) GetMessages() Messages {
	messages, err := q.ch.Consume(
		q.queueName, // queue name
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no local
		false,       // no wait
		nil,         // arguments
	)
	if err != nil {
		log.Println(err)
	}
	return messages
}
