package game_server

import (
	"context"

	"github.com/Diaku49/grpc-game-server/internal/repositories/models"
	"github.com/Diaku49/grpc-game-server/pb"
	"github.com/Diaku49/grpc-game-server/pkg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (gs *GameServer) SignUpUser(ctx context.Context, req *pb.SignUpUserReq) (*pb.Message, error) {
	hashPass, err := pkg.Hash(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPass,
	}

	// I could use Transaction but i didnt :)
	// User Exists
	id, err := gs.db.FindUserIdByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if id != "" {
		return nil, status.Error(codes.AlreadyExists, "User with this email already exist")
	}

	// Creating User
	err = gs.db.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &pb.Message{
		Message: "Signed up successfully.",
	}, nil
}

func (gs *GameServer) LoginUser(ctx context.Context, req *pb.LoginUserReq) (*pb.LoginUserRes, error) {
	user, err := gs.db.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if isValidPass := pkg.Compare(user.Password, req.Password); isValidPass == false {
		return nil, status.Error(codes.Internal, "Password not correct.")
	}

	token, err := pkg.GenerateToken(user.Id, gs.cfg.JwtSecret)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.LoginUserRes{
		Id:         user.Id,
		Name:       user.Name,
		TotalWin:   user.Total_win,
		TotalGames: user.Total_games,
		Token:      token,
	}, nil
}

func (gs *GameServer) GetGameRooms(ctx context.Context, req *pb.GetGameRoomsReq) (*pb.GetGameRoomsRes, error) {
	gameRooms, err := gs.db.ListGameRooms(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	GameRooms := models.MapToPBGetGameRoomsRes(gameRooms)

	return &pb.GetGameRoomsRes{
		GameRooms: GameRooms,
	}, nil
}
