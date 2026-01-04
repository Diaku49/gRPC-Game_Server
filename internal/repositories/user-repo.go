package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Diaku49/grpc-game-server/internal/repositories/models"
)

func (gr *GameDB) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (name, email, password)
	          VALUES (:name, :email, :password)`
	_, err := gr.db.NamedExecContext(ctx, query, user)
	if err != nil {
		return fmt.Errorf("create user failed, err: %v", err)
	}

	return nil
}

func (gr *GameDB) FindUserIdByEmail(ctx context.Context, email string) (string, error) {
	var id string
	query := `SELECT id FROM users WHERE email = $1`

	err := gr.db.GetContext(ctx, &id, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", fmt.Errorf("query failed, err: %v", err)
	}

	return id, nil
}

func (gr *GameDB) FindUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query :=
		`SELECT u.id, u.name, u.total_win, COUNT(g.id) AS total_games
	    FROM users u 
		LEFT JOIN games g ON u.id = g.player1_id OR u.id = g.player2_id
		WHERE email = $1
		GROUP BY u.id, u.name, u.total_win
	`

	err := gr.db.GetContext(ctx, &user, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("query failed, err: %v", err)
	}

	return &user, nil
}
