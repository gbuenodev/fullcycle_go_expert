package app

import (
	"database/sql"
	"net/http"
	"server/internal/api"
	"server/internal/store"
	"server/internal/utils"
	"server/migrations"
)

type App struct {
	ExchangeHandler *api.ExchangeHandler
	DBConn          *sql.DB
}

func NewApp() (*App, error) {
	dbConfig := store.DBConfig{
		Provider: "SQLite",
		Driver:   "sqlite3",
		DBPath:   "database/exchange.db",
	}

	DBConn, err := store.Open(&dbConfig)
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(DBConn, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// STORES
	exchangeStore := store.NewSQLiteExchangeStore(DBConn)

	// HANDLERS
	exchangeHandler := api.NewExchangeHandler(exchangeStore)

	app := &App{
		ExchangeHandler: exchangeHandler,
		DBConn:          DBConn,
	}

	return app, nil
}

func (a *App) HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"alive": "true"})
}
