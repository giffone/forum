package main

import (
	"forum/conf"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

var confDriver = "sqlite3"

func main() {
	b := conf.Config(confDriver)
	b.Base()

	defer b.DB.Close()
}
