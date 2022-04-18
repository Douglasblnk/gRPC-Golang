package services

import (
	"products/internal/models"
	"products/internal/repository"
)

func FindAllProducts() ([]*models.Product, error) {
	products, err := repository.FindAllProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}
