package session

import (
	"database/sql"
	"forum/internal/constant"
	"forum/internal/model"
	"log"
)

func rowSession(row *sql.Row) (*model.Session, error) {
	resp := new(model.Session)
	if err := row.Scan(&resp.Login, &resp.UUID, &resp.Date); err != nil {
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			} else {
				log.Printf("rowSession: %v", err)
				return nil, constant.Err500
			}
		}
	}
	return resp, nil
}
