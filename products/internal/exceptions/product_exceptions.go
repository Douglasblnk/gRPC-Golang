package exceptions

import "errors"

var ErrProductNotFound = errors.New("produto não encontrado")
var ErrProductAlreadyExists = errors.New("produto já existe")
