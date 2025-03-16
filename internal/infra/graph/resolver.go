package graph

import "github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ListOrdersUseCase usecase.ListOrdersUseCase
}
