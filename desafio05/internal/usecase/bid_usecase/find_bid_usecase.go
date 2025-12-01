package bidusecase

import (
	"context"

	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
)

func (bu *BidUseCase) FindBidByAuctionId(ctx context.Context, auctionId string) ([]*BidOutputDTO, *internalerrors.InternalError) {
	bids, err := bu.BidRepository.FindBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	var output []*BidOutputDTO
	for _, bid := range bids {
		output = append(output, &BidOutputDTO{
			ID:        bid.ID,
			UserID:    bid.UserId,
			AuctionID: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}

	return output, nil
}

func (bu *BidUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*BidOutputDTO, *internalerrors.InternalError) {
	bid, err := bu.BidRepository.FindWinningBidByAuctionId(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	return &BidOutputDTO{
		ID:        bid.ID,
		UserID:    bid.UserId,
		AuctionID: bid.AuctionId,
		Amount:    bid.Amount,
		Timestamp: bid.Timestamp,
	}, nil
}
