package auction

import (
	"context"
	"fmt"
	"time"

	"github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/logger"
	auctionentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/auction_entity"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ar *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*auctionentity.Auction, *internalerrors.InternalError) {
	filter := bson.M{"_id": id}

	var auction AuctionEntityMongo
	err := ar.Collection.FindOne(ctx, filter).Decode(&auction)
	if err != nil {
		msg := fmt.Sprintf("error finding auction with id: %s", id)
		logger.Error(msg, err)
		return nil, internalerrors.NewInternalServerError(msg)
	}

	return &auctionentity.Auction{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   auction.Condition,
		Status:      auction.Status,
		Timestamp:   time.Unix(auction.Timestamp, 0),
	}, nil
}

func (ar *AuctionRepository) FindAuctions(ctx context.Context, status auctionentity.AuctionStatus, category, productName string) ([]*auctionentity.Auction, *internalerrors.InternalError) {
	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["product_name"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	cursor, err := ar.Collection.Find(ctx, filter)
	if err != nil {
		msg := "error finding auctions"
		logger.Error(msg, err)
		return nil, internalerrors.NewInternalServerError(msg)
	}
	defer cursor.Close(ctx)

	var auctionEntityMongo []AuctionEntityMongo
	if err := cursor.All(ctx, &auctionEntityMongo); err != nil {
		msg := "error decoding auctions"
		logger.Error(msg, err)
		return nil, internalerrors.NewInternalServerError(msg)
	}

	var auctions []*auctionentity.Auction
	for _, auction := range auctionEntityMongo {
		auctions = append(auctions, &auctionentity.Auction{
			ID:          auction.ID,
			ProductName: auction.ProductName,
			Category:    auction.Category,
			Description: auction.Description,
			Condition:   auction.Condition,
			Status:      auction.Status,
			Timestamp:   time.Unix(auction.Timestamp, 0),
		})
	}

	return auctions, nil
}
