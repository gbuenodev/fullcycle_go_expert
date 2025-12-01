package auction

import (
	"context"

	"github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/logger"
	auctionentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/auction_entity"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	ID          string                         `bson:"_id"`
	ProductName string                         `bson:"product_name"`
	Category    string                         `bson:"category"`
	Description string                         `bson:"description"`
	Condition   auctionentity.ProductCondition `bson:"condition"`
	Status      auctionentity.AuctionStatus    `bson:"status"`
	Timestamp   int64                          `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) Create(ctx context.Context, auction *auctionentity.Auction) *internalerrors.InternalError {
	auctionEntityMongo := AuctionEntityMongo{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   auction.Timestamp.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		msg := "error creating auction"
		logger.Error(msg, err)
		return internalerrors.NewInternalServerError(msg)
	}

	return nil
}
