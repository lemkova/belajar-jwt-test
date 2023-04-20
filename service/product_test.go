package service

import (
	"jwt-h8/models"
	"jwt-h8/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var productRepository = &repository.ProductRepositoryMock{
	Mock:         mock.Mock{},
	ProductStore: map[uint]*models.Product{},
	Id_counter:   0,
}

func TestCreateProduct(t *testing.T) {
	productRepository.Mock.On("CreateProduct", mock.Anything).Return(nil)

	product := &models.Product{
		Title:       "test",
		Description: "test",
		UserID:      1,
	}

	err := productRepository.CreateProduct(product)

	assert.Nil(t, err)
}

func TestCreateInvalidProduct(t *testing.T) {
	productRepository.Mock.On("CreateProduct", mock.Anything).Return(nil)

	product := &models.Product{
		Title:       "",
		Description: "test",
		UserID:      1,
	}

	err := productRepository.CreateProduct(product)

	assert.NotNil(t, err)
}

func TestUpdateProduct(t *testing.T) {
	productRepository.Mock.On("UpdateProduct", mock.Anything).Return(nil)

	product := &models.Product{
		Title:       "test",
		Description: "test update",
		UserID:      1,
	}

	product.ID = 1

	err := productRepository.UpdateProduct(product)

	assert.Nil(t, err)
}

func TestUpdateEmptyProduct(t *testing.T) {
	productRepository.Mock.On("UpdateProduct", mock.Anything).Return(nil)

	product := &models.Product{
		Title:       "test",
		Description: "test update",
		UserID:      1,
	}

	product.ID = 2

	err := productRepository.UpdateProduct(product)

	assert.NotNil(t, err)
}
