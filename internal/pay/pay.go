package pay

import (
	"github.com/google/uuid"
	"github.com/harvenz/txd/internal/btc"
)

type Payment struct {
	Id       string
	Currency string
	Amount   int
	Address  string
	Status   string
	TXID     string
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