package graphql

import (
	"context"

	"github.com/GigioSanches/Desafio3Golang/internal/domain"
	"github.com/GigioSanches/Desafio3Golang/internal/usecase"
)

type Resolver struct {
	ListOrdersUC *usecase.ListOrdersUseCase
}

func (r *Resolver) ListOrders(ctx context.Context) ([]domain.Order, error) {
	return r.ListOrdersUC.Execute()
}
