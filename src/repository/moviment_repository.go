package repository

import (
	"github.com/mutinho/src/models"
	"gorm.io/gorm"
)

type MovimentRepository struct {
	db *gorm.DB
}

func NewMovimentRepository(db *gorm.DB) *MovimentRepository {
	return &MovimentRepository{db}
}

func (r *MovimentRepository) Create(moviment *models.Moviment) error {
	return r.db.Create(moviment).Error
}

func (r *MovimentRepository) FindAll() ([]models.Moviment, error) {
	var moviments []models.Moviment
	err := r.db.Preload("Producto").Preload("Usuario").Find(&moviments).Error
	return moviments, err
}

func (r *MovimentRepository) FindProductBySlug(slug string) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, "slug = ?", slug).Error
	return &product, err
}

func (r *MovimentRepository) FindUserBySlug(slug string) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, "slug = ?", slug).Error
	return &user, err
}
