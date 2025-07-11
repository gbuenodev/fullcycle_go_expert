package store

import (
	"context"

	"github.com/google/uuid"
)

type Exchange struct {
	Id              string `json:"id"`
	Currency        string `json:"currency"`
	DesiredCurrency string `json:"desiredCurrency"`
	Bid             string `json:"bid"`
}

func NewExchange() *Exchange {
	return &Exchange{
		Id:              uuid.New().String(),
		Currency:        "USD",
		DesiredCurrency: "BRL",
	}
}

type ExchangeStore interface {
	SaveExchange(ctx context.Context, exchange *Exchange) (string, error)
}
