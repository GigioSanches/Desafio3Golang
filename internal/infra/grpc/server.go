package grpc

import (
	"context"

	pb "github.com/GigioSanches/Desafio3Golang/internal/infra/grpc/proto"
	"github.com/GigioSanches/Desafio3Golang/internal/usecase"
)

type OrderServiceServer struct {
	pb.UnimplementedOrderServiceServer
	ListOrdersUC *usecase.ListOrdersUseCase
}

func NewOrderServiceServer(listOrdersUC *usecase.ListOrdersUseCase) *OrderServiceServer {
	return &OrderServiceServer{ListOrdersUC: listOrdersUC}
}

func (s *OrderServiceServer) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.ListOrdersUC.Execute()
	if err != nil {
		return nil, err
	}
	resp := &pb.ListOrdersResponse{}
	for _, o := range orders {
		resp.Orders = append(resp.Orders, &pb.Order{
			Id:          o.ID,
			Description: o.Description,
			Price:       o.Price,
			CreatedAt:   o.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		})
	}
	return resp, nil
}
