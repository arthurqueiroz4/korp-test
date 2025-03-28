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
	GetAll(page, size int, name string) (*dto.Page[dto.ProductReadDto], error)
	Delete(id int) error

	ValidateQuantity(ips []dto.InvoiceProductDto) error
	UpdateBalance(ips []dto.InvoiceProductDto) error
}

type ProductRepository interface {
	Create(p *Product) error
	FindByName(name string) (*Product, error)
	FindAll(page, size int, name string) ([]Product, int, error)
	Delete(id int) error
	FindAllByIds(ids []uint) ([]Product, error)
	UpdateBalance(discountById map[uint]int) error
}
