package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	endpoint := fmt.Sprintf("%s/current.json", c.baseURL)

	params := url.Values{}
	params.Add("key", c.apiKey)
	params.Add("q", city)
	params.Add("aqi", "no")

	fullURL := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("weatherapi returned status %d", resp.StatusCode)
	}

	var weatherResp weatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return 0, err
	}

	return weatherResp.Current.TempC, nil
}
