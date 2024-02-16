package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connect() *sql.DB {

	conn, err := sql.Open("postgres", "postgres://admin:123@localhost:5432/fiber")
	if err != nil {
		panic(err)
	}
	db = conn

	return db

}
