package repositories

import (
	"context"

	"github.com/Diaku49/grpc-game-server/internal/repositories/models"
	"github.com/jmoiron/sqlx"
)

type GameRepo interface {
	// Lobby
	CreateUser(ctx context.Context, user *models.User) error
	FindUserIdByEmail(ctx context.Context, email string) (string, error)
	FindUserByEmail(ctx context.Context, email string) (*models.User, error)
	// Game
	ListGameRooms(ctx context.Context) ([]models.GetGameRoomDTO, error)
	CreateGameRoom(ctx context.Context, user_id string) (string, error)
	CloseGameRoom(ctx context.Context, game_id, user_id string) error
	// Gameplay
}

type GameDB struct {
	db *sqlx.DB
}

func NewGameDB(db *sqlx.DB) *GameDB {
	return &GameDB{
		db: db,
	}
}
