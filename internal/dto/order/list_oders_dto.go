package dto

type ListOrders struct {
	ReceivedOrders  []OrderDTO
	PreparingOrders []OrderDTO
	ReadyOrders     []OrderDTO
	CompletedOrders []OrderDTO
}
