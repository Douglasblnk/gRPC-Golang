package schemas

type CreateProduct struct {
	Title string `json:"title" validate:"required"`
	Price int    `json:"price" validate:"required"`
}

type UpdateProduct struct {
	Title *string `json:"title"`
	Price *int    `json:"price"`
}
