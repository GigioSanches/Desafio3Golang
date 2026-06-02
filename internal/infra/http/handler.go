package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/GigioSanches/Desafio3Golang/internal/domain"
	"github.com/GigioSanches/Desafio3Golang/internal/usecase"
)

type OrderHandler struct {
	ListOrdersUC *usecase.ListOrdersUseCase
	OrderRepo    domain.OrderRepository
}

func NewOrderHandler(listOrdersUC *usecase.ListOrdersUseCase) *OrderHandler {
	return &OrderHandler{ListOrdersUC: listOrdersUC, OrderRepo: listOrdersUC.OrderRepo}
}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.ListOrdersUC.Execute()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order domain.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid payload"})
		return
	}
	if order.CreatedAt.IsZero() {
		order.CreatedAt = time.Now()
	}
	err := h.OrderRepo.CreateOrder(order)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
