package post

import (
	"context"
	"database/sql"
	"fmt"
	"forum/internal/adapters/repo/lib"
	"forum/internal/config"
	"forum/internal/constant"
	"forum/internal/model"
	"forum/internal/service"
)

type storagePost struct {
	db      *sql.DB
	connect *config.Conn
	query   *config.Query
}

func NewStorage(db *sql.DB, cn *config.Conn, que *config.Query) service.StoragePost {
	return &storagePost{
		db:      db,
		connect: cn,
		query:   que,
	}
}

func (sp *storagePost) GetAll(ctx context.Context) ([]*model.Post, error) {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	query := sp.query.Queries[config.KeySelectAllPosts]
	if !sp.connect.Tx {
		rows, err := lib.Query(ctx2, sp.db, query, nil)
		if err != nil {
			return nil, err // 500
		}
		defer lib.CloseRows(rows)

		return rowsPost(rows)
	}
	tx, err := lib.TxBegin(ctx2, sp.db)
	if err != nil {
		return nil, err
	}
	defer lib.TxRollBack(tx)

	rows, err := lib.QueryTx(ctx2, tx, query, nil)
	if err != nil {
		return nil, err
	}

	defer lib.CloseRows(rows)

	resp, err := rowsPost(rows)
	if err != nil {
		return nil, err
	}

	lib.TxCommit(tx)
	return resp, nil
}

func (sp *storagePost) GetByID(ctx context.Context, id int) (*model.Post, error) {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	query := fmt.Sprintf(sp.query.Queries[config.KeySelectPostBy], "id")

	if !sp.connect.Tx {
		row, err := lib.QueryRow(ctx2, sp.db, query, id)
		if err != nil {
			return nil, err
		}
		return rowPost(row)

	}

	tx, err := lib.TxBegin(ctx2, sp.db)
	if err != nil {
		return nil, err
	}
	defer lib.TxRollBack(tx)

	row, err := lib.QueryRowTx(ctx2, tx, query, id)
	if err != nil {
		return nil, err
	}

	resp, err := rowPost(row)
	if err != nil {
		return nil, err
	}

	lib.TxCommit(tx)
	return resp, nil
}

func (sp *storagePost) Create(ctx context.Context, post *model.CreatePostDTO) error {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	if !sp.connect.Tx {
		return sp.create(ctx2, post)
	}
	return sp.createTx(ctx2, post)
}

func (sp *storagePost) Delete(ctx context.Context, id int) error {
	//ctx2, cancel := context.WithTimeout(ctx, lib.TimeLimit)
	//defer cancel()

	//if !sp.connect.Tx {
	//	return sp.delete(ctx2, id)
	//}
	//return sp.deleteTx(ctx2, id)
	return nil
}
