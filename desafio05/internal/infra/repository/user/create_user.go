package user

import (
	"context"

	"github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/logger"
	userentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/user_entity"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

type UserRepository struct {
	Collection *mongo.Collection
}

func NewUserRepository(database *mongo.Database) *UserRepository {
	return &UserRepository{
		Collection: database.Collection("users"),
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *userentity.User) *internalerrors.InternalError {
	userEntityMongo := UserEntityMongo{
		ID:   user.ID,
		Name: user.Name,
	}

	_, err := ur.Collection.InsertOne(ctx, userEntityMongo)
	if err != nil {
		msg := "error creating user"
		logger.Error(msg, err)
		return internalerrors.NewInternalServerError(msg)
	}

	return nil
}
