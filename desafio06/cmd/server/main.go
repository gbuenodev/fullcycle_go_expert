package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gbuenodev/fullcycle_go_expert/desafio06/config"
	"github.com/gbuenodev/fullcycle_go_expert/desafio06/internal/limiter"
	"github.com/gbuenodev/fullcycle_go_expert/desafio06/internal/middleware"
	"github.com/gbuenodev/fullcycle_go_expert/desafio06/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	store, err := storage.NewRedisStorage(cfg.RedisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	rateLimiter := limiter.NewRateLimiter(store, cfg)

	r := chi.NewRouter()

	// Wrap middleware to make it compatible with Chi
	r.Use(func(next http.Handler) http.Handler {
		return middleware.RateLimitMiddleware(rateLimiter, next)
	})

	r.Get("/auth", authHandler)
	r.Get("/", testHandler)

	addr := ":" + cfg.ServerPort
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	tier := r.URL.Query().Get("tier")
	if tier == "" {
		tier = "basic"
	}

	if tier != "basic" && tier != "premium" {
		http.Error(w, "Invalid tier. Use 'basic' or 'premium'", http.StatusBadRequest)
		return
	}

	token := fmt.Sprintf("%s:%s", tier, uuid.New().String())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"tier":  tier,
	})
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("OK\n"))
}
