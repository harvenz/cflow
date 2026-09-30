package pay

import (
	"github.com/google/uuid"
	"github.com/harvenz/txd/internal/btc"
)

type Payment struct {
    Id       string `json:"id"`
    Currency string `json:"currency"`
    Amount   int    `json:"amount"`
    Address  string `json:"address"`
    Status   string `json:"status"`
    TXID     string `json:"txid"`
}

type Service struct {
	btc *btc.Client
}

// create payment service with bitcoin client
func NewService(btcClient *btc.Client) *Service {
    return &Service{
        btc: btcClient,
    }
}

func (s *Service) CreatePayment(currency string, amount int) (*Payment, error) {
	addr, err := s.btc.GetNewAddress()
	if err != nil {
    	return nil, err
	}

	payment := &Payment{
		Id: uuid.NewString(),
		Currency: currency,
		Amount: amount,
		Address: addr,
		Status: "pending",
	}

	return payment, err
}