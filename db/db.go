package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DBConn() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=gobiodata sslmode=disable password=222 host=localhost")
	if err != nil {
		return nil, err
	}

	return db, nil
}
