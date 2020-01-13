package conn

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type DbConnection struct {
	db *sql.DB
}

func (d DbConnection) GetConnection() *sql.DB {
	if d.db == nil {
		db, err := sql.Open("postgres", "host=db port=5432 dbname=mydb user=root password=root sslmode=disable")
		if err != nil {
			log.Println(err)
		}
		return db
	} else {
		return d.db
	}
}
