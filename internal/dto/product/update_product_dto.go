package dto

type UpdateProductDTO struct {
	Id          string
	Name        string
	Description string
	Image       string
	Value       float64
	CategoryId  string
}
