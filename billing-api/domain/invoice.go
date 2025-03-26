package domain

import (
	"billing-api/dto"

	"gorm.io/gorm"
)

type Status string

const (
	Opened     Status = "OPENED"
	Processing Status = "PROCESSING"
	Closed     Status = "CLOSED"
)

type Invoice struct {
	gorm.Model
	Numeration string            `gorm:"type:varchar(50)"`
	Status     string            `gorm:"type:varchar(20);default:'OPENED'"`
	Items      []*InvoiceProduct `gorm:"foreignKey:InvoiceID"`
}

type InvoiceProduct struct {
	InvoiceID uint `gorm:"not null;index"`
	ProductID uint `gorm:"not null;index"`
	Quantity  int
}

type InvoiceService interface {
	Create(creatingDto *dto.InvoiceCreateDto) (*dto.InvoiceReadDto, error)
	GetAll(page, size int) (*dto.Page[dto.InvoiceReadDto], error)
}

type InvoiceRepository interface {
	Create(i *Invoice) error
	FindByNumeration(n string) (*Invoice, error)
	FindAll(page, size int) ([]Invoice, error)
}
