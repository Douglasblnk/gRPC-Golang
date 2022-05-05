package repository

import (
	"microservice-poc/orders/internal/database"
	"microservice-poc/orders/internal/exceptions"
	"microservice-poc/orders/internal/models"
)

func FindOrderByID(id string) (*models.Order, error) {
	order := &models.Order{}

	if err := database.DB.Preload("Items").Where("id = ?", id).First(order).Error; err != nil {
		return nil, exceptions.ErrOrderNotFound
	}

	return order, nil
}
