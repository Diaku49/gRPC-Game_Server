package repositories

import (
	"context"
	"fmt"

	"github.com/Diaku49/grpc-game-server/internal/repositories/models"
)

func (gr *GameDB) ListGameRooms(ctx context.Context) ([]models.GetGameRoomDTO, error) {
	var gameRooms []models.GetGameRoomDTO
	query := `
	SELECT g.id, rounds_num, g.status,
	p1.name as player1_name,
	p2.name as player2_name
	FROM games g
	LEFT JOIN users p1 ON g.player1_id = p1.id
	LEFT JOIN users p2 ON g.player2_id = p2.id
	WHERE status = open OR 'status' = 'in_progress' `

	err := gr.db.SelectContext(ctx, &gameRooms, query, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to get gameRooms, err: %v", err)
	}

	return gameRooms, nil
}

func (gr *GameDB) CreateGameRoom() {}

func (gr *GameDB) UpdateGameRoom() {}

func (gr *GameDB) CreateRound() {}

func (gr *GameDB) UpdateRound() {}
