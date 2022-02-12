package user

import (
	"database/sql"
	"forum/internal/constant"
	"forum/internal/model"
	"log"
)

func rowsUser(rows *sql.Rows) (resp []*model.User, err error) {
	for rows.Next() {
		p := new(model.User)
		if err := rows.Scan(&p.ID, &p.Login, &p.Password, &p.Email, &p.Root); err != nil {
			if err == sql.ErrNoRows {
				break
			}
			log.Printf("rowsUser: %v", err)
			return nil, constant.Err500
		}
		resp = append(resp, p)
	}
	if err := rows.Err(); err != nil {
		log.Printf("rowsUser: end with: %v", err)
		return nil, constant.Err500
	}
	if resp == nil {
		return nil, nil
	}
	return resp, nil
}

func rowUser(row *sql.Row) (*model.User, error) {
	resp := new(model.User)
	if err := row.Scan(&resp.ID, &resp.Login, &resp.Password, &resp.Email, &resp.Root); err != nil {
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			} else {
				log.Printf("rowUser: %v", err)
				return nil, constant.Err500
			}
		}
	}
	return resp, nil
}
