package category

import (
	"context"
	"database/sql"
	"forum/internal/adapters/repo/lib"
	"forum/internal/config"
	"forum/internal/constant"
	"forum/internal/model"
	"forum/internal/service"
)

type storageCategory struct {
	db      *sql.DB
	connect *config.Conn
	query   *config.Query
}

func NewStorage(db *sql.DB, cn *config.Conn, que *config.Query) service.StorageCategory {
	return &storageCategory{
		db:      db,
		connect: cn,
		query:   que,
	}
}

func (sc *storageCategory) GetAll(ctx context.Context) ([]*model.Category, error) {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	query := sc.query.Queries[config.KeySelectAllCategories]
	if !sc.connect.Tx {
		rows, err := lib.Query(ctx2, sc.db, query, nil)
		if err != nil {
			return nil, err // 500
		}
		defer lib.CloseRows(rows)

		return rowsCategory(rows)
	}
	tx, err := lib.TxBegin(ctx2, sc.db)
	if err != nil {
		return nil, err
	}
	defer lib.TxRollBack(tx)

	rows, err := lib.QueryTx(ctx2, tx, query, nil)
	if err != nil {
		return nil, err
	}

	defer lib.CloseRows(rows)

	resp, err := rowsCategory(rows)
	if err != nil {
		return nil, err
	}

	lib.TxCommit(tx)
	return resp, nil
}

func (sc *storageCategory) Create(ctx context.Context, dto *model.CreateCategoryDTO) error {
	return nil
}

func (sc *storageCategory) Delete(ctx context.Context, id int) error {
	return nil
}
