package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/logger"
	userentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/user_entity"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserEntityMongo struct {
	Id   string `bson:"_id"`
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

func (*UserRepository) Create(user *UserEntityMongo) error {
	return nil
}

func (ur *UserRepository) FindById(ctx context.Context, id string) (*userentity.User, *internalerrors.InternalError) {
	filter := bson.M{"_id": id}

	var user *UserEntityMongo
	err := ur.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			msg := fmt.Sprintf("could not find user with id: %s", id)
			logger.Error(msg, err)
			return nil, internalerrors.NewNotFoundError(msg)
		}
		msg := "error trying to find user"
		logger.Error(msg, err)
		return nil, internalerrors.NewInternalServerError(msg)
	}

	return &userentity.User{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}
