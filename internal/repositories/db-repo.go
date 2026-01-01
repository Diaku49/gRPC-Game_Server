package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type GameRepo interface {
	CreateUser(ctx context.Context)
}

type GameDB struct {
	db *sqlx.DB
}

func NewGameDB(db *sqlx.DB) *GameDB {
	return &GameDB{
		db: db,
	}
}
