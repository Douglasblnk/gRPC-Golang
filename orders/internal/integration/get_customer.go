package integration

import (
	"context"
	pb "microservice-poc/orders/proto"
)

func GetCustomer(customerID string) (*pb.CustomerResponse, error) {
	conn, err := Connect("3200")

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	customerService := pb.NewCustomerServiceClient(conn)

	customer, err := customerService.FindCustomer(context.TODO(), &pb.CustomerRequest{
		Id: customerID,
	})

	return customer, nil
}
