package config

const (
	TabUsers                 = "src_users"
	TabCategories            = "src_categories"
	TabLikes                 = "src_likes"
	TabPosts                 = "posts"
	TabPostsLikes            = "posts_likes"
	TabPostsCategories       = "posts_categories"
	TabComments              = "comments"
	TabCommentsLikes         = "comments_likes"
	TabSessions              = "sessions"
	KeyAttach                = "attach"
	KeyDetach                = "detach"
	KeyRestore               = "restore"
	KeyInsert2               = "insert_2"
	KeyInsert3               = "insert_3"
	KeyInsert4               = "insert_4"
	KeySelectUserBy          = "select_user_by"
	KeySelectSessionBy       = "select_session_by"
	KeySelectAllPosts        = "select_all_posts"
	KeySelectAllCategories   = "select_all_categories"
	KeySelectPostBy          = "select_post_byID"
	KeySelectPostsCategories = "select_posts_categories"
	KeyDeleteSessionBy       = "delete_session_by"
)

type Query struct {
	Queries map[string]string
	Tables  []string
}

func (q *Query) MakeQ(driver string) {
	q.Queries = make(map[string]string)
	q.Tables = []string{TabUsers, TabCategories, TabLikes, TabPosts, TabPostsLikes, TabComments, TabCommentsLikes, TabSessions}

	q.sqlQ()
	switch driver {
	case "others": // others driver
		q.posgresQ()
	case "mysql": // others driver
		q.mysqlQ()
	default: // lite3 by default
		q.sqliteQ()
	}
}

func (q *Query) sqlQ() { // for all databases
}

func (q *Query) sqliteQ() {
	q.Queries[KeyAttach] = `ATTACH DATABASE ? AS ?;`

	q.Queries[KeyDetach] = `DETACH DATABASE ?;`

	q.Queries[KeyRestore] = `INSERT INTO %s SELECT * FROM %s.%s;`

	q.Queries[KeyInsert2] = `INSERT INTO %s (%s, %s)  
		VALUES (?,?)
	;`
	q.Queries[KeyInsert3] = `INSERT INTO %s (%s, %s, %s)  
		VALUES (?,?,?)
	;`

	q.Queries[KeyInsert4] = `INSERT INTO %s (%s, %s, %s, %s)  
		VALUES (?,?,?,?)
	;`

	q.Queries[TabUsers] = `CREATE TABLE IF NOT EXISTS %s (
		"id"		INTEGER NOT NULL,
		"login"		TEXT NOT NULL UNIQUE,
		"password"	TEXT NOT NULL,
		"email"		TEXT NOT NULL UNIQUE,
		"root"		INTEGER NOT NULL DEFAULT 0,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	q.Queries[TabCategories] = `CREATE TABLE IF NOT EXISTS %s (
		"id"	INTEGER NOT NULL,
		"name"	TEXT NOT NULL UNIQUE,
		PRIMARY KEY("id")
	);`

	q.Queries[TabLikes] = `CREATE TABLE IF NOT EXISTS %s (
		"id"	INTEGER NOT NULL,
		"name"	TEXT NOT NULL UNIQUE,
		PRIMARY KEY("id")
	);`

	q.Queries[TabPosts] = `CREATE TABLE IF NOT EXISTS %s (
		"id"			INTEGER NOT NULL,
		"date"			TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
		"user"			INTEGER NOT NULL,
		"text"			TEXT NOT NULL,
		FOREIGN KEY("user") REFERENCES "src_users"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	q.Queries[TabPostsLikes] = `CREATE TABLE IF NOT EXISTS %s (
		"id"	INTEGER NOT NULL,
		"user"	INTEGER NOT NULL,
		"post"	INTEGER NOT NULL,
		"like"	INTEGER NOT NULL,
		FOREIGN KEY("user") REFERENCES "src_users"("id"),
		FOREIGN KEY("post") REFERENCES "posts"("id"),
		FOREIGN KEY("like") REFERENCES "src_likes"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	q.Queries[TabPostsCategories] = `CREATE TABLE IF NOT EXISTS %s (
		"id"		INTEGER NOT NULL UNIQUE,
		"post"		INTEGER NOT NULL,
		"category"	INTEGER NOT NULL,
		FOREIGN KEY("post") REFERENCES "posts"("id"),
		FOREIGN KEY("category") REFERENCES "src_categories"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	q.Queries[TabComments] = `CREATE TABLE  IF NOT EXISTS %s (
		"id"		INTEGER NOT NULL UNIQUE,
		"user"		INTEGER NOT NULL,
		"post"		INTEGER NOT NULL,
		"date"	  	TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
		"comment"	TEXT NOT NULL,
		FOREIGN KEY("user") REFERENCES "src_users"("id"),
		FOREIGN KEY("post") REFERENCES "posts"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	q.Queries[TabCommentsLikes] = `CREATE TABLE IF NOT EXISTS %s (
		"id"		INTEGER NOT NULL UNIQUE,
		"user"		INTEGER NOT NULL,
		"comment"	INTEGER NOT NULL,
		"like"	INTEGER NOT NULL,
		FOREIGN KEY("user") REFERENCES "src_users"("id"),
		FOREIGN KEY("comment") REFERENCES "comments"("id"),
		FOREIGN KEY("like") REFERENCES "src_likes"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	q.Queries[TabSessions] = `CREATE TABLE IF NOT EXISTS %s (
		"login"	INTEGER NOT NULL UNIQUE,
		"uuid"	TEXT NOT NULL UNIQUE,
		"date"	TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	q.Queries[KeySelectAllPosts] = `SELECT posts.id, posts.text, posts.user, src_users.login, posts.date
		FROM		posts 
		LEFT JOIN	src_users ON src_users.id = posts.user
		ORDER BY	posts.id DESC;`

	q.Queries[KeySelectAllCategories] = `SELECT src_categories.id, src_categories.name
		FROM		src_categories
		ORDER BY	src_categories.id ASC;`

	q.Queries[KeySelectPostBy] = `SELECT posts.id, posts.text, posts.user, src_users.login, posts.date
		FROM		posts
		LEFT JOIN	src_users ON src_users.id = posts.user
		WHERE 		posts.%s = ?;`

	q.Queries[KeySelectUserBy] = `SELECT src_users.id, src_users.login, src_users.password, src_users.email, src_users.root
		FROM	src_users
		WHERE	src_users.%s = ?;`

	q.Queries[KeySelectSessionBy] = `SELECT sessions.login, sessions.uuid, sessions.date
		FROM	sessions
		WHERE	sessions.%s = ?;`

	q.Queries[KeyDeleteSessionBy] = `DELETE
		FROM	sessions
		WHERE	sessions.%s = ?;`

}

func (q *Query) posgresQ() {
	q.Queries[KeyInsert2] = `INSERT INTO %s (%s, %s)  
		VALUES ($1,$2)
	;`
	q.Queries[KeyInsert3] = `INSERT INTO %s (%s, %s, %s)  
		VALUES ($1,$2,$3)
	;`

	q.Queries[KeyInsert4] = `INSERT INTO %s (%s, %s, %s, %s)  
		VALUES ($1,$2,$3,$4)
	;`
}

func (q *Query) mysqlQ() {
}
