package models

import (
	"time"

	"github.com/Diaku49/grpc-game-server/pb"
)

type GameRoom struct {
	Id       string  `db:"id"`
	Status   string  `db:"status"`
	RoundNum int32   `db:"rounds_num"`
	Player1  *string `db:"player1_id"`
	Player2  *string `db:"player2_id"`

	Created_at time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}

// ----------- DTO
type GetGameRoomDTO struct {
	Id          string `db:"id"`
	Status      string `db:"status"`
	RoundNum    int32  `db:"rounds_num"`
	Player1Name string `db:"player1_name"`
	Player2Name string `db:"player2_name"`
}

func MapToPBGetGameRoomsRes(ggrItems []GetGameRoomDTO) []*pb.GameRooms {
	var ggrRes []*pb.GameRooms
	for _, item := range ggrItems {
		newItem := &pb.GameRooms{
			Id:          item.Id,
			Status:      item.Status,
			RoundsNum:   item.RoundNum,
			Player1Name: item.Player1Name,
			Player2Name: item.Player2Name,
		}
		ggrRes = append(ggrRes, newItem)
	}

	return ggrRes
}
