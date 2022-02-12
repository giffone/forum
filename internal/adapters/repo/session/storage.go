package session

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

type storageSession struct {
	db      *sql.DB
	connect *config.Conn
	query   *config.Query
}

func NewStorage(db *sql.DB, cn *config.Conn, que *config.Query) service.StorageSession {
	return &storageSession{
		db:      db,
		connect: cn,
		query:   que,
	}
}

func (ss *storageSession) GetByLogin(ctx context.Context, login string) (*model.Session, error) {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	query := fmt.Sprintf(ss.query.Queries[config.KeySelectSessionBy], "login")

	if !ss.connect.Tx {
		row, err := lib.QueryRow(ctx2, ss.db, query, login)
		if err != nil {
			return nil, err
		}

		return rowSession(row)
	}

	tx, err := lib.TxBegin(ctx2, ss.db)
	if err != nil {
		return nil, err
	}
	defer lib.TxRollBack(tx)

	row, err := lib.QueryRowTx(ctx2, tx, query, login)
	if err != nil {
		return nil, err
	}

	resp, err := rowSession(row)
	if err != nil {
		return nil, err
	}

	lib.TxCommit(tx)
	return resp, nil
}

func (ss *storageSession) Create(ctx context.Context, dto *model.CreateSessionDTO) error {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	query := fmt.Sprintf(ss.query.Queries[config.KeyInsert3], config.TabUsers, "login", "uuid", "date")
	_, err := ss.db.ExecContext(ctx2, query, dto.Login, dto.UUID, dto.Date)
	if err != nil {
		log.Printf("create user: %v\n", err)
		return err
	}
	return nil
}

func (ss *storageSession) DeleteByLogin(ctx context.Context, login string) error {
	ctx2, cancel := context.WithTimeout(ctx, constant.TimeLimit)
	defer cancel()

	query := fmt.Sprintf(ss.query.Queries[config.KeyDeleteSessionBy], "login")
	_, err := ss.db.ExecContext(ctx2, query, login)
	if err != nil {
		log.Printf("create user: %v\n", err)
		return err
	}
	return nil
}
