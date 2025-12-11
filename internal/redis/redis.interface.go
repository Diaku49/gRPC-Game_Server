package redis

import "context"

type Redis interface {
	GetUser(context.Context, string) (GetUser, error)
	SetUser(context.Context, CreateUserDTO) (string, error)
	SetGameRoom()
	GetGameRoomById()
	GetGameRooms()
}
