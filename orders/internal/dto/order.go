package dto

type OrderResponse struct {
	ID            int
	CustomerID    string
	CustomerName  string
	CustomerEmail string
	Items         []*OrderItemResponse
}

type OrderItemResponse struct {
	ID        int
	ProductID string
	Title     string
	Price     int32
	Qtd       int
}
