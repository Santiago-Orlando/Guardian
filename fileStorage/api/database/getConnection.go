package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB) {

	connStr := os.Getenv("POSTGRESQL_URI")
	db, err := sql.Open("postgres", connStr + "Guardian_Files")
	if err != nil {
		panic(err)
	}
	return db
}
