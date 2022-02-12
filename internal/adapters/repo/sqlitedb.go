package repo

import (
	"context"
	"database/sql"
	"fmt"
	"forum/internal/config"
	"log"
	"os"
)

type Lite struct {
	Db    *sql.DB
	Conn  *config.Conn
	Query *config.Query
}

func (l *Lite) NewClient(ctx context.Context) *sql.DB {
	if dbNotExist(l.Conn.Path) { // if main base not exist
		l.create() // create new base

		if l.BackupExist() { // have backup
			l.backup()
		}
	}
	l.open()
	l.make() // make tables without restore from backup
	InsertSrc(l.Db, l.Query)
	return l.Db
}

func (l *Lite) create() {
	log.Printf("Creating %s...\n", l.Conn.Name)

	file, err := os.Create(l.Conn.Path)
	if err != nil {
		log.Fatalf("create: can not create base - %v\n", err)
	}
	err = file.Close()
	if err != nil {
		return
	}

	log.Printf("%s created\n", l.Conn.Name)
}

func (l *Lite) open() {
	var err error

	l.Db, err = sql.Open(l.Conn.Driver, l.Conn.Path)
	if err != nil {
		log.Fatalf("base: open: %v\n", err)
	}
}

func (l *Lite) make() {
	for _, table := range l.Query.Tables {
		l.tables(table)
	}
}

func (l *Lite) tables(table string) {
	value, ok := l.Query.Queries[table]
	if !ok {
		log.Fatalf("%s was not created, no such query to make a table. Fatal exit\n", table)
	}

	val := fmt.Sprintf(value, table)
	_, err := l.Db.Exec(val)
	if err != nil {
		log.Fatalf("tables: execute %s: %v. Fatal exit\n", table, err) // if can not build table in Db, stop program
	}
}
