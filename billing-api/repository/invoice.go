package repository

import (
	"billing-api/domain"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var _ domain.InvoiceRepository = &InvoiceRepository{}

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{db}
}

func (ir *InvoiceRepository) Create(i *domain.Invoice) error {
	err := ir.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("Items").Create(&i).Error; err != nil {
			return err
		}
		addInvoiceIdInInvoiceProducts(i.ID, i.Items)
		if err := tx.Create(&i.Items).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func addInvoiceIdInInvoiceProducts(invoiceId uint, ips []*domain.InvoiceProduct) {
	for _, v := range ips {
		v.InvoiceID = invoiceId
	}
}

func (ir *InvoiceRepository) FindByNumeration(n string) (*domain.Invoice, error) {
	var invoice domain.Invoice
	err := ir.db.
		Where("numeration = ?", n).
		First(&invoice).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &invoice, nil
}

func (ir *InvoiceRepository) FindAll(page, size int) ([]domain.Invoice, int, error) {
	var invoices []domain.Invoice
	offset := page * size

	result := ir.db.
		Offset(offset).
		Limit(size).
		Preload("Items").
		Find(&invoices)

	if result.Error != nil {
		return nil, 0, result.Error
	}

	var count int64
	ir.db.Model(&domain.Invoice{}).Count(&count)

	return invoices, int(count), nil
}

func (ir *InvoiceRepository) UpdateStatus(id uint, status string, detail string) error {
	result := ir.db.Model(&domain.Invoice{}).
		Where("id = ?", id).
		Update("status", status).
		Update("detail", detail)

	if result.Error != nil {
		return fmt.Errorf("failed to update invoice status: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("invoice with ID %d not found", id)
	}

	return nil
}

func (ir *InvoiceRepository) FindInvoiceProductsById(id uint) ([]domain.InvoiceProduct, error) {
	var invoiceProducts []domain.InvoiceProduct

	err := ir.db.Where("invoice_id = ?", id).Find(&invoiceProducts).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find invoice products: %w", err)
	}

	if len(invoiceProducts) == 0 {
		return nil, fmt.Errorf("no products found for invoice ID %d", id)
	}

	return invoiceProducts, nil
}
