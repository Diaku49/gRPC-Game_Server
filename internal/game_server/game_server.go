package game_server

import (
	"context"

	"github.com/Diaku49/grpc-game-server/config"
	"github.com/Diaku49/grpc-game-server/internal/redis"
	"github.com/Diaku49/grpc-game-server/internal/repositories"
	"github.com/Diaku49/grpc-game-server/pb"
)

type GameServer struct {
	pb.UnimplementedGameServerServer
	ctx context.Context
	cfg *config.Config
	rdb *redis.RedisClient
	db  repositories.GameRepo
}

func NewGameServer(ctx context.Context, config *config.Config, rdb *redis.RedisClient, db repositories.GameRepo) (*GameServer, error) {
	// gameRooms, err := rdb.GetGameRooms(ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to load game rooms from Redis at startup: %w", err)
	// }

	return &GameServer{
		cfg: config,
		ctx: ctx,
		rdb: rdb,
		db:  db,
	}, nil
}
