package domain

type QueueConfig struct {
	Address string
	Name    string
}

type QueueService interface {
	Create(qc QueueConfig) error
	Consume()
}
