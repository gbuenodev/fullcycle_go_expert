package usecase

import (
	"context"

	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/dto"
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/internal/gateway"
	"github.com/gbuenodev/fullcycle_go_expert/desafio04/backend/pkg/validator"
)

type GetAddressUseCase struct {
	viaCEPGateway gateway.ViaCEPGateway
}

func NewGetAddressUseCase(viaCEPGateway gateway.ViaCEPGateway) *GetAddressUseCase {
	return &GetAddressUseCase{
		viaCEPGateway: viaCEPGateway,
	}
}

func (uc *GetAddressUseCase) Execute(ctx context.Context, zipCode string) (*dto.AddressResponseDTO, error) {
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

	return &dto.AddressResponseDTO{
		City: address.City,
	}, nil
}
