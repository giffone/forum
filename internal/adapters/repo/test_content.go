package repo

import (
	"database/sql"
	"fmt"
	"forum/internal/config"
)

func InsertTestUsers(db *sql.DB, q *config.Query) {
	que := fmt.Sprintf(q.Queries[config.KeyInsert3], config.TabUsers, "login", "password", "email")
	db.Exec(que, "user2", "123", "user2@mail.ru")
	db.Exec(que, "user3", "123", "user3@mail.ru")
	db.Exec(que, "user4", "123", "user4@mail.ru")
	db.Exec(que, "user5", "123", "user5@mail.ru")
}

func InsertTestPosts(db *sql.DB, q *config.Query) {
	que := fmt.Sprintf(q.Queries[config.KeyInsert2], config.TabPosts, "user", "text")
	db.Exec(que, "2", "Lorem Ipsum is simply dummy text of the printing and typesetting industry.")
	db.Exec(que, "3", "Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, "+
		"when an unknown printer took a galley of type and scrambled it to make a type specimen book.")
	db.Exec(que, "4", "It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.")
	db.Exec(que, "5", "It was popularised in the 1960s with the release of Letraset sheets containing"+
		" Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.")
}
