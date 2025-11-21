//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/config"
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/gateway"
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/infra/api"
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/infra/clients"
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/infra/webserver"
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

func InitializeWebServer(cfg *config.Config) (*webserver.WebServer, error) {
	wire.Build(
		// Clients (implementações concretas)
		provideViaCEPClient,
		provideWeatherAPIClient,

		// Bind: liga implementações concretas às interfaces
		wire.Bind(new(gateway.ViaCEPGateway), new(*clients.ViaCEPClient)),
		wire.Bind(new(gateway.WeatherAPIGateway), new(*clients.WeatherAPIClient)),

		// Use Cases
		usecase.NewGetWeatherUseCase,

		// Handlers
		api.NewWeatherHandler,

		// Router
		webserver.NewRouter,
		provideChiRouter,

		// WebServer
		provideWebServer,
	)
	return nil, nil
}

func provideViaCEPClient(cfg *config.Config) *clients.ViaCEPClient {
	return clients.NewViaCEPClient(cfg.ViaCEP.BaseURL)
}

func provideWeatherAPIClient(cfg *config.Config) *clients.WeatherAPIClient {
	return clients.NewWeatherAPIClient(cfg.WeatherAPI.BaseURL, cfg.WeatherAPI.Key)
}

func provideChiRouter(router *webserver.Router) *chi.Mux {
	return router.SetupRoutes()
}

func provideWebServer(cfg *config.Config, chiRouter *chi.Mux) *webserver.WebServer {
	return webserver.NewWebServer(chiRouter, cfg.Server.Port)
}
