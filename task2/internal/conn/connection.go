package conn

import "database/sql"

type DbConnection struct {
	db *sql.DB
}

func (d DbConnection) GetConnection() *sql.DB {
	if d.db == nil {
		db, err := sql.Open("postgres", "host=db port=5432 dbname=mydb user=root password=root sslmode=disable")
		if err != nil {
			panic(err)
		}
		return db
	} else {
		return d.db
	}
}
