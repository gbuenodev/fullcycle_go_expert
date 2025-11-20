package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/entity"
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
	url := fmt.Sprintf("%s/%s/json/", c.baseURL, zipCode)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("viacep returned status %d", resp.StatusCode)
	}

	var viaCEPResp viaCEPReponse
	if err := json.NewDecoder(resp.Body).Decode(&viaCEPResp); err != nil {
		return nil, err
	}

	if viaCEPResp.Erro {
		return nil, fmt.Errorf("zipcode not found")
	}

	address := entity.NewAddress(viaCEPResp.Localidade)
	return address, nil
}
