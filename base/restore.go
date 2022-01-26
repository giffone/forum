package base

import (
	"fmt"
	"log"
	"os"
)

func (d *DataBase) backup() {
	d.open()

	if d.attach() { // attach backup
		d.makeAndRestore("src") // make tables and restore data from backup
		d.detach("src")
	}
}

func (d *DataBase) backupExist() bool {
	d.pathBackup = fmt.Sprintf("db/backup/%s.db", d.fileName)

	_, err := os.Stat(d.pathBackup)
	return !os.IsNotExist(err)
}

func (d *DataBase) attach() bool {
	log.Printf("Copying data from backup %s to %s\n", d.pathBackup, d.path)
	src := "src"

	value, ok := d.Q.Queries["attach"]
	if !ok {
		log.Println("query: can not find attach")
		return false
	}

	_, err := d.DB.Exec(value, d.pathBackup, src)
	if err != nil {
		log.Println("execute: can not attach backup")
		return false
	}
	log.Println("backup attached successfully")
	return true
}

func (d *DataBase) detach(as string) {
	value, ok := d.Q.Queries["detach"]
	if !ok {
		log.Println("query: can not find detach")
		return
	}
	_, err := d.DB.Exec(value, as)
	if err != nil {
		log.Println("execute: can not detach backup")
	}
}

func (d *DataBase) makeAndRestore(src string) {
	for _, table := range d.Q.tables {
		d.tables(table)
		d.restore(table, src)
	}
	log.Println("database copied")
}

func (d *DataBase) restore(table, src string) {
	value, ok := d.Q.Queries["restore"]
	if !ok {
		log.Println("query: can not find restore")
		return
	}

	val := fmt.Sprintf(value, table, src, table)
	result, err := d.DB.Exec(val)
	if err != nil {
		log.Printf("restore: \"%s\" was not restored, no such table in backup\n", table)
		return
	}
	numberLines, _ := result.RowsAffected()
	log.Printf("restore: %d lines added to \"%s\" table\n", numberLines, table)
}
