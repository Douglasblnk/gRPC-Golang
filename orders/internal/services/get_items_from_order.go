package services

import (
	"microservice-poc/orders/internal/dto"
	"microservice-poc/orders/internal/integration"
	"microservice-poc/orders/internal/models"
)

func GetItemsFromOrder(order *models.Order) ([]*dto.OrderItemResponse, error) {
	var product []*dto.OrderItemResponse

	for _, p := range order.Items {
		prod, err := integration.GetProduct(p.ProductID)

		if err != nil {
			return nil, err
		}

		orderItemResponse := &dto.OrderItemResponse{
			ProductID: prod.Id,
			Title:     prod.Title,
			Price:     prod.Price,
		}

		product = append(product, orderItemResponse)
	}

	return product, nil
}
