package repo

import (
	"database/sql"
	"fmt"
	"forum/internal/config"
	"log"
)

func InsertSrc(db *sql.DB, q *config.Query) {
	src := make(map[string][]string)
	src[config.TabLikes] = []string{"like", "dislike"}
	src[config.TabCategories] = []string{"анекдоты", "карикатуры", "мемы", "фразы", "стишки", "18+", "черный юмор"}

	que := fmt.Sprintf(q.Queries[config.KeyInsert4], config.TabUsers, "login", "password", "email", "root")
	_, err := db.Exec(que, "admin", "admin", "admin@mail.ru", 1) // root=1 for admin
	if err != nil {
		log.Printf("insert source: admin did not created: %v\n", err)
	}

	for table, source := range src {
		numberLines := 0
		que = fmt.Sprintf(q.Queries[config.KeyInsert2], table, "id", "name")
		for id, value := range source {
			_, err := db.Exec(que, id+1, value)
			if err != nil {
				log.Printf("insert source: \"%s\" did not inserted to \"%s\"\n", value, table)
				continue
			}
			numberLines++
		}
		log.Printf("insert source: %d lines added to \"%s\" table\n", numberLines, table)
	}
}
