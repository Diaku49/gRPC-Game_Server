package redis

type Redis interface {
	SetGameRoom()
	GetGameRoomById()
	GetGameRooms()
}
