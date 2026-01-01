package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type GameDB struct {
	db *sqlx.DB
}

func InitializeDb(dbUrl string) (*GameDB, error) {
	db, err := sqlx.Connect("postgres", dbUrl)
	if err != nil {
		return nil, fmt.Errorf("database initialization failed, err:%w", err)
	}

	return &GameDB{
		db: db,
	}, nil
}
