package webserver

import (
	"net/http"

	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/infra/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Router struct {
	WeatherHandler *api.WeatherHandler
}

func NewRouter(weatherHandler *api.WeatherHandler) *Router {
	return &Router{
		WeatherHandler: weatherHandler,
	}
}

func (rt *Router) SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// OTEL middleware
	r.Use(otelhttp.NewMiddleware("http-server"))

	// Health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	// Metrics exporter route
	r.Handle("/metrics", promhttp.Handler())

	r.Post("/weather", rt.WeatherHandler.PostWeatherByZipCode)

	return r
}
