package repository

import (
	"microservice-poc/products/internal/database"
	"microservice-poc/products/internal/exceptions"
	"microservice-poc/products/internal/models"
)

func FindAllProducts() ([]*models.Product, error) {
	products := []*models.Product{}

	if err := database.DB.Find(&products).Error; err != nil {
		return nil, exceptions.ErrProductNotFound
	}

	return products, nil
}
