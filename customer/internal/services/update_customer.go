package services

import (
	"customer/api/schemas"
	"customer/internal/models"
	"customer/internal/repository"
)

func UpdateCustomer(id string, data *schemas.UpdateCustomer) (*models.Customer, error) {
	customer, err := repository.FindCustomerByID(id)

	if err != nil {
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
