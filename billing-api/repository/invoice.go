package repository

import (
	"billing-api/domain"
	"errors"

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

func (ir *InvoiceRepository) FindAll(page, size int) ([]domain.Invoice, error) {
	var invoices []domain.Invoice
	offset := (page - 1) * size

	result := ir.db.
		Offset(offset).
		Limit(size).
		Preload("Items").
		Find(&invoices)

	if result.Error != nil {
		return nil, result.Error
	}

	return invoices, nil
}
