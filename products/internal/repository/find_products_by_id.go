package repository

import (
	"microservice-poc/products/internal/database"
	"microservice-poc/products/internal/exceptions"
	"microservice-poc/products/internal/models"
)

func FindProductByID(id string) (*models.Product, error) {
	product := &models.Product{}

	if err := database.DB.Where("id = ?", id).First(product).Error; err != nil {
		return nil, exceptions.ErrProductNotFound
	}

	return product, nil
}
