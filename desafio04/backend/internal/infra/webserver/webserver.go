package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebServer struct {
	Router *chi.Mux
	Port   string
}

func NewWebServer(router *chi.Mux, port string) *WebServer {
	return &WebServer{
		Router: router,
		Port:   port,
	}
}

func (s *WebServer) Start() error {
	addr := fmt.Sprintf(":%s", s.Port)
	log.Printf("Server running on port %s", s.Port)
	return http.ListenAndServe(addr, s.Router)
}
