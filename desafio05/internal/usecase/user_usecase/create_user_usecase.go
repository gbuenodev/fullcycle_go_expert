package userusecase

import (
	"context"

	userentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/user_entity"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
)

type UserUseCase struct {
	UserRepository userentity.UserRepositoryInterface
}

type UserInputDTO struct {
	Name string `json:"name" binding:"required, min=2"`
}

type UserOutputDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserUsecaseInterface interface {
	CreateUser(ctx context.Context, input *UserInputDTO) *internalerrors.InternalError
	FindUserById(ctx context.Context, id string) (*UserOutputDTO, *internalerrors.InternalError)
}

func NewUserUseCase(userRepository userentity.UserRepositoryInterface) UserUsecaseInterface {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, input *UserInputDTO) *internalerrors.InternalError {
	user, err := userentity.NewUser(input.Name)
	if err != nil {
		return err
	}

	if err := uc.UserRepository.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}
