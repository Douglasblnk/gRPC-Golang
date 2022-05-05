package grpcHandlers

import (
	"context"
	"microservice-poc/products/internal/services"
	pb "microservice-poc/products/proto"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
}

func (ps *ProductService) FindProduct(ctx context.Context, data *pb.ProductRequest) (*pb.ProductResponse, error) {
	product, err := services.FindProduct(data.Id)

	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Id:    product.ID,
		Title: product.Title,
		Price: int32(product.Price),
	}, nil
}
