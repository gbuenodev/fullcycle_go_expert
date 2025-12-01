package bid

import (
	"context"
	"fmt"
	"sync"

	"github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/logger"
	auctionentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/auction_entity"
	bidentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/bid_entity"
	"github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/infra/repository/auction"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type BidEntityMongo struct {
	ID        string  `bson:"id"`
	UserId    string  `bson:"user_id"`
	AuctionId string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func (br *BidRepository) Create(ctx context.Context, bids []bidentity.Bid) *internalerrors.InternalError {
	var wg sync.WaitGroup
	for _, bid := range bids {
		wg.Add(1)
		go func(bid bidentity.Bid) {
			defer wg.Done()

			auction, err := br.AuctionRepository.FindAuctionById(ctx, bid.AuctionId)
			if err != nil {
				msg := fmt.Sprintf("error trying to find auction with id: %s", bid.AuctionId)
				logger.Error(msg, err)
				return
			}

			if auction.Status != auctionentity.Active {
				msg := fmt.Sprintf("auction with id: %s is not active", bid.AuctionId)
				logger.Error(msg, err)
				return
			}

			bidEntityMongo := &BidEntityMongo{
				ID:        bid.ID,
				UserId:    bid.UserId,
				AuctionId: bid.AuctionId,
				Amount:    bid.Amount,
				Timestamp: bid.Timestamp.Unix(),
			}

			if _, err := br.Collection.InsertOne(ctx, bidEntityMongo); err != nil {
				msg := fmt.Sprintf("error trying to insert bid with id: %s", bid.ID)
				logger.Error(msg, err)
				return
			}
		}(bid)
	}

	wg.Wait()
	return nil
}
