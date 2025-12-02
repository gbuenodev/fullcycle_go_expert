package userentity

import (
	"context"

	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	"github.com/google/uuid"
)

type User struct {
	ID   string
	Name string
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *User) *internalerrors.InternalError
	FindUserById(ctx context.Context, id string) (*User, *internalerrors.InternalError)
}

func NewUser(name string) (*User, *internalerrors.InternalError) {
	user := &User{
		ID:   uuid.New().String(),
		Name: name,
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() *internalerrors.InternalError {
	if len(u.Name) == 0 {
		return internalerrors.NewBadRequestError("Name cannot be empty")
	}

	return nil
}
