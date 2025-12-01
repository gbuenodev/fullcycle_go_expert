package bid

import (
	"context"
	"fmt"
	"time"

	"github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/logger"
	bidentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/bid_entity"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (br *BidRepository) FindBidByAuctionId(ctx context.Context, auctionId string) ([]*bidentity.Bid, *internalerrors.InternalError) {
	filter := bson.M{"auction_id": auctionId}

	cursor, err := br.Collection.Find(ctx, filter)
	if err != nil {
		msg := fmt.Sprintf("error finding bids by auction id: %s", auctionId)
		logger.Error(msg, err)
		return nil, internalerrors.NewInternalServerError(msg)
	}
	defer cursor.Close(ctx)

	var bidEntityMongo []BidEntityMongo
	if err := cursor.All(ctx, &bidEntityMongo); err != nil {
		msg := fmt.Sprintf("error decoding bids by auction id: %s", auctionId)
		logger.Error(msg, err)
		return nil, internalerrors.NewInternalServerError(msg)
	}

	var bids []*bidentity.Bid
	for _, bid := range bidEntityMongo {
		bids = append(bids, &bidentity.Bid{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: time.Unix(bid.Timestamp, 0),
		})
	}

	return bids, nil
}

func (br *BidRepository) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*bidentity.Bid, *internalerrors.InternalError) {
	filter := bson.M{"auction_id": auctionId}
	opts := options.FindOne().SetSort(bson.D{{Key: "amount", Value: -1}})

	var bidEntityMongo BidEntityMongo
	if err := br.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo); err != nil {
		msg := fmt.Sprintf("error finding winning bid by auction id: %s", auctionId)
		logger.Error(msg, err)
		return nil, internalerrors.NewInternalServerError(msg)
	}

	return &bidentity.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil
}
