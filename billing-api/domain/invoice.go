package domain

import (
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
	Numeration string           `gorm:"type:varchar(50)"`
	Status     Status           `gorm:"type:varchar(20);default:'OPENED'"`
	Items      []InvoiceProduct `gorm:"foreignKey:InvoiceID"`
}

type InvoiceProduct struct {
	InvoiceID uint `gorm:"not null;index"`
	ProductID uint `gorm:"not null;index"`
}
