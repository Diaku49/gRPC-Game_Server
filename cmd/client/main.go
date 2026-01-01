package main

import (
	"fmt"

	"github.com/Diaku49/grpc-game-server/config"
)

func InitClient() {
	config, err := config.LoadConfig()
	if err != nil {
		panic("Retrieving config failed")
	}
	fmt.Print(config)
}
