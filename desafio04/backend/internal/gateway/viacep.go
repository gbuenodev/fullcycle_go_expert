package gateway

import (
	"context"

	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/entity"
)

type ViaCEPGateway interface {
	GetAddressByZipCode(ctx context.Context, zipCode string) (*entity.Address, error)
}
