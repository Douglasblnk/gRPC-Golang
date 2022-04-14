package services

import (
	"customer/api/schemas"
	"customer/internal/exceptions"
	"customer/internal/models"
	"customer/internal/repository"
)

func UpdateCustomer(data *schemas.UpdateCustomer) (*models.Customer, error) {
	customer, err := repository.FindCustomerByEmail(*data.Email)

	if err != nil && err != exceptions.ErrCustomerNotFound {
		return nil, err
	}

	if data.Email != nil {
		customer.Email = *data.Email
	}

	if data.Name != nil {
		customer.Name = *data.Name
	}

	err = repository.UpdateCustomer(customer)

	if err != nil {
		return nil, err
	}

	return customer, nil
}
