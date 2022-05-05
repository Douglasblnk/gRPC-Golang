package services

import (
	"microservice-poc/orders/api/schemas"
	"microservice-poc/orders/internal/exceptions"
	"microservice-poc/orders/internal/integration"
	"microservice-poc/orders/internal/models"
	"microservice-poc/orders/internal/repository"
	pb "microservice-poc/orders/proto"
)

func getProduct(data *schemas.Order) ([]*pb.ProductResponse, error) {
	var product []*pb.ProductResponse

	for _, p := range data.Items {
		prod, err := integration.GetProduct(p.ProductID)

		if err != nil {
			return nil, err
		}

		product = append(product, prod)
	}

	return product, nil
}

func mountProducts(products []*pb.ProductResponse) []models.OrderItem {
	items := []models.OrderItem{}

	for _, product := range products {
		items = append(items, models.OrderItem{
			ProductID: product.Id,
		})
	}

	return items
}

func CreateOrder(data *schemas.Order) (*models.Order, error) {
	customer, err := integration.GetCustomer(data.CustomerID)

	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, exceptions.ErrCustomerNotFound
	}

	products, err := getProduct(data)

	if err != nil {
		return nil, err
	}

	for _, product := range products {
		if product == nil {
			return nil, exceptions.ErrProductNotFound
		}
	}

	order := &models.Order{
		CustomerID:    customer.Id,
		CustomerName:  customer.Name,
		CustomerEmail: customer.Email,
		Items:         mountProducts(products),
	}

	err = repository.CreateOrder(order)

	if err != nil {
		return nil, err
	}

	return order, nil
}
