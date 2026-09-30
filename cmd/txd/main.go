package main

import (
	"log"
	"net/http"

	"github.com/harvenz/txd/internal/api"
	"github.com/harvenz/txd/internal/btc"
	"github.com/harvenz/txd/internal/env"
	"github.com/harvenz/txd/internal/pay"
	"github.com/harvenz/txd/internal/storage"
)

func main() {
	// load environment configuration
    cfg := env.Load()

    // prepare database
    storage.PrepareDB()

	// create bitcoin core client
    client, err := btc.NewClient(cfg)
    if err != nil {
        log.Fatal(err)
    }

	// create payment service
    service := pay.NewService(client)

	// create api router
    router := api.NewRouter(service)

    log.Println("starting api on :8080")

	// start http server
    err = http.ListenAndServe(":8080", router)
    if err != nil {
        log.Fatal(err)
    }
}