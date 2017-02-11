package main

import "github.com/Nekaerth/Strid/gui"
import "github.com/Nekaerth/Strid/game"
import "github.com/Nekaerth/Strid/network"

func main() {
	state := game.SetUpGame()
	gui.SetUpInterface(state)
	network.SetUpNetwork(state)
}
