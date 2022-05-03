package schemas

type CreateOrder []struct {
	Qtd       int    `json:"qtd" validate:"required"`
	ProductID string `json:"productID" validate:"required"`
}

// type UpdateOrder struct {
// 	Title *string `json:"title"`
// 	Price *int    `json:"price"`
// }
