package webserver

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	router := chi.NewRouter()
	// O middleware é adicionado aqui, no momento da criação do router.
	router.Use(middleware.Logger)

	return &WebServer{
		Router:        router,
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handlerFunc http.HandlerFunc, method string) {
	switch strings.ToUpper(method) {
	case "GET":
		s.Router.Get(path, handlerFunc)
	case "POST":
		s.Router.Post(path, handlerFunc)
	case "PUT":
		s.Router.Put(path, handlerFunc)
	case "DELETE":
		s.Router.Delete(path, handlerFunc)
	default:
		s.Router.Handle(path, handlerFunc)
	}
}

// O método Start agora apenas inicia o servidor.
func (s *WebServer) Start() {
	http.ListenAndServe(s.WebServerPort, s.Router)
}
