package services

import (
	"microservice-poc/products/api/rest/schemas"
	"microservice-poc/products/internal/exceptions"
	"microservice-poc/products/internal/models"
	"microservice-poc/products/internal/repository"
)

func CreateProduct(data *schemas.CreateProduct) (*models.Product, error) {
	productWithSameTitle, err := repository.FindProductByTitle(data.Title)

	if err != nil && err != exceptions.ErrProductNotFound {
		return nil, err
	}

	if productWithSameTitle != nil {
		return nil, exceptions.ErrProductNotFound
	}

	productModel := &models.Product{
		Title: data.Title,
		Price: data.Price,
	}

	err = repository.CreateProduct(productModel)

	if err != nil {
		return nil, err
	}

	return productModel, nil
}
