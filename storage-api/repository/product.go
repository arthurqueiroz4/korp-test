package repository

import (
	"errors"
	"log/slog"
	"storage-api/domain"

	"gorm.io/gorm"
)

// Force implemention in comptime
var _ domain.ProductRepository = &ProductRepository{}

type ProductRepository struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (pr *ProductRepository) Create(p *domain.Product) error {
	result := pr.db.Create(p)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (pr *ProductRepository) FindByName(name string) (*domain.Product, error) {
	var product domain.Product
	result := pr.db.Where("name = ?", name).First(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		slog.Error("Error on query", "err", result.Error)
		return nil, result.Error
	}
	return &product, nil
}

func (pr *ProductRepository) FindAll(page, size int) ([]domain.Product, error) {
	var products []domain.Product
	offset := (page - 1) * size

	result := pr.db.
		Offset(offset).
		Limit(size).
		Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}
