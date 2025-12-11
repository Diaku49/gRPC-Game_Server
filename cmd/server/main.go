package server

import (
	"fmt"
	"net"

	Config "github.com/Diaku49/grpc-game-server/internal/config"
	gs "github.com/Diaku49/grpc-game-server/internal/game_server"
	"github.com/Diaku49/grpc-game-server/internal/redis"
	"github.com/Diaku49/grpc-game-server/pb"
	"google.golang.org/grpc"
)

func InitServer() {
	// Getting config
	cfg, err := Config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("Retrieving config failed error: %v", err))
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", cfg.Port))
	if err != nil {
		panic(fmt.Sprintf("Could not start listening on port %s: %v", cfg.Port, err))
	}

	// Getting redis db
	rdb := redis.InitRedis(cfg)

	var opts []grpc.ServerOption

	// gRPC server
	grpcServer := grpc.NewServer(opts...)
	// GameServer implementation
	gameServer, err := gs.NewGameServer(cfg, rdb)
	if err != nil {
		panic(fmt.Sprintf("server error: %s", err.Error()))
	}
	// Registering service to gRPC
	pb.RegisterGameServerServer(grpcServer, gameServer)
	if err := grpcServer.Serve(lis); err != nil {
		panic(fmt.Sprintf("Failed to serve gRPC server: %s", err.Error()))
	}
}
