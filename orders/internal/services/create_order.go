package services

import (
	"fmt"
	"microservice-poc/orders/api/schemas"
	"microservice-poc/orders/internal/integration"
	"microservice-poc/orders/internal/models"
	"microservice-poc/orders/internal/repository"
	pb "microservice-poc/orders/proto"
)

func GetProduct(data *schemas.CreateOrder) ([]*pb.ProductResponse, error) {
	var product []*pb.ProductResponse

	for _, p := range *data {
		prod, err := integration.GetProduct(p.ProductID)

		if err != nil {
			return nil, err
		}

		product = append(product, prod)
	}

	return product, nil
}

func MountProducts(products []*pb.ProductResponse) []models.OrderItem {
	items := []models.OrderItem{}

	for _, product := range products {
		items = append(items, models.OrderItem{
			ProductID: product.Id,
		})
	}

	return items
}

func CreateOrder(customerId string, data *schemas.CreateOrder) (*models.Order, error) {
	customer, err := integration.GetCustomer(customerId)

	if err != nil {
		return nil, err
	}

	product, err := GetProduct(data)

	if err != nil {
		return nil, err
	}

	order := &models.Order{
		CustomerID: customer.Id,
		Items:      MountProducts(product),
	}

	fmt.Println("order", order)

	err = repository.CreateOrder(order)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
