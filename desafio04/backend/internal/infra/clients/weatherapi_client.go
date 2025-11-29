package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type WeatherAPIClient struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewWeatherAPIClient(baseURL, apiKey string) *WeatherAPIClient {
	return &WeatherAPIClient{
		baseURL:    baseURL,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

type weatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func (c *WeatherAPIClient) GetTemperatureByCity(ctx context.Context, city string) (float64, error) {
	tracer := otel.Tracer("weatherapi-client")
	ctx, span := tracer.Start(ctx, "WeatherAPI.GetTemperature")
	defer span.End()

	span.SetAttributes(
		attribute.String("city", city),
		attribute.String("api", "weatherapi"),
	)

	endpoint := fmt.Sprintf("%s/current.json", c.baseURL)

	params := url.Values{}
	params.Add("key", c.apiKey)
	params.Add("q", city)
	params.Add("aqi", "no")

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to create request")
		return 0, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "http request failed")
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("weatherapi returned status %d", resp.StatusCode)
		span.RecordError(err)
		span.SetStatus(codes.Error, "non-200 status code")
		span.SetAttributes(attribute.Int("http.status_code", resp.StatusCode))
		return 0, err
	}

	var weatherResp weatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to decode response")
		return 0, err
	}

	temp := weatherResp.Current.TempC
	span.SetAttributes(attribute.Float64("temperature_celsius", temp))
	span.SetStatus(codes.Ok, "success")
	return temp, nil
}
