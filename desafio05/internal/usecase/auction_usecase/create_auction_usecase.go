package auctionusecase

import (
	"context"
	"time"

	auctionentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/auction_entity"
	bidentity "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/entity/bid_entity"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	bidusecase "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/usecase/bid_usecase"
)

type AuctionUseCase struct {
	AuctionRepository auctionentity.AuctionRepositoryInterface
	BidRepository     bidentity.BidRepositoryInterface
}

type AuctionInputDTO struct {
	ProductName string           `json:"product_name" binding:"required, min=1"`
	Category    string           `json:"category" binding:"required, min=2"`
	Description string           `json:"description" binding:"required, min=2, max=200"`
	Condition   ProductCondition `json:"condition" binding:"oneOf=1 2 3"`
}

type AuctionOutputDTO struct {
	ID          string           `json:"id"`
	ProductName string           `json:"product_name"`
	Category    string           `json:"category"`
	Description string           `json:"description"`
	Condition   ProductCondition `json:"condition"`
	Status      AuctionStatus    `json:"status"`
	Timestamp   time.Time        `json:"timestamp" time_format:"2006-01-02 15:04:05"`
}

type WinningInfoOutputDTO struct {
	Auction AuctionOutputDTO         `json:"auction"`
	Bid     *bidusecase.BidOutputDTO `json:"bid,omitempty"`
}

type ProductCondition int64
type AuctionStatus int64

type AuctionUseCaseInterface interface {
	CreateAuction(ctx context.Context, input *AuctionInputDTO) *internalerrors.InternalError
	FindAuctionById(ctx context.Context, id string) (*AuctionOutputDTO, *internalerrors.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]*AuctionOutputDTO, *internalerrors.InternalError)
	FindWinningBidByAuctionId(context.Context, string) (*WinningInfoOutputDTO, *internalerrors.InternalError)
}

func NewAuctionUseCase(auctionRepository auctionentity.AuctionRepositoryInterface, bidRepository bidentity.BidRepositoryInterface) AuctionUseCaseInterface {
	return &AuctionUseCase{
		AuctionRepository: auctionRepository,
		BidRepository:     bidRepository,
	}
}

func (au *AuctionUseCase) CreateAuction(ctx context.Context, auctionInput *AuctionInputDTO) *internalerrors.InternalError {
	auction, err := auctionentity.NewAuction(
		auctionInput.ProductName,
		auctionInput.Category,
		auctionInput.Description,
		auctionentity.ProductCondition(auctionInput.Condition),
	)
	if err != nil {
		return err
	}

	if err := au.AuctionRepository.CreateAuction(ctx, auction); err != nil {
		return err
	}

	return nil
}
