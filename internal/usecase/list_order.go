package usecase

import (
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/dto"
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() ([]dto.OrderOutputDTO, error) {
	orders, err := l.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var ordersDTO []dto.OrderOutputDTO
	for _, order := range orders {
		ordersDTO = append(ordersDTO, dto.OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ordersDTO, nil
}
