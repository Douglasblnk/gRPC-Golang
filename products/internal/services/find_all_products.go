package services

import (
	"microservice-poc/products/internal/models"
	"microservice-poc/products/internal/repository"
)

func FindAllProducts() ([]*models.Product, error) {
	products, err := repository.FindAllProducts()

	if err != nil {
		return nil, err
	}

	return products, nil
}
