package repository

import (
	"microservice-poc/orders/internal/database"
	"microservice-poc/orders/internal/exceptions"
	"microservice-poc/orders/internal/models"
)

func FindAllOrders() ([]*models.Order, error) {
	orders := []*models.Order{}

	if err := database.DB.Preload("Items").Find(&orders).Error; err != nil {
		return nil, exceptions.ErrOrderNotFound
	}

	return orders, nil
}
