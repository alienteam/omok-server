package main

import (
	"github.com/alienteam/omok-server/core"
	"github.com/alienteam/omok-server/game"
)

func main() {
	g := &game.Game{}
	server := core.NewServer("localhost:8000", "/", g)
	server.Start()
}
