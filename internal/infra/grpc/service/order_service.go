package service

import (
	"context"

	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/infra/grpc/pb"
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	ListOrderUseCase usecase.ListOrdersUseCase
}

func NewOrderService(listOrderUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		ListOrderUseCase: listOrderUseCase,
	}
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.ListOrdersResponse, error) {
	orders := []*pb.Order{}
	output, err := s.ListOrderUseCase.Execute()
	if err != nil {
		return nil, err
	}
	for _, o := range output {
		orders = append(orders, &pb.Order{
			Id:         o.ID,
			Price:      float32(o.Price),
			Tax:        float32(o.Tax),
			FinalPrice: float32(o.FinalPrice),
		})
	}
	return &pb.ListOrdersResponse{Orders: orders}, nil
}
