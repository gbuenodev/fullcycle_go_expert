//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/gbuenodev/fullcycle_go_expert/desafio03/internal/entity"
	"github.com/gbuenodev/fullcycle_go_expert/desafio03/internal/event"
	"github.com/gbuenodev/fullcycle_go_expert/desafio03/internal/infra/database"
	"github.com/gbuenodev/fullcycle_go_expert/desafio03/internal/infra/web"
	"github.com/gbuenodev/fullcycle_go_expert/desafio03/internal/usecase"
	"github.com/gbuenodev/fullcycle_go_expert/desafio03/pkg/events"

	"github.com/google/wire"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

// Este é um injetor para o caso de uso de criação, usado pelo gRPC e outros.
func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewListOrdersUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.ListOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewListOrdersUseCase,
	)
	return &usecase.ListOrdersUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
		usecase.NewListOrdersUseCase,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
