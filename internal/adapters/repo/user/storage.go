package user

import (
	"context"
	"database/sql"
	"fmt"
	"forum/internal/adapters/repo/lib"
	"forum/internal/config"
	"forum/internal/constant"
	"forum/internal/model"
	"forum/internal/service"
	"log"
)

type storageUser struct {
	db      *sql.DB
	connect *config.Conn
	query   *config.Query
}

func NewStorage(db *sql.DB, cn *config.Conn, que *config.Query) service.StorageUser {
	return &storageUser{
		db:      db,
		connect: cn,
		query:   que,
	}
}

func (su *storageUser) GetByID(ctx context.Context, id int) (*model.User, error) {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	query := fmt.Sprintf(su.query.Queries[config.KeySelectUserBy], "id")

	if !su.connect.Tx {
		row, err := lib.QueryRow(ctx2, su.db, query, id)
		if err != nil {
			return nil, err
		}

		return rowUser(row)
	}

	tx, err := lib.TxBegin(ctx2, su.db)
	if err != nil {
		return nil, err
	}
	defer lib.TxRollBack(tx)

	row, err := lib.QueryRowTx(ctx2, tx, query, id)
	if err != nil {
		return nil, err
	}

	resp, err := rowUser(row)
	if err != nil {
		return nil, err
	}

	lib.TxCommit(tx)
	return resp, nil
}

func (su *storageUser) Create(ctx context.Context, dto *model.CreateUserDTO) error {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	query := fmt.Sprintf(su.query.Queries[config.KeyInsert3], config.TabUsers, "login", "password", "email")
	_, err := su.db.ExecContext(ctx2, query, dto.Login, dto.Password, dto.Email)
	if err != nil {
		log.Printf("create user: %v\n", err)
		return err
	}
	return nil
}

func (su *storageUser) Delete(ctx context.Context, id int) error {
	return nil
}

func (su *storageUser) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	query := fmt.Sprintf(su.query.Queries[config.KeySelectUserBy], "login")

	if !su.connect.Tx {
		row, err := lib.QueryRow(ctx2, su.db, query, login)
		if err != nil {
			return nil, err
		}

		return rowUser(row)
	}

	tx, err := lib.TxBegin(ctx2, su.db)
	if err != nil {
		return nil, err
	}
	defer lib.TxRollBack(tx)

	row, err := lib.QueryRowTx(ctx2, tx, query, login)
	if err != nil {
		return nil, err
	}

	resp, err := rowUser(row)
	if err != nil {
		return nil, err
	}

	lib.TxCommit(tx)
	return resp, nil
}
