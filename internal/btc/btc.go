package btc

import (
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/harvenz/txd/internal/env"
)

type Client struct {
	rpc *rpcclient.Client
}

func NewClient(env env.Env) (*Client, error) {
	config := &rpcclient.ConnConfig{
		Host: env.BTC.Host,
		User: env.BTC.User,
		Pass: env.BTC.Pass,
		HTTPPostMode: true,
		DisableTLS: true,
	}

	rpc, err := rpcclient.New(config, nil)
	if err != nil {
		return nil, err
	}

	return &Client{
		rpc: rpc,
	}, nil
}

func (c *Client) GetBlockchainInfo() (*btcjson.GetBlockChainInfoResult, error) {
	return c.rpc.GetBlockChainInfo()
}

func (c *Client) GetNewAddress() (string, error) {
	address, err := c.rpc.GetNewAddress("")
	if err != nil {
		return "", err
	}

	return address.EncodeAddress(), nil
}