package services

import (
	"microservice-poc/orders/internal/dto"
	"microservice-poc/orders/internal/repository"
)

func FindOrder(id string) (*dto.OrderResponse, error) {
	order, err := repository.FindOrderByID(id)

	if err != nil {
		return nil, err
	}

	products, err := GetItemsFromOrder(order)

	if err != nil {
		return nil, err
	}

	orderResponse := &dto.OrderResponse{
		ID:            order.ID,
		CustomerID:    order.CustomerID,
		CustomerName:  order.CustomerName,
		CustomerEmail: order.CustomerEmail,
		Items:         products,
	}

	return orderResponse, nil
}
