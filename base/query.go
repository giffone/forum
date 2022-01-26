package base

const (
	users         = "src_users"
	categories    = "src_categories"
	likes         = "src_likes"
	posts         = "posts"
	postslikes    = "posts_likes"
	comments      = "comments"
	commentslikes = "comments_likes"
	sessions      = "sessions"
)

type Query struct {
	Queries map[string]string
	src     map[string][]string
	tables  []string
}

func (q *Query) MakeQ(driver string) {
	q.src = make(map[string][]string)
	q.Queries = make(map[string]string)
	q.tables = []string{users, categories, likes, posts, postslikes, comments, commentslikes, sessions}

	switch driver {
	case "pq": // postgres driver
		q.pqQ()
	default: // sqlite3 by default
		q.sqliteQ()
	}
}

func (q *Query) sqliteQ() {
	q.Queries["attach"] = `ATTACH DATABASE ? AS ?;`

	q.Queries["detach"] = `DETACH DATABASE ?;`

	q.Queries["restore"] = `INSERT INTO %s SELECT * FROM %s.%s;`

	q.Queries["src_insert"] = `INSERT INTO %s (id, name)  
		VALUES (?,?)
	;`

	q.Queries[users] = `CREATE TABLE IF NOT EXISTS %s (
		"id"	INTEGER NOT NULL UNIQUE,
		"login"	TEXT NOT NULL UNIQUE,
		"password"	TEXT NOT NULL UNIQUE,
		"email"	TEXT NOT NULL UNIQUE,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	q.Queries[categories] = `CREATE TABLE IF NOT EXISTS %s (
		"id"	INTEGER NOT NULL UNIQUE,
		"name"	TEXT NOT NULL UNIQUE,
		PRIMARY KEY("id")
	);`

	q.Queries[likes] = `CREATE TABLE IF NOT EXISTS %s (
		"id"	INTEGER NOT NULL UNIQUE,
		"name"	TEXT NOT NULL UNIQUE,
		PRIMARY KEY("id")
	);`

	q.Queries[posts] = `CREATE TABLE IF NOT EXISTS %s (
		"id"	INTEGER NOT NULL UNIQUE,
		"date"	TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
		"user"	INTEGER NOT NULL,
		"category"	INTEGER,
		"post"	TEXT NOT NULL,
		FOREIGN KEY("user") REFERENCES "users"("id"),
		FOREIGN KEY("category") REFERENCES "categories"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	q.Queries[postslikes] = `CREATE TABLE IF NOT EXISTS %s (
		"id"	INTEGER NOT NULL UNIQUE,
		"value"	TEXT,
		"user"	INTEGER NOT NULL,
		"post"	INTEGER NOT NULL,
		FOREIGN KEY("user") REFERENCES "users"("id"),
		FOREIGN KEY("post") REFERENCES "posts"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	q.Queries[comments] = `CREATE TABLE  IF NOT EXISTS %s (
		"id"	INTEGER NOT NULL UNIQUE,
		"user"	INTEGER NOT NULL,
		"post"	INTEGER NOT NULL,
		"date"	TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
		"comment"	TEXT NOT NULL,
		FOREIGN KEY("user") REFERENCES "users"("id"),
		FOREIGN KEY("post") REFERENCES "posts"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	q.Queries[commentslikes] = `CREATE TABLE IF NOT EXISTS %s (
		"id"	INTEGER NOT NULL UNIQUE,
		"value"	INTEGER,
		"user"	INTEGER NOT NULL,
		"comment"	INTEGER NOT NULL,
		FOREIGN KEY("user") REFERENCES "users"("id"),
		FOREIGN KEY("comment") REFERENCES "comments"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`

	q.Queries[sessions] = `CREATE TABLE IF NOT EXISTS %s (
		"id"	INTEGER NOT NULL UNIQUE,
		"user"	TEXT NOT NULL UNIQUE,
		FOREIGN KEY("user") REFERENCES "users"("id"),
		PRIMARY KEY("id" AUTOINCREMENT)
	);`
}

func (q *Query) pqQ() {
}
