package game_server

import (
	"context"
	"fmt"

	"github.com/Diaku49/grpc-game-server/config"
	"github.com/Diaku49/grpc-game-server/db"
	"github.com/Diaku49/grpc-game-server/internals/redis"
	"github.com/Diaku49/grpc-game-server/pb"
)

type GameServer struct {
	pb.UnimplementedGameServerServer
	ctx       context.Context
	cfg       *config.Config
	rdb       *redis.RedisClient
	db        *db.GameDB
	gameRooms map[string][2]string
}

func NewGameServer(ctx context.Context, config *config.Config, rdb *redis.RedisClient, gameDB *db.GameDB) (*GameServer, error) {
	gameRooms, err := rdb.GetGameRooms(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load game rooms from Redis at startup: %w", err)
	}

	return &GameServer{
		cfg:       config,
		rdb:       rdb,
		db:        gameDB,
		gameRooms: gameRooms,
	}, nil
}
