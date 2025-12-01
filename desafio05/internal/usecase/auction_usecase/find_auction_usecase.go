package auctionusecase

import (
	"context"

	auctionentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/auction_entity"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	bidusecase "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/usecase/bid_usecase"
)

func (au *AuctionUseCase) FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internalerrors.InternalError) {
	auction, err := au.AuctionRepository.FindAuctionById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &AuctionOutputDTO{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}, nil
}

func (au *AuctionUseCase) FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]*AuctionOutputDTO, *internalerrors.InternalError) {
	auctions, err := au.AuctionRepository.FindAuctions(
		ctx,
		auctionentity.AuctionStatus(status),
		category,
		productName,
	)
	if err != nil {
		return nil, err
	}

	var output []*AuctionOutputDTO
	for _, auction := range auctions {
		output = append(output, &AuctionOutputDTO{
			ID:          auction.ID,
			ProductName: auction.ProductName,
			Category:    auction.Category,
			Description: auction.Description,
			Condition:   ProductCondition(auction.Condition),
			Status:      AuctionStatus(auction.Status),
			Timestamp:   auction.Timestamp,
		})
	}

	return output, nil
}

func (au *AuctionUseCase) FindWinningBidByAuctionId(ctx context.Context, auctionId string) (*WinningInfoOutputDTO, *internalerrors.InternalError) {
	auction, err := au.AuctionRepository.FindAuctionById(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	auctionOutput := AuctionOutputDTO{
		ID:          auction.ID,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}

	bid, err := au.BidRepository.FindWinningBidByAuctionId(ctx, auction.ID)
	if err != nil {
		return &WinningInfoOutputDTO{
			Auction: auctionOutput,
			Bid:     nil,
		}, nil
	}

	bidOutput := &bidusecase.BidOutputDTO{
		ID:        bid.ID,
		UserID:    bid.UserId,
		AuctionID: bid.AuctionId,
		Amount:    bid.Amount,
		Timestamp: bid.Timestamp,
	}

	return &WinningInfoOutputDTO{
		Auction: auctionOutput,
		Bid:     bidOutput,
	}, nil
}
