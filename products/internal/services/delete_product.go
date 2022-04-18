package services

import (
	"products/internal/repository"
)

func DeleteProduct(id string) error {
	product, err := repository.FindProductByID(id)

	if err != nil {
		return err
	}

	if err = repository.DeleteProduct(product); err != nil {
		return err
	}

	return nil
}
