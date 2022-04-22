package services

import (
	"microservice-poc/orders/api/schemas"
	"microservice-poc/orders/internal/config"
	"microservice-poc/orders/internal/models"
	pbcustomer "microservice-poc/proto"
)

func CreateOrder(data *schemas.CreateOrder) (*models.Order, error) {
	conn := config.Connection

	customer := pbcustomer.NewCustomerServiceClient(conn)

	customer.GetCustomer()
}
