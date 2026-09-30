package storage

import (
	"database/sql"
	"log"

	"github.com/harvenz/txd/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

func SavePayment(payment *model.Payment) error {
	db, err := sql.Open("sqlite3", "./txd.db")
		if err != nil {
			log.Fatal(err)
		}

	defer db.Close()
	
	stmt := `
	INSERT INTO payments (
		id,
		currency,
		amount,
		satoshis,
		address,
		status,
		txid
	)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err = db.Exec(
		stmt,
		payment.Id,
		payment.Currency,
		payment.Amount,
		payment.Satoshis,
		payment.Address,
		payment.Status,
		payment.TXID,
	)

	if err != nil {
		log.Fatal(err)
	}

	return err
}