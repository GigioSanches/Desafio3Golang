package db

import (
	"database/sql"

	"github.com/GigioSanches/Desafio3Golang/internal/domain"
)

type OrderRepositoryDB struct {
	DB *sql.DB
}

func NewOrderRepositoryDB(db *sql.DB) *OrderRepositoryDB {
	return &OrderRepositoryDB{DB: db}
}

func (r *OrderRepositoryDB) ListOrders() ([]domain.Order, error) {
	rows, err := r.DB.Query("SELECT id, description, price, created_at FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	orders := []domain.Order{}
	for rows.Next() {
		var o domain.Order
		if err := rows.Scan(&o.ID, &o.Description, &o.Price, &o.CreatedAt); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}
	return orders, nil
}

func (r *OrderRepositoryDB) CreateOrder(order domain.Order) error {
	_, err := r.DB.Exec("INSERT INTO orders (id, description, price, created_at) VALUES ($1, $2, $3, $4)", order.ID, order.Description, order.Price, order.CreatedAt)
	return err
}
