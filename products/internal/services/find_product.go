package services

import (
	"microservice-poc/products/internal/models"
	"microservice-poc/products/internal/repository"
)

func FindProduct(id string) (*models.Product, error) {
	product, err := repository.FindProductByID(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
