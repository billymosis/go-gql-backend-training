package mydb

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", "./test.db")
	if err != nil {
		log.Fatalln(err)
	}
	return db, err
}
