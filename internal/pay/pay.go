package pay

import (
	"log"
	"strconv"

	"github.com/google/uuid"
	"github.com/harvenz/txd/internal/btc"
	"github.com/harvenz/txd/internal/model"
	"github.com/harvenz/txd/internal/rate"
	"github.com/harvenz/txd/internal/storage"
)

type Service struct {
	btc *btc.Client
}

// create payment service with bitcoin client
func NewService(btcClient *btc.Client) *Service {
    return &Service{
        btc: btcClient,
    }
}

func (s *Service) CreatePayment(currency string, amount int) (*model.Payment, error) {
	addr, err := s.btc.GetNewAddress()
	if err != nil {
    	return nil, err
	}

	// get current bitcoin price
	rate := rate.Client{}
	price, err := rate.GetPrice(currency)
	if err != nil {
		log.Fatal(err)
	}

	// parse bitcoin price
	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
    	log.Fatal(err)
	}

	// calculate bitcoin amount
	btc := float64(amount) / priceFloat

	// convert bitcoin to satoshis
	satoshis := int64(btc * 100_000_000)


	payment := &model.Payment{
		Id: uuid.NewString(),
		Currency: currency,
		Amount: amount,
		Satoshis: satoshis,
		Address: addr,
		Status: "pending",
	}

	// save payment to storage
	storage.SavePayment(payment)

	return payment, err
}