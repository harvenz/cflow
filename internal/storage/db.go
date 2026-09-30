package storage

import (
	"database/sql"
	"log"

	"github.com/harvenz/txd/internal/model"
	_ "github.com/mattn/go-sqlite3"
)

func SavePayment(payment *model.Payment) {
	db, err := sql.Open("sqlite3", "./txd.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	
	stmt := `
    CREATE TABLE IF NOT EXISTS transactions (
        id INTEGER NOT NULL PRIMARY KEY,
        currency TEXT NOT NULL,
		amount INTEGER NOT NULL,
		satoshis INTEGER NOT NULL,
		address TEXT NOT NULL,
		status TEXT NOT NULL,
		txid TEXT
    );
    `
	_, err = db.Exec(stmt)
	if err != nil {
		log.Fatal(err)
	}
}