package service

import (
	"billing-api/domain"
	"billing-api/dto"
	"billing-api/exception"

	"github.com/peteprogrammer/go-automapper"
)

var _ domain.InvoiceService = &InvoiceService{}

type InvoiceService struct {
	ir domain.InvoiceRepository
}

func NewInvoiceService(ir domain.InvoiceRepository) *InvoiceService {
	return &InvoiceService{ir}
}

func (is *InvoiceService) Create(creatingDto *dto.InvoiceCreateDto) (*dto.InvoiceReadDto, error) {
	existingInvoice := new(domain.Invoice)
	existingInvoice, err := is.ir.FindByNumeration(creatingDto.Numeration)
	if err != nil {
		return nil, err
	}
	if existingInvoice != nil {
		return nil, exception.NewErrBadRequest("Numeration already created")
	}

	invoice := new(domain.Invoice)
	automapper.MapLoose(creatingDto, invoice)
	invoice.Items = makeInvoiceProduct(creatingDto.Products)

	err = is.ir.Create(invoice)
	if err != nil {
		return nil, err
	}
	readingDto := new(dto.InvoiceReadDto)
	automapper.MapLoose(invoice, readingDto)
	readingDto.Items = makeInvoiceProductDto(invoice.Items)

	return readingDto, nil
}

func makeInvoiceProduct(dtos []dto.InvoiceProductDto) []*domain.InvoiceProduct {
	ips := make([]*domain.InvoiceProduct, len(dtos))
	for i := range ips {
		ips[i] = &domain.InvoiceProduct{
			InvoiceID: 0,
			ProductID: dtos[i].ID,
			Quantity:  dtos[i].Quantity,
		}
	}
	return ips
}

func makeInvoiceProductDto(entities []*domain.InvoiceProduct) []dto.InvoiceProductDto {
	ips := make([]dto.InvoiceProductDto, len(entities))
	for i := range ips {
		ips[i] = dto.InvoiceProductDto{
			ID:       entities[i].ProductID,
			Quantity: entities[i].Quantity,
		}
	}
	return ips
}

func (is *InvoiceService) GetAll(page, size int) (*dto.Page[dto.InvoiceReadDto], error) {
	if page < 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	invoices, total, err := is.ir.FindAll(page, size)
	if err != nil {
		return nil, err
	}

	responseDtos := make([]dto.InvoiceReadDto, len(invoices))
	automapper.MapLoose(invoices, &responseDtos)
	for i, v := range invoices {
		responseDtos[i].Items = makeInvoiceProductDto(v.Items)
	}

	return &dto.Page[dto.InvoiceReadDto]{
		Content: responseDtos,
		Page:    page,
		Size:    size,
		Total:   total,
	}, nil
}
