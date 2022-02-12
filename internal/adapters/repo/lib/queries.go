package lib

import (
	"context"
	"database/sql"
	"forum/internal/constant"
	"log"
)

func Query(ctx context.Context, db *sql.DB, query string, key interface{}) (*sql.Rows, error) {
	rows, err := db.QueryContext(ctx, query, key)
	if err != nil {
		log.Printf("query \"%s\"\nended with error: %v", query, err)
		return nil, constant.Err500
	}
	return rows, nil
}

func QueryTx(ctx context.Context, tx *sql.Tx, query string, key interface{}) (*sql.Rows, error) {
	rows, err := tx.QueryContext(ctx, query, key)
	if err != nil {
		log.Printf("queryTx \"%s\"\nended with error: %v", query, err)
		return nil, constant.Err500
	}
	return rows, nil
}

func QueryRow(ctx context.Context, db *sql.DB, query string, key interface{}) (*sql.Row, error) {
	row := db.QueryRowContext(ctx, query, key)
	if err := row.Err(); err != nil {
		log.Printf("queryRow \"%s\"\nended with error: %v", query, err)
		return nil, constant.Err500
	}
	return row, nil
}

func QueryRowTx(ctx context.Context, tx *sql.Tx, query string, key interface{}) (*sql.Row, error) {
	row := tx.QueryRowContext(ctx, query, key)
	if err := row.Err(); err != nil {
		log.Printf("queryRowTx \"%s\"\nended with error: %v", query, err)
		return nil, constant.Err500
	}
	return row, nil
}
