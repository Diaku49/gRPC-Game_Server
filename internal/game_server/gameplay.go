package game_server

import (
	"context"

	"github.com/Diaku49/grpc-game-server/pb"
	"google.golang.org/grpc"
)

func (gs *GameServer) JoinGame(ctx context.Context, req *pb.JoinGameReq) (*pb.JoinGameRes, error) {
	// playerId := req.GetPlayerId()
	// // freeRoom, err := gs.findRoom(playerId)
	// if err != nil {
	// 	return nil, fmt.Errorf("faild finding game room, error:%v", err)
	// }

	// if freeRoom != "" {
	// 	players, exists := gs.gameRooms[freeRoom]
	// 	if !exists {
	// 		return nil, fmt.Errorf("room %s not found", freeRoom)
	// 	}
	// 	if players[0] == "" {
	// 		players[0] = playerId
	// 	} else {
	// 		players[1] = playerId
	// 	}
	// 	gs.gameRooms[freeRoom] = players
	// }

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

// func (gs *GameServer) findRoom(id string) (string, error) {
// 	var roomIds []string
// 	// for key, players := range gs.gameRooms {
// 	// 	for _, pid := range players {
// 	// 		if pid == id {
// 	// 			return "", fmt.Errorf("this plater already joined a game room")
// 	// 		}
// 	// 	}
// 	// 	if len(players) < 2 {
// 	// 		roomIds = append(roomIds, key)
// 	// 	}
// 	// }

// 	if roomIds[0] == "" {
// 		return "", nil
// 	}

// 	return roomIds[0], nil
// }
