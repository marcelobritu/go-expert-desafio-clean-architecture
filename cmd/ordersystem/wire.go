//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/entity"
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/infra/database"
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/infra/web"
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/usecase"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

func NewListOrdersUseCase(db *sql.DB) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
