package conn

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type DbConnection struct {
	db *sql.DB
}

func (d DbConnection) GetConnection() *sql.DB {
	if d.db == nil {
		db, err := sql.Open("postgres", "host=localhost port=5432 dbname=mydb user=root password=root sslmode=disable")
		if err != nil {
			panic(err)
		}
		return db
	} else {
		return d.db
	}
}
