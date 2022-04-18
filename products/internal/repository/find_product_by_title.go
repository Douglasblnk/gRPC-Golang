package repository

import (
	"products/internal/database"
	"products/internal/exceptions"
	"products/internal/models"
)

func FindProductByTitle(title string) (*models.Product, error) {
	product := &models.Product{}

	if err := database.DB.Where("title = ?", title).First(product).Error; err != nil {
		return nil, exceptions.ErrProductNotFound
	}

	return product, nil
}
