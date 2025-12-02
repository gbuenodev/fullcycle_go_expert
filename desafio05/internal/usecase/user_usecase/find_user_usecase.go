package userusecase

import (
	"context"

	userentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/user_entity"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
)

type UserUseCase struct {
	UserRepository userentity.UserRepositoryInterface
}

type UserOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserUsecaseInterface interface {
	FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internalerrors.InternalError)
}

func NewUserUseCase(userRepository userentity.UserRepositoryInterface) UserUsecaseInterface {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

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
