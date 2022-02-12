package lib

import (
	"database/sql"
	"log"
)

func CloseRows(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		log.Printf("close rows: %v", err)
	}
}
