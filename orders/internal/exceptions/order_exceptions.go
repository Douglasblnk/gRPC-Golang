package exceptions

import "errors"

var ErrCustomerNotFound = errors.New("cliente não encontrado")
var ErrProductNotFound = errors.New("produto não encontrado")
var ErrOrderNotFound = errors.New("pedido não encontrado")
var ErrOrderAlreadyExists = errors.New("pedido já existe")
