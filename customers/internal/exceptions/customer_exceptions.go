package exceptions

import "errors"

var ErrCustomerNotFound = errors.New("cliente não encontrado")
var ErrCustomerAlreadyExists = errors.New("cliente já existe")
