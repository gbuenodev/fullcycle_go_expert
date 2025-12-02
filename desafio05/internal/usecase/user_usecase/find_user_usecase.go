package userusecase

import (
	"context"

	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
)

func (uc *UserUseCase) FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internalerrors.InternalError) {
	user, err := uc.UserRepository.FindUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserOutputDTO{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}
