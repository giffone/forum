package conf

import "forum/base"

func Config(driver string) *base.DataBase {
	switch driver {
	case "pq": // postgres driver
		b := &base.DataBase{
			Name:   "postgres",
			Prefix: "database",
			Driver: "pq",
			Q:      new(base.Query),
		}
		b.Q.MakeQ("pq")
		return b
	default: // sqlite3 by default
		b := &base.DataBase{
			Name:   "sqlite",
			Prefix: "database",
			Driver: "sqlite3",
			Q:      new(base.Query),
		}
		b.Q.MakeQ("sqlite3") // make queries
		return b
	}
	
}
