package client

import (
	"fmt"

	Config "github.com/Diaku49/grpc-game-server/internal/config"
)

func InitClient() {
	config, err := Config.LoadConfig()
	if err != nil {
		panic("Retrieving config failed")
	}
	fmt.Print(config)
}
