package model

type Payment struct {
	Id       string `json:"id"`
	Currency string `json:"currency"`
	Amount   int    `json:"amount"`
	Satoshis int64  `json:"satoshis"`
	Address  string `json:"address"`
	Status   string `json:"status"`
	TXID     string `json:"txid"`
}