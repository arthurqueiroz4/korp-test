package service

import (
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
