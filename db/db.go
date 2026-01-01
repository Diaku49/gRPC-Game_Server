package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitializeDb(dbUrl string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("database initialization failed, err:%w", err)
	}

	return db, nil
}
