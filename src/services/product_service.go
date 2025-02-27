package services

import (
	"github.com/google/uuid"
	"github.com/mutinho/src/dtos"
	"github.com/mutinho/src/models"
	"github.com/mutinho/src/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) CreateProduct(req dtos.CreateProductRequest) (*dtos.ProductResponse, error) {
	id := uuid.New()

	product := &models.Product{
		ID:          id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, err
	}

	return &dtos.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Slug:        product.Slug,
	}, nil

}
