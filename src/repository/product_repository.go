package repository

import (
	"github.com/mutinho/src/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db}
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}
