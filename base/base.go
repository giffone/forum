package base

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

type DataBase struct {
	Name, Prefix, Driver, fileName, path, pathBackup string
	DB                                               *sql.DB
	Q                                                *Query
}

func (d *DataBase) Base() {
	d.fileName = fmt.Sprintf("%s-%s", d.Prefix, d.Name)

	if d.dbNotExist() { // if main base not exist
		d.create() // create new base

		if d.backupExist() { // if have backup
			d.backup()
			return
		}
	}

	d.open()
	d.make() // make tables without restore from backup
	insertSrc(d)
}

func (d *DataBase) dbNotExist() bool {
	d.path = fmt.Sprintf("db/%s.db", d.fileName)

	_, err := os.Stat(d.path)
	return os.IsNotExist(err)
}

func (d *DataBase) create() {
	log.Printf("Creating %s...\n", d.path)

	file, err := os.Create(d.path)
	if err != nil {
		log.Fatalf("create: can not create base - %v\n", err)
	}
	file.Close()
	log.Printf("%s created\n", d.path)
}

func (d *DataBase) open() {
	var err error
	d.DB, err = sql.Open(d.Driver, d.path)
	if err != nil {
		log.Fatalf("base: open: %v\n", err)
	}
}

func (d *DataBase) make() {
	for _, table := range d.Q.tables {
		d.tables(table)
	}
}

func (d *DataBase) tables(table string) {
	value, ok := d.Q.Queries[table]
	if !ok {
		log.Fatalf("%s was not created, no such query to make a table. Fatal exit\n", table)
	}

	val := fmt.Sprintf(value, table)
	_, err := d.DB.Exec(val)
	if err != nil {
		log.Fatalf("tables: execute %s: %v. Fatal exit\n", table, err) // if can not build table in db, stop programm
	}
}

// func (d *DataBase) Base() {
// 	d.fileName = fmt.Sprintf("%s-%s", d.Prefix, d.Name)
// 	restore := false

// 	if d.dbNotExist() { // if main base not exist
// 		d.create() // create new base

// 		if d.backupExist() { // if have backup
// 			restore = true
// 		}
// 	}

// 	d.open()

// 	if restore {
// 		if d.attach() { // attach backup
// 			d.makeAndRestore("src") // make tables and restore data from backup
// 			d.detach("src")
// 			return
// 		}
// 	}
// 	d.make() // make tables without restore from backup
// 	insertSrc(d)
// }
