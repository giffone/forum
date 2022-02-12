package main

import (
	"context"
	"database/sql"
	"forum/internal/app"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()

	db, router, port := app.NewApp(ctx).Run("postgres")

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
