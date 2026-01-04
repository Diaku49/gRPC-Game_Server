package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Diaku49/grpc-game-server/config"
	"github.com/Diaku49/grpc-game-server/db"
	gs "github.com/Diaku49/grpc-game-server/internal/game_server"
	"github.com/Diaku49/grpc-game-server/internal/interceptors"
	"github.com/Diaku49/grpc-game-server/internal/redis"
	"github.com/Diaku49/grpc-game-server/internal/repositories"
	"github.com/Diaku49/grpc-game-server/pb"
	"google.golang.org/grpc"
)

func main() {
	// Getting config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Retrieving config failed error: %v", err)
	}

	InitServer(cfg)
}

func InitServer(cfg *config.Config) {
	// Shutdown setup
	errCh := make(chan error, 1)
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.Port))
	if err != nil {
		log.Fatalf("Could not start listening on port %s: %v", cfg.Port, err)
	}

	// Initializig dbs
	db, err := db.InitializeDb(cfg.DbUrl)
	if err != nil {
		log.Fatalf("Database initialization failed, error: %v", err)
	}
	fmt.Println("Game database initialized successfully")

	gameRp := repositories.NewGameDB(db)
	rdb := redis.InitRedis(cfg)

	//------------------- gRPC server
	var opts []grpc.ServerOption
	opts = []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptors.RegisterAuthInterceptor(cfg.JwtSecret)),
	}

	grpcServer := grpc.NewServer(opts...)
	// GameServer implementation
	gameServer, err := gs.NewGameServer(ctx, cfg, rdb, gameRp)
	if err != nil {
		log.Fatalf("server error: %s", err.Error())
	}
	// Registering service to gRPC
	pb.RegisterGameServerServer(grpcServer, gameServer)

	// gRPC listening
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			if !errors.Is(err, grpc.ErrServerStopped) {
				errCh <- err
			}
		}
	}()

	//------------- Gracefull shutdown process
	select {
	case <-ctx.Done():
		log.Println("Shutdown signal recieved.")
	case err = <-errCh:
		log.Printf("gRPC error recieved Error:%v", err)
	}

	done := make(chan struct{})
	go func() {
		grpcServer.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("Server stopped gracefully.")
	case <-time.After(10 * time.Second):
		grpcServer.Stop()
		log.Println("Server stopped with force.")
	}
}
