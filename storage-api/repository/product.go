package repository

import (
	"errors"
	"fmt"
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

func (pr *ProductRepository) FindAll(page, size int) ([]domain.Product, int, error) {
	var products []domain.Product
	offset := page * size

	result := pr.db.
		Offset(offset).
		Limit(size).
		Find(&products)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	var count int64
	pr.db.Model(&domain.Product{}).Count(&count)

	return products, int(count), nil
}

func (pr *ProductRepository) Delete(id int) error {
	result := pr.db.Delete(&domain.Product{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (pr *ProductRepository) FindAllByIds(ids []uint) ([]domain.Product, error) {
	var products []domain.Product

	if len(ids) == 0 {
		return products, nil
	}

	err := pr.db.Where("id IN ?", ids).Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find products by ids: %w", err)
	}

	return products, nil
}

func (pr *ProductRepository) UpdateBalance(discountById map[uint]int) error {
	tx := pr.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for productID, quantity := range discountById {
		result := tx.Model(&domain.Product{}).
			Where("id = ? AND balance >= ?", productID, quantity).
			Update("balance", gorm.Expr("balance - ?", quantity))

		if result.Error != nil {
			tx.Rollback()
			return fmt.Errorf("failed to update balance for product %d: %w", productID, result.Error)
		}

		if result.RowsAffected == 0 {
			tx.Rollback()
			return fmt.Errorf("product %d not found or insufficient balance", productID)
		}
	}

	return tx.Commit().Error
}
