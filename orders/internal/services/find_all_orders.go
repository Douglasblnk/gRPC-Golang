package services

import (
	"microservice-poc/orders/internal/dto"
	"microservice-poc/orders/internal/repository"
)

func FindAllOrders() ([]*dto.OrderResponse, error) {
	orders, err := repository.FindAllOrders()

	if err != nil {
		return nil, err
	}

	var orderResponse []*dto.OrderResponse

	for index, order := range orders {
		products, err := GetItemsFromOrder(order)

		if err != nil {
			return nil, err
		}

		orderResponse = append(
			orderResponse[:index],
			&dto.OrderResponse{
				ID:            order.ID,
				CustomerID:    order.CustomerID,
				CustomerName:  order.CustomerName,
				CustomerEmail: order.CustomerEmail,
				Items:         products,
			},
		)
	}

	return orderResponse, nil
}
