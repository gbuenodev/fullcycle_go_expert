package routes

import (
	"server/internal/app"

	"github.com/go-chi/chi/v5"
)

func Routes(app *app.App) *chi.Mux {
	r := chi.NewRouter()

	// HEALTH CHECK
	r.Get("/health", app.HealthCheck)

	// EXCHANGE ROUTES
	r.Get("/cotacao", app.ExchangeHandler.HandleGetExchange)

	return r
}
