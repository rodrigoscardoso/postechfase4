package dto

type CreateOrderProductDTO struct {
	OrderId     string
	ProductId   string
	Quantity    int
	SubTotal    float64
	Observation string
}
