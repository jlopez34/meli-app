package storage

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

func SetupStorage() (*Storage, error) {
	db, err := sql.Open("mysql", "root:federer1/meli_app")
	if err != nil {
		return nil, sql.ErrConnDone
	}

	return &Storage{db: db}, nil
}
