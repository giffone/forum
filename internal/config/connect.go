package config

type Conn struct {
	Name, Path, PathB, Driver, Port, Connection string
	Tx                                          bool
}

func (c *Conn) Connect(driver string) {
	if driver == "postgres" {
		c.Tx = true
		c.Name = "database-others.db"
		c.Path = "db/database-others.db"
		c.PathB = "db/backup/database-others.db"
		c.Driver = driver
		c.Port = ":3306"
		c.Connection = "user=admin password=admin " +
			"host=localhost port=3306 dbname= forum_db" +
			"connect_timeout=20 sslmode=disable" //user=username password=password host=dbname.host.com port=5433 dbname=dbname connect_timeout=20 sslmode=disable

	} else if driver == "mysql" {
		c.Tx = true
		c.Name = "database-mysql.db"
		c.Path = "db/database-mysql.db"
		c.PathB = "db/backup/database-mysql.db"
		c.Driver = driver
		c.Port = ":3306"
		c.Connection = "admin:admin@tcp(localhost:3306)/forum_db" //<username>:<pw>@tcp(<HOST>:<port>)/<dbname>
	} else {
		c.Name = "database-sqlite3.db"
		c.Path = "db/database-sqlite3.db"
		c.PathB = "db/backup/database-sqlite3.db"
		c.Driver = "sqlite3"
		c.Port = ":8080"
	}
}
