package app

import (
	"context"
	"forum/internal/adapters/repository"
	"forum/internal/adapters/repository/mysqldb"
	"forum/internal/adapters/repository/sqlitedb"
	"log"
)

func switcher(ctx context.Context, driver string) repository.Repo {
	switch driver {
	case "mysql":
		return repository.NewRepoTx(ctx, &mysqldb.MySql{})
	case "sqlite3":
		return repository.NewRepo(ctx, &sqlitedb.Lite{})
	default:
		log.Fatalf("switcher: unknow driver %s\n", driver)
	}
	return nil
}
