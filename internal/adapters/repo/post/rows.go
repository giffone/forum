package post

import (
	"database/sql"
	"forum/internal/constant"
	"forum/internal/model"
	"log"
)

func rowsPost(rows *sql.Rows) (resp []*model.Post, err error) {
	for rows.Next() {
		p := new(model.Post)
		if err := rows.Scan(&p.ID, &p.Text, &p.UserID, &p.User, &p.Date); err != nil {
			if err == sql.ErrNoRows {
				break
			}
			log.Printf("rowsPost: %v", err)
			return nil, constant.Err500
		}
		resp = append(resp, p)
	}
	if err := rows.Err(); err != nil {
		log.Printf("rowsPost: end with: %v", err)
		return nil, constant.Err500
	}
	if resp == nil {
		return nil, nil
	}
	return resp, nil
}

func rowPost(row *sql.Row) (*model.Post, error) {
	resp := new(model.Post)
	if err := row.Scan(&resp.ID, &resp.Text, &resp.UserID, &resp.User, &resp.Date); err != nil {
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			} else {
				log.Printf("rowPost: %v", err)
				return nil, constant.Err500
			}
		}
	}
	return resp, nil
}
