package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/entity"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type ViaCEPClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewViaCEPClient(baseURL string) *ViaCEPClient {
	return &ViaCEPClient{
		baseURL:    baseURL,
		httpClient: &http.Client{},
	}
}

type viaCEPReponse struct {
	Localidade string `json:"localidade"`
	Erro       bool   `json:"erro,omitempty"`
}

func (c *ViaCEPClient) GetAddressByZipCode(ctx context.Context, zipCode string) (*entity.Address, error) {
	tracer := otel.Tracer("viacep_client")
	ctx, span := tracer.Start(ctx, "ViaCEP.GetAddress")
	defer span.End()

	span.SetAttributes(
		attribute.String("zipcode", zipCode),
		attribute.String("api", "viacep"),
	)

	url := fmt.Sprintf("%s/%s/json/", c.baseURL, zipCode)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to create request")
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "http request failed")
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("viacep returned status %d", resp.StatusCode)
		span.RecordError(err)
		span.SetStatus(codes.Error, "non-200 status code")
		span.SetAttributes(attribute.Int("http.status_code", resp.StatusCode))
		return nil, err
	}

	var viaCEPResp viaCEPReponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResp); err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, "failed to decode response")
		return nil, err
	}

	if viaCEPResp.Erro {
		err := fmt.Errorf("zipcode not found")
		span.RecordError(err)
		span.SetStatus(codes.Error, "zipcode not found in viacep")
		return nil, err
	}

	address := entity.NewAddress(viaCEPResp.Localidade)
	span.SetAttributes(attribute.String("city", viaCEPResp.Localidade))
	span.SetStatus(codes.Ok, "success")
	return address, nil
}
