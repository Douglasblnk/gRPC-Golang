package schemas

type CreateCustomer struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateCustomer struct {
	Name  *string `json:"name"`
	Email *string `json:"email" validate:"email"`
}
