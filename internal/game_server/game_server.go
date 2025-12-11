package game_server

import (
	"context"
	"fmt"

	"github.com/Diaku49/grpc-game-server/internal/config"
	"github.com/Diaku49/grpc-game-server/internal/redis"
	"github.com/Diaku49/grpc-game-server/pb"
	"google.golang.org/grpc"
)

type GameServer struct {
	pb.UnimplementedGameServerServer
	cfg       *config.Config
	rdb       *redis.RedisClient
	gameRooms map[string][2]string
}

func NewGameServer(config *config.Config, rdb *redis.RedisClient) (*GameServer, error) {
	ctx := context.Background()
	gameRooms, err := rdb.GetGameRooms(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to load game rooms from Redis at startup: %w", err)
	}

	return &GameServer{
		cfg:       config,
		rdb:       rdb,
		gameRooms: gameRooms,
	}, nil
}

func (gs *GameServer) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	user := redis.CreateUserDTO{
		Email: req.GetEmail(),
		Name:  req.GetName(),
	}
	id, err := gs.rdb.SetUser(ctx, user)
	if err != nil {
		return &pb.CreateUserRes{}, fmt.Errorf("couldnt create user, error:%v", err)
	}

	message := "User created successfully."
	return &pb.CreateUserRes{Id: id, Message: message}, nil
}

func (gs *GameServer) JoinGame(ctx context.Context, req *pb.JoinGameReq) (*pb.JoinGameRes, error) {
	playerId := req.GetPlayerId()
	freeRoom, err := gs.findRoom(playerId)
	if err != nil {
		return nil, fmt.Errorf("faild finding game room, error:%v", err)
	}

	if freeRoom != "" {
		players, exists := gs.gameRooms[freeRoom]
		if !exists {
			return nil, fmt.Errorf("room %s not found", freeRoom)
		}
		if players[0] == "" {
			players[0] = playerId
		} else {
			players[1] = playerId
		}
		gs.gameRooms[freeRoom] = players
	}

	// creating a game room

	// update redis

	return nil, nil
}

func (gs *GameServer) StartGame(req *pb.StartGameReq, stream grpc.ServerStreamingServer[pb.StartGameRes]) error {

	return nil
}

func (gs *GameServer) MakeMove(ctx context.Context, req *pb.MakeMoveReq) (*pb.MakeMoveRes, error) {

	return nil, nil
}

func (gs *GameServer) findRoom(id string) (string, error) {
	var roomIds []string
	for key, players := range gs.gameRooms {
		for _, pid := range players {
			if pid == id {
				return "", fmt.Errorf("this plater already joined a game room")
			}
		}
		if len(players) < 2 {
			roomIds = append(roomIds, key)
		}
	}

	if roomIds[0] == "" {
		return "", nil
	}

	return roomIds[0], nil
}
