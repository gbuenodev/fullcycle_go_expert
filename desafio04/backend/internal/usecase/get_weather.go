package usecase

import (
	"context"
	"errors"
	"math"

	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/dto"
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/gateway"
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/pkg/validator"
)

var (
	ErrInvalidZipCode  = errors.New("invalid zipcode")
	ErrZipCodeNotFound = errors.New("can not find zipcode")
)

type GetWeatherUseCase struct {
	viaCEPGateway     gateway.ViaCEPGateway
	weatherAPIGateway gateway.WeatherAPIGateway
}

func NewGetWeatherUseCase(
	viaCepGateway gateway.ViaCEPGateway,
	weatherAPIGateway gateway.WeatherAPIGateway,
) *GetWeatherUseCase {
	return &GetWeatherUseCase{
		viaCEPGateway:     viaCepGateway,
		weatherAPIGateway: weatherAPIGateway,
	}
}

func (uc *GetWeatherUseCase) Execute(ctx context.Context, zipCode string) (*dto.WeatherResponseDTO, error) {
	if !validator.IsValidZipCode(zipCode) {
		return nil, ErrInvalidZipCode
	}

	address, err := uc.viaCEPGateway.GetAddressByZipCode(ctx, zipCode)
	if err != nil {
		return nil, ErrZipCodeNotFound
	}

	if address == nil || address.City == "" {
		return nil, ErrZipCodeNotFound
	}

	tempCelsius, err := uc.weatherAPIGateway.GetTemperatureByCity(ctx, address.City)
	if err != nil {
		return nil, err
	}

	// Compor: Address tem Weather
	address.SetWeather(tempCelsius)

	return &dto.WeatherResponseDTO{
		City:  address.City,
		TempC: round(address.Weather.Temperature.Celsius, 2),
		TempF: round(address.Weather.Temperature.Fahrenheit, 2),
		TempK: round(address.Weather.Temperature.Kelvin, 2),
	}, nil
}

// round arredonda um float64 para n casas decimais
func round(value float64, decimals int) float64 {
	multiplier := math.Pow(10, float64(decimals))
	return math.Round(value*multiplier) / multiplier
}
