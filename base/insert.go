package base

import (
	"fmt"
	"log"
)

func insertSrc(d *DataBase) {
	d.Q.src[categories] = []string{"анекдоты","карикатуры","мемы","фразы","стишки","18+","черный юмор"}
	d.Q.src[likes] = []string{"like", "dislike"}

	for table, source := range d.Q.src {
		numberLines := 0
		for id, value := range source {
			query := fmt.Sprintf(d.Q.Queries["src_insert"], table)
			_, err := d.DB.Exec(query, id+1, value)
			if err != nil {
				log.Printf("insert source: %s did not inserted to %s\n", value, table)
				continue
			}
			numberLines++
		}
		log.Printf("insert source: %d lines added to \"%s\" table\n", numberLines, table)
	}
}
