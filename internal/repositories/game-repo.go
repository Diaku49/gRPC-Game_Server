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

func (gr *GameDB) CreateGameRoom(ctx context.Context, user_id string) (string, error) {
	var game_id string
	query := `
	INSERT INTO games (player1_id, status, rounds_num)
	VALUES ($1, $2, $3)
	RETURNING id
	`

	row := gr.db.QueryRowxContext(ctx, query, user_id, "open", 0)
	if err := row.Scan(&game_id); err != nil {
		return "", fmt.Errorf("Couldnt creat game room, err: %v", err)
	}

	return game_id, nil
}

func (gr *GameDB) CloseGameRoom(ctx context.Context, game_id, user_id string) error {
	query := `
	UPDATE games
	SET status = 'closed',
	updated_at = NOW()
	WHERE id = $1 AND (player1_id = $2 OR player_2 = $2)
	`

	result, err := gr.db.ExecContext(ctx, query, game_id, user_id)
	if err != nil {
		return fmt.Errorf("Failed close game room, err: %v", err)
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Failed affect game room row, err: %v", err)
	}
	if affect == 0 {
		return fmt.Errorf("game room with this id:%s not found", game_id)
	}

	return nil
}

func (gr *GameDB) UpdateGameRoom() {}

func (gr *GameDB) CreateRound() {}

func (gr *GameDB) UpdateRound() {}
