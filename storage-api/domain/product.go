package domain

import (
	"storage-api/dto"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"type:varchar(128)"`
	Description string `gorm:"type:varchar(256)"`
	Balance     int
}

type ProductService interface {
	Create(dto *dto.ProductCreateDto) (*dto.ProductReadDto, error)
}

type ProductRepository interface {
	Create(p *Product) error
	FindByName(name string) (*Product, error)
}
