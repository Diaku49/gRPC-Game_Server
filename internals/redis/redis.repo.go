package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Diaku49/grpc-game-server/config"
	"github.com/google/uuid"

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

// User

func (r *RedisClient) SetUser(ctx context.Context, data CreateUserDTO) (string, error) {
	if data.Email == "" || data.Name == "" {
		return "", fmt.Errorf("invalid input")
	}
	userId := uuid.New()
	id := userId.String()

	r.client.HSet(ctx, (USERS + id), &CreateUserDTO{
		Email: data.Email,
		Name:  data.Name,
	})

	return id, nil
}

func (r *RedisClient) GetUser(ctx context.Context, id string) (GetUser, error) {
	getUser, err := r.client.HGetAll(ctx, id).Result()
	if err != nil {
		return GetUser{}, fmt.Errorf("couldnt find user")
	}

	return GetUser{Email: getUser["email"], Name: getUser["name"]}, nil
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
