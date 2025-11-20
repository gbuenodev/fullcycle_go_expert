package api

import (
	"encoding/json"
	"net/http"

	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/dto"
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type WeatherHandler struct {
	GetWeatherUseCase *usecase.GetWeatherUseCase
}

func NewWeatherHandler(getWeatherUseCase *usecase.GetWeatherUseCase) *WeatherHandler {
	return &WeatherHandler{
		GetWeatherUseCase: getWeatherUseCase,
	}
}

func (h *WeatherHandler) GetWeatherByZipCode(w http.ResponseWriter, r *http.Request) {
	zipCode := chi.URLParam(r, "zipcode")

	res, err := h.GetWeatherUseCase.Execute(r.Context(), zipCode)
	if err != nil {
		h.handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *WeatherHandler) handleError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	var statusCode int
	var message string

	switch err {
	case usecase.ErrInvalidZipCode:
		statusCode = http.StatusUnprocessableEntity
		message = "invalid zipcode"
	case usecase.ErrZipCodeNotFound:
		statusCode = http.StatusNotFound
		message = "can not find zipcode"
	default:
		statusCode = http.StatusInternalServerError
		message = "internal server error"
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(dto.ErrorResponseDTO{Message: message})
}
