package dto

type CreateOrderProduct struct {
	ProductId   string  `json:"productId"`
	Quantity    int     `json:"quantity"`
	Observation string  `json:"observation"`
	SubTotal    float64 `json:"subTotal"`
}

type CreateOrderDTO struct {
	Status    string               `json:"status"`
	ClientId  string               `json:"clientId"`
	PaymentId string               `json:"paymentId"`
	Total     float64              `json:"total"`
	Products  []CreateOrderProduct `json:"products"`
}
