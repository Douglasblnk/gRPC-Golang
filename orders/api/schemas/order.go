package schemas

type Order struct {
	CustomerID string     `json:"customerId" validate:"required"`
	Items      OrderItems `json:"items" validate:"required"`
}

type OrderItems []struct {
	Qtd       int    `json:"qtd" validate:"required"`
	ProductID string `json:"productId" validate:"required"`
}

// type UpdateOrder struct {
// 	Title *string `json:"title"`
// 	Price *int    `json:"price"`
// }
