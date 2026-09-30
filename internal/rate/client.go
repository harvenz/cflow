package rate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
}

type priceResponse struct {
	Data struct {
		Amount   string `json:"amount"`
		Base     string `json:"base"`
		Currency string `json:"currency"`
	} `json:"data"`
}

func (c *Client) GetPrice(currency string) (string, error) {
	url := fmt.Sprintf(
		"https://api.coinbase.com/v2/prices/BTC-%s/spot",
		currency,
	)

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var data priceResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	return data.Data.Amount, nil
}