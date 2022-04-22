package services

import (
	"microservice-poc/customers/internal/repository"
)

func DeleteCustomer(id string) error {
	customer, err := repository.FindCustomerByID(id)

	if err != nil {
		return err
	}

	if err = repository.DeleteCustomer(customer); err != nil {
		return err
	}

	return nil
}
