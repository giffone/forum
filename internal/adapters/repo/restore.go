package repo

import (
	"fmt"
	"forum/internal/config"
	"log"
	"os"
)

func (l *Lite) backup() {
	l.open()

	if l.attach() { // attach backup
		l.makeAndRestore("src") // make tables and restore data from backup
		l.detach("src")
	}
}

func (l *Lite) BackupExist() bool {
	_, err := os.Stat(l.Conn.PathB)
	return !os.IsNotExist(err)
}

func (l *Lite) attach() bool {
	log.Printf("Copying data from backup %s to %s\n", l.Conn.PathB, l.Conn.Path)
	src := "src"

	value, ok := l.Query.Queries[config.KeyAttach]
	if !ok {
		log.Println("query: can not find attach")
		return false
	}

	_, err := l.Db.Exec(value, l.Conn.PathB, src)
	if err != nil {
		log.Println("execute: can not attach backup")
		return false
	}
	log.Println("backup attached successfully")
	return true
}

func (l *Lite) detach(as string) {
	value, ok := l.Query.Queries[config.KeyDetach]
	if !ok {
		log.Println("query: can not find detach")
		return
	}
	_, err := l.Db.Exec(value, as)
	if err != nil {
		log.Println("execute: can not detach backup")
	}
}

func (l *Lite) makeAndRestore(src string) {
	for _, table := range l.Query.Tables {
		l.tables(table)
		l.restore(table, src)
	}
	log.Println("database copied")
}

func (l *Lite) restore(table, src string) {
	value, ok := l.Query.Queries[config.KeyRestore]
	if !ok {
		log.Println("query: can not find restore")
		return
	}

	val := fmt.Sprintf(value, table, src, table)
	result, err := l.Db.Exec(val)
	if err != nil {
		log.Printf("restore: \"%s\" was not restored, no such table in backup\n", table)
		return
	}
	numberLines, _ := result.RowsAffected()
	log.Printf("restore: %d lines added to \"%s\" table\n", numberLines, table)
}
