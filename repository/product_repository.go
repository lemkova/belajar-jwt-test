package repository

import (
	"jwt-h8/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	UpdateProduct(product *models.Product) error
}

type productRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepositoryImpl {
	return &productRepositoryImpl{db}
}

func (p *productRepositoryImpl) CreateProduct(product *models.Product) error {
	return p.db.Create(product).Error
}

func (p *productRepositoryImpl) UpdateProduct(product *models.Product) error {
	return p.db.Save(product).Error
}
