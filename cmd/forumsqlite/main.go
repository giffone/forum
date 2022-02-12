package main

import (
	"context"
	"database/sql"
	"forum/internal/app"
	_ "github.com/go-sql-driver/mysql" // import mysql library
	_ "github.com/lib/pq"              // import others library
	_ "github.com/mattn/go-sqlite3"    // Import go-lite3 library (by default)
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	db, router, port := app.NewApp(ctx).Run("sqlite3")

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("close database: %v\n", err)
		}
	}(db)

	log.Printf("localhost%s is listening...\n", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Printf("listening error: %v", err)
	}
}
