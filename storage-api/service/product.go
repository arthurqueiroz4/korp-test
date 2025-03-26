package service

import (
	"fmt"
	"log/slog"
	"storage-api/domain"
	"storage-api/dto"
	"storage-api/exception"

	"github.com/peteprogrammer/go-automapper"
)

// Force implemention in comptime
var _ domain.ProductService = &ProductService{}

type ProductService struct {
	pr domain.ProductRepository
}

func NewProductService(pr domain.ProductRepository) *ProductService {
	return &ProductService{pr}
}

func (ps *ProductService) Create(creatingDto *dto.ProductCreateDto) (*dto.ProductReadDto, error) {
	if creatingDto.Balance < 0 {
		return nil, exception.NewErrBadRequest("", "balance cannot be negative")
	}

	existingProduct, err := ps.pr.FindByName(creatingDto.Name)
	if err != nil {
		return nil, err
	}
	if existingProduct != nil {
		return nil, exception.NewErrBadRequest("", "product with this name already exists")
	}
	p := new(domain.Product)
	automapper.MapLoose(creatingDto, p)

	err = ps.pr.Create(p)
	if err != nil {
		return nil, exception.NewErrBadRequest("", err.Error())
	}

	readingDto := new(dto.ProductReadDto)
	automapper.Map(p, readingDto)
	return readingDto, nil
}

func (ps *ProductService) GetAll(page, size int) (*dto.Page[dto.ProductReadDto], error) {
	if page < 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}

	products, total, err := ps.pr.FindAll(page, size)
	if err != nil {
		return nil, err
	}

	responseDtos := make([]dto.ProductReadDto, len(products))
	automapper.Map(products, &responseDtos)

	return &dto.Page[dto.ProductReadDto]{
		Content: responseDtos,
		Page:    page,
		Size:    size,
		Total:   total,
	}, nil
}

func (ps *ProductService) Delete(id int) error {
	_ = ps.pr.Delete(id)
	return nil
}

func (ps *ProductService) ValidateQuantity(ips []dto.InvoiceProductDto) error {
	if len(ips) == 0 {
		return nil
	}

	ids := make([]uint, len(ips))
	for i, ip := range ips {
		ids[i] = ip.ProductID
	}

	products, err := ps.pr.FindAllByIds(ids)
	if err != nil {
		return err
	}
	slog.Info("ProductService#ValidateQuantity", "products", products)
	stockByProductID := make(map[uint]int, len(products))
	for _, product := range products {
		stockByProductID[product.ID] = product.Balance
	}

	for _, ip := range ips {
		availableStock, exists := stockByProductID[ip.ProductID]
		if !exists {
			return exception.NewErrBadRequest("", fmt.Sprintf("Product with ID %d not found", ip.ProductID))
		}
		if ip.Quantity > stockByProductID[ip.ProductID] {
			return exception.NewErrBadRequest("", fmt.Sprintf("Insufficient stock for product %d. Requested: %d, Available: %d",
				ip.ProductID, ip.Quantity, availableStock))
		}
	}

	return nil
}

func (ps *ProductService) UpdateBalance(ips []dto.InvoiceProductDto) error {
	if len(ips) == 0 {
		return nil
	}
	discountById := make(map[uint]int, len(ips))
	for _, ip := range ips {
		discountById[ip.ProductID] = ip.Quantity
	}
	if err := ps.ValidateQuantity(ips); err != nil {
		return fmt.Errorf("validation failed: %w", err)
	}
	err := ps.pr.UpdateBalance(discountById)
	if err != nil {
		return err
	}
	return nil
}
