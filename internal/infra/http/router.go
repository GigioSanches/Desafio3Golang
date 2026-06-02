package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(orderHandler *OrderHandler) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/order", orderHandler.ListOrders).Methods("GET")
	r.HandleFunc("/order", orderHandler.CreateOrder).Methods("POST")
	return r
}
