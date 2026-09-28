package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type PaymentRequest struct {
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
}

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /payments", createPayment)

	return mux
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	var req PaymentRequest

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(bytes, &req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Currency == "" {
		http.Error(w, "invalid currency", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 {
		http.Error(w, "invalid amount value", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}