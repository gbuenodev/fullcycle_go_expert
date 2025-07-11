package store

import (
	"context"
	"database/sql"
)

type SQLiteExchangeStore struct {
	DBConn *sql.DB
}

func NewSQLiteExchangeStore(DBConn *sql.DB) *SQLiteExchangeStore {
	return &SQLiteExchangeStore{
		DBConn: DBConn,
	}
}

func (es *SQLiteExchangeStore) SaveExchange(ctx context.Context, exchange *Exchange) (string, error) {
	stmt, err := es.DBConn.Prepare("INSERT INTO exchanges(id, currency, desired_currency, bid) VALUES (?, ?, ?, ?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, exchange.Id, exchange.Currency, exchange.DesiredCurrency, exchange.Bid)
	if err != nil {
		return "", err
	}

	return exchange.Bid, nil
}
