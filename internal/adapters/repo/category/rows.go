package category

import (
	"database/sql"
	"forum/internal/constant"
	"forum/internal/model"
	"log"
)

func rowsCategory(rows *sql.Rows) (resp []*model.Category, err error) {
	for rows.Next() {
		p := new(model.Category)
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			if err == sql.ErrNoRows {
				break
			}
			log.Printf("rowsCategory: %v", err)
			return nil, constant.Err500
		}
		resp = append(resp, p)
	}
	if err := rows.Err(); err != nil {
		log.Printf("rowsCategory: end with: %v", err)
		return nil, constant.Err500
	}
	if resp == nil {
		return nil, nil
	}
	return resp, nil
}

func rowCategory(row *sql.Row) (*model.Category, error) {
	resp := new(model.Category)
	if err := row.Scan(&resp.ID, &resp.Name); err != nil {
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			} else {
				log.Printf("rowCategory: %v", err)
				return nil, constant.Err500
			}
		}
	}
	return resp, nil
}
