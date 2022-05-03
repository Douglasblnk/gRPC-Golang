package repository

import (
	"microservice-poc/orders/internal/database"
	"microservice-poc/orders/internal/models"
)

func CreateOrder(order *models.Order) error {
	if err := database.DB.Create(order).Error; err != nil {
		return err
	}

	return nil
}
