package auctionentity

import (
	"context"
	"time"

	"github.com/gbuenodev/fullcycle_go_expert/desafio05/configs/logger"
	internalerrors "github.com/gbuenodev/fullcycle_go_expert/desafio05/internal/internal_errors"
	"github.com/google/uuid"
)

type Auction struct {
	ID          string
	ProductName string
	Category    string
	Description string
	Condition   ProductCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auction *Auction) *internalerrors.InternalError
	FindAuctionById(ctx context.Context, id string) (*Auction, *internalerrors.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]*Auction, *internalerrors.InternalError)
}

type ProductCondition int64
type AuctionStatus int64

const (
	Active AuctionStatus = iota
	Completed
)

const (
	New ProductCondition = iota
	Used
	Refurbished
)

func NewAuction(productName, category, description string, condition ProductCondition) (*Auction, *internalerrors.InternalError) {
	auction := &Auction{
		ID:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}

	if err := auction.Validate(); err != nil {
		msg := "error instantiating new auction: invalid fields"
		logger.Error(msg, err)
		return nil, internalerrors.NewBadRequestError(msg)
	}

	return auction, nil
}

func (a *Auction) Validate() *internalerrors.InternalError {
	if len(a.ProductName) == 0 {
		return internalerrors.NewBadRequestError("ProductName cannot be empty")
	}

	if len(a.Category) < 3 {
		return internalerrors.NewBadRequestError("Category must be at least 3 characters")
	}

	if len(a.Description) < 10 {
		return internalerrors.NewBadRequestError("Description must be at least 10 characters")
	}

	if len(a.Description) > 50 {
		return internalerrors.NewBadRequestError("Description must not exceed 50 characters")
	}

	if a.Condition < 0 || a.Condition > 2 {
		return internalerrors.NewBadRequestError("Condition must be New (0), Used (1), or Refurbished (2)")
	}

	if a.Status < 0 || a.Status > 1 {
		return internalerrors.NewBadRequestError("Status must be Active (0) or Completed (1)")
	}

	return nil
}
