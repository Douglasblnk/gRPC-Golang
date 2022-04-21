package exceptions

import "errors"

var ErrOrderNotFound = errors.New("pedido não encontrado")
var ErrOrderAlreadyExists = errors.New("pedido já existe")
