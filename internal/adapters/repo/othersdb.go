package repo

import (
	"context"
	"database/sql"
	"fmt"
	"forum/internal/config"
	"log"
)

type Others struct {
	Db    *sql.DB
	Conn  *config.Conn
	Query *config.Query
}

func (o *Others) NewClient(ctx context.Context) *sql.DB {
	o.open(ctx)
	o.make(ctx) // make tables without restore from backup
	InsertSrc(o.Db, o.Query)
	return o.Db
}

func (o *Others) open(ctx context.Context) {
	var err error
	o.Db, err = sql.Open(o.Conn.Driver, o.Conn.Connection)
	if err != nil {
		log.Fatalf("base: open: %v\n", err)
	}

	err = o.Db.PingContext(ctx)
	if err != nil {
		log.Fatalf("base: ping: %v\n", err)
	}
}

func (o *Others) make(ctx context.Context) {
	tx, err := o.Db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("function begin tx: %v", err)
	}

	for _, table := range o.Query.Tables {
		o.tables(tx, ctx, table)
	}

	if err = tx.Commit(); err != nil {
		log.Fatal(err)
	}
}

func (o *Others) tables(tx *sql.Tx, ctx context.Context, table string) {
	value, ok := o.Query.Queries[table]
	if !ok {
		log.Fatalf("%s was not created, no such query to make a table. Fatal exit\n", table)
	}

	val := fmt.Sprintf(value, table)

	_, err := tx.ExecContext(ctx, val)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Fatalf("tables: execute %s: unable to rollback: %v", table, rollbackErr)
		}
		log.Fatalf("tables: execute %s: %v. Fatal exit\n", table, err)
	}
}
