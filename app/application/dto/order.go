package dto

type CreateOrderRequest struct {
	UserId     int64 `json:"user_id"`
	OrderItems []struct {
		ProductId string  `json:"product_id"`
		Quantity  int     `json:"quantity"`
		Price     float64 `json:"price"`
	} `json:"order_items"`
}
