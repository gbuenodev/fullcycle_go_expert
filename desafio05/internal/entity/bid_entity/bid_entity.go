package bidentity

import (
	"context"
	"time"

	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	"github.com/google/uuid"
)

type Bid struct {
	ID        string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

type BidRepositoryInterface interface {
	CreateBid(ctx context.Context, bids []Bid) *internalerrors.InternalError
	FindBidByAuctionId(ctx context.Context, auctionId string) ([]*Bid, *internalerrors.InternalError)
	FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*Bid, *internalerrors.InternalError)
}

func NewBid(userId, auctionId string, amount float64) (*Bid, *internalerrors.InternalError) {
	bid := &Bid{
		ID:        uuid.New().String(),
		UserId:    userId,
		AuctionId: auctionId,
		Amount:    amount,
		Timestamp: time.Now(),
	}

	if err := bid.Validate(); err != nil {
		return nil, err
	}

	return bid, nil
}

func (b *Bid) Validate() *internalerrors.InternalError {
	if err := uuid.Validate(b.ID); err != nil {
		return internalerrors.NewBadRequestError("UserID is not a valid ID")
	}

	if err := uuid.Validate(b.AuctionId); err != nil {
		return internalerrors.NewBadRequestError("AuctionID is not a valid ID")
	}

	if b.Amount <= 0 {
		return internalerrors.NewBadRequestError("Amount must be greater than zero")
	}

	return nil
}
