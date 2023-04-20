package repository

import (
	"errors"
	"jwt-h8/models"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	Mock         mock.Mock
	ProductStore map[uint]*models.Product
	Id_counter   uint
}

func (p *ProductRepositoryMock) CreateProduct(product *models.Product) error {
	args := p.Mock.Called(product)
	if _, err := govalidator.ValidateStruct(product); err != nil {
		return err
	}

	if args.Get(0) == nil {
		p.Id_counter += 1
		product.ID = p.Id_counter
		p.ProductStore[product.ID] = product
		return nil
	}

	return args.Error(0)
}

func (p *ProductRepositoryMock) UpdateProduct(product *models.Product) error {
	args := p.Mock.Called(product)
	if _, err := govalidator.ValidateStruct(product); err != nil {
		return err
	}

	if _, ok := p.ProductStore[product.ID]; !ok {
		return errors.New("product not found")
	}

	if args.Get(0) == nil {
		p.ProductStore[product.ID] = product
		return nil
	}

	return args.Error(0)
}
