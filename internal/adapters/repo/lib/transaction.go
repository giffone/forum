package lib

import (
	"context"
	"database/sql"
	"forum/internal/constant"
	"log"
)

func TxBegin(ctx context.Context, db *sql.DB) (*sql.Tx, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("transaction begin: %v", err)
		return nil, constant.Err500
	}
	return tx, nil
}

func TxRollBack(tx *sql.Tx) {
	if err := tx.Rollback(); err != nil {
		log.Printf("transaction: can not rollback: %v", err)
	}
}

func TxCommit(tx *sql.Tx) {
	if err := tx.Commit(); err != nil {
		log.Printf("transaction commit: %v", err)
	}
}
