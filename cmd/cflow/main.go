package main

import (
	"log"
	"net/http"

	"github.com/mlloc/cflow/internal/api"
	"github.com/mlloc/cflow/internal/btc"
	"github.com/mlloc/cflow/internal/env"
	"github.com/mlloc/cflow/internal/pay"
)

func main() {
	// load environment configuration
    cfg := env.Load()

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