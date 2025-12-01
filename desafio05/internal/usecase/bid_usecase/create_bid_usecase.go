package bidusecase

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/logger"
	bidentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/bid_entity"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
)

type BidUseCase struct {
	BidRepository bidentity.BidRepositoryInterface

	timer               *time.Timer
	maxBatchSize        int
	batchInsertInterval time.Duration
	bidChannel          chan bidentity.Bid
}

type BidInputDTO struct {
	UserID    string  `json:"user_id"`
	AuctionID string  `json:"auction_id"`
	Amount    float64 `json:"amount"`
}

type BidOutputDTO struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	AuctionID string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type BidUseCaseInterface interface {
	CreateBid(ctx context.Context, bidInput *BidInputDTO) (*BidOutputDTO, *internalerrors.InternalError)
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]*BidOutputDTO, *internalerrors.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internalerrors.InternalError)
}

var bidBatch []bidentity.Bid

func NewBidUseCase(bidRepository bidentity.BidRepositoryInterface) *BidUseCase {
	ctx := context.Background()
	maxSizeInterval := getMaxBatchSizeInterval()
	maxBatchSize := getMaxBatchSize()

	bid := &BidUseCase{
		BidRepository:       bidRepository,
		maxBatchSize:        maxBatchSize,
		batchInsertInterval: maxSizeInterval,
		timer:               time.NewTimer(maxSizeInterval),
		bidChannel:          make(chan bidentity.Bid, maxBatchSize),
	}

	go bid.triggerBatchInsert(ctx)

	return bid
}

func (bu *BidUseCase) CreateBid(ctx context.Context, bidInput *BidInputDTO) *internalerrors.InternalError {

	bid, err := bidentity.NewBid(bidInput.UserID, bidInput.AuctionID, bidInput.Amount)
	if err != nil {
		return err
	}

	bu.bidChannel <- *bid

	return nil

}

func getMaxBatchSizeInterval() time.Duration {
	batchInsertInterval := os.Getenv("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)
	if err != nil {
		return 3 * time.Minute
	}
	return duration
}

func getMaxBatchSize() int {
	maxBatchSize := os.Getenv("MAX_BATCH_SIZE")
	batchSize, err := strconv.Atoi(maxBatchSize)
	if err != nil {
		return 5
	}
	return batchSize
}

func (bu *BidUseCase) triggerBatchInsert(ctx context.Context) {
	defer close(bu.bidChannel)

	for {
		select {
		case bid, ok := <-bu.bidChannel:
			if !ok {
				if len(bidBatch) > 0 {
					if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
						msg := "error trying to process bid batch list"
						logger.Error(msg, err)
					}
				}
				return
			}
			bidBatch = append(bidBatch, bid)

			if len(bidBatch) >= bu.maxBatchSize {
				if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
					msg := "error trying to process bid batch list"
					logger.Error(msg, err)
				}
				bidBatch = nil
				bu.timer.Reset(bu.batchInsertInterval)
			}
		case <-bu.timer.C:
			if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
				msg := "error trying to process bid batch list"
				logger.Error(msg, err)
			}
			bidBatch = nil
			bu.timer.Reset(bu.batchInsertInterval)
		}
	}
}
