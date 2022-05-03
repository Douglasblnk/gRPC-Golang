package integration

import (
	"context"
	pb "microservice-poc/orders/proto"
)

func GetProduct(productID string) (*pb.ProductResponse, error) {
	conn, err := Connect("3300")

	if err != nil {
		return nil, err
	}

	defer conn.Close()

	productService := pb.NewProductServiceClient(conn)

	product, err := productService.FindProduct(context.TODO(), &pb.ProductRequest{
		Id: productID,
	})

	return product, nil
}
