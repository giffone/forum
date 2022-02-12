package app

import (
	"context"
	"database/sql"
	"forum/internal/adapters/repo"
	"forum/internal/config"
)

type Switch interface {
	NewClient(ctx context.Context) *sql.DB
}

func start(ctx context.Context, s Switch) *sql.DB {
	return s.NewClient(ctx)
}

func Switcher(ctx context.Context, driver string) (*sql.DB, *config.Conn, *config.Query) {
	if driver == "mysql" || driver == "postgres" {
		o := &repo.Others{
			Conn:  new(config.Conn),
			Query: new(config.Query),
		}
		o.Conn.Connect(driver)
		o.Query.MakeQ(driver) // make queries
		return start(ctx, o), o.Conn, o.Query
	}

	l := &repo.Lite{
		Conn:  new(config.Conn),
		Query: new(config.Query),
	}
	l.Conn.Connect("sqlite3")
	l.Query.MakeQ(driver) // make queries
	return start(ctx, l), l.Conn, l.Query
}
