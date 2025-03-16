package web

import (
	"encoding/json"
	"net/http"

	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/dto"
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/entity"
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/usecase"
)

type WebOrderHandler struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewWebOrderHandler(
	orderRepository entity.OrderRepositoryInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		OrderRepository: orderRepository,
	}
}

func (h *WebOrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
	listOrdersUseCase := usecase.NewListOrdersUseCase(h.OrderRepository)
	output, err := listOrdersUseCase.Execute()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(&dto.ListOrdersOutputDTO{Orders: output})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
