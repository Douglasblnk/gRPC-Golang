package services

import (
	"microservice-poc/customers/api/schemas"
	"microservice-poc/customers/internal/exceptions"
	"microservice-poc/customers/internal/models"
	"microservice-poc/customers/internal/repository"
)

func CreateCustomer(data *schemas.CreateCustomer) (*models.Customer, error) {
	customerWithSameEmail, err := repository.FindCustomerByEmail(data.Email)

	if err != nil && err != exceptions.ErrCustomerNotFound {
		return nil, err
	}

	if customerWithSameEmail != nil {
		return nil, exceptions.ErrCustomerAlreadyExists
	}

	customerModel := &models.Customer{
		Name:  data.Name,
		Email: data.Email,
	}

	err = repository.CreateCustomer(customerModel)

	if err != nil {
		return nil, err
	}

	return customerModel, nil
}
