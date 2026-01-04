package game_server

import (
	"context"

	"github.com/Diaku49/grpc-game-server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (gs *GameServer) CreateGameRoom(ctx context.Context, req *pb.CreateGameRoomReq) (*pb.Message, error) {
	game_id, err := gs.db.CreateGameRoom(ctx, req.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Message{
		Id:      game_id,
		Message: "Game room created successfully.",
	}, nil
}

func (gs *GameServer) CloseGameRoom(ctx context.Context, req *pb.CloseGameRoomReq) (*pb.Message, error) {
	err := gs.db.CloseGameRoom(ctx, req.GameId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Message{}, nil
}
