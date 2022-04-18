package services

import (
	"products/api/schemas"
	"products/internal/models"
	"products/internal/repository"
)

func UpdateProduct(id string, data *schemas.UpdateProduct) (*models.Product, error) {
	product, err := repository.FindProductByID(id)

	if err != nil {
		return nil, err
	}

	if data.Title != nil {
		product.Title = *data.Title
	}

	if data.Price != nil {
		product.Price = *data.Price
	}

	err = repository.UpdateProduct(product)

	if err != nil {
		return nil, err
	}

	return product, nil
}
