package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Diaku49/grpc-game-server/config"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func InitRedis(cfg *config.Config) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPass,
		DB:       1,
	})
	return &RedisClient{client: client}
}

// Game Room

func (r *RedisClient) SetGameRoom(ctx context.Context) {

}

func (r *RedisClient) GetGameRooms(ctx context.Context) (map[string][2]string, error) {
	var results = make(map[string]*redis.StringCmd)
	roomIds, err := r.client.SMembers(ctx, GAME_ROOMS).Result()
	if err != nil {
		return nil, fmt.Errorf("couldnt fetch rooms")
	}

	// making pipeline redis for better performance
	rdbPipe := r.client.Pipeline()
	for _, id := range roomIds {
		key := GAME_ROOMS + id
		results[id] = rdbPipe.HGet(ctx, key, "players_id")
	}
	if _, err := rdbPipe.Exec(ctx); err != nil {
		return nil, fmt.Errorf("couldnt exec redis pipe")
	}

	// Run the commands
	gameRooms := make(map[string][2]string)
	for id, cmd := range results {
		playersId, err := cmd.Result()
		if err != nil {
			return nil, fmt.Errorf("couldnt fetch rooms")
		}

		var players_id []string
		err = json.Unmarshal([]byte(playersId), &playersId)
		if err != nil {
			return nil, fmt.Errorf("unmarshaling failed")
		}

		gameRooms[id] = [2]string(players_id)
	}

	return gameRooms, nil
}

func (r *RedisClient) GetGameRoomById(ctx context.Context, roomId string) {

}
