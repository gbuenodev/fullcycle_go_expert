package auction

import (
	"context"
	"os"
	"time"

	"github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/logger"
	auctionentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/auction_entity"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	"go.mongodb.org/mongo-driver/bson"
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

func (ar *AuctionRepository) CreateAuction(ctx context.Context, auction *auctionentity.Auction) *internalerrors.InternalError {
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

	go ar.completeAuction(context.Background(), auction.ID)

	return nil
}

func getAuctionDuration() time.Duration {
	auctionDuration := os.Getenv("AUCTION_DURATION")
	duration, err := time.ParseDuration(auctionDuration)
	if err != nil {
		return 24 * time.Hour
	}
	return duration
}

func (ar *AuctionRepository) completeAuction(ctx context.Context, auctionID string) {
	duration := getAuctionDuration()
	timeoutCtx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()

	<-timeoutCtx.Done()

	if timeoutCtx.Err() == context.DeadlineExceeded {
		update := bson.M{"$set": bson.M{"status": auctionentity.Completed}}
		filter := bson.M{"_id": auctionID}

		if _, err := ar.Collection.UpdateOne(ctx, filter, update); err != nil {
			logger.Error("error completing auction", err)
		}
	}
}
