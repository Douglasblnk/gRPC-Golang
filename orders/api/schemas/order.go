package schemas

type CreateOrder struct {
	Title string `json:"title" validate:"required"`
	Price int    `json:"price" validate:"required"`
}

type UpdateOrder struct {
	Title *string `json:"title"`
	Price *int    `json:"price"`
}
