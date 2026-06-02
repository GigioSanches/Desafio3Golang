package domain

import "time"

type Order struct {
	ID          string    `json:"id"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

type OrderRepository interface {
	ListOrders() ([]Order, error)
	CreateOrder(order Order) error
}
