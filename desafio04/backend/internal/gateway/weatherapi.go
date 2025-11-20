package gateway

import "context"

type WeatherAPIGateway interface {
	GetTemperatureByCity(ctx context.Context, city string) (float64, error)
}
