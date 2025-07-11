package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"server/internal/store"
	"server/internal/utils"
	"time"
)

const EXCHANGE_API = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

type ExchangeHandler struct {
	exchangeStore store.ExchangeStore
}

func NewExchangeHandler(exchangeStore store.ExchangeStore) *ExchangeHandler {
	return &ExchangeHandler{
		exchangeStore: exchangeStore,
	}
}

func (eh *ExchangeHandler) HandleGetExchange(w http.ResponseWriter, r *http.Request) {
	reqCtx, cancelReq := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancelReq()

	req, err := http.NewRequestWithContext(reqCtx, "GET", EXCHANGE_API, nil)
	if err != nil {
		fmt.Printf("error preparing request: %s\n", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error requesting exchange: %s\n", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}
	defer resp.Body.Close()

	var wrapper map[string]json.RawMessage

	if err := json.NewDecoder(resp.Body).Decode(&wrapper); err != nil {
		fmt.Printf("error decoding JSON: %s\n", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	exchange := store.NewExchange()
	if err := json.Unmarshal(wrapper["USDBRL"], exchange); err != nil {
		fmt.Printf("error decoding bid: %s\n", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	saveCtx, cancelSave := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancelSave()
	bid, err := eh.exchangeStore.SaveExchange(saveCtx, exchange)
	if err != nil {
		fmt.Printf("error saving exchange: %s\n", err)
		utils.WriteJSON(w, http.StatusInternalServerError, utils.Envelope{"error": "internal server error"})
		return
	}

	utils.WriteJSON(w, http.StatusOK, utils.Envelope{"bid": bid})
}
