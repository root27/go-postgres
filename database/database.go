package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() (*sql.DB, error) {

	conn, err := sql.Open("postgres", "postgres://admin:123@localhost:5432/fiber")
	if err != nil {

		log.Fatal(err)

		return nil, err
	}
	db = conn

	return db, nil

}
