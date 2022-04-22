package services

import (
	pbcustomer "microservices/proto"
	"order/api/schemas"
	"order/internal/config"
	"order/internal/models"
)

func CreateOrder(data *schemas.CreateOrder) (*models.Order, error) {
	conn := config.Connection

	customer := pbcustomer.NewCustomerServiceClient(conn)

	customer.GetCustomer()
}
