package usecase

import "github.com/GigioSanches/Desafio3Golang/internal/domain"

type ListOrdersUseCase struct {
	OrderRepo domain.OrderRepository
}

func (uc *ListOrdersUseCase) Execute() ([]domain.Order, error) {
	return uc.OrderRepo.ListOrders()
}
