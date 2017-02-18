package main

import "github.com/Nekaerth/Strid/gui"
import "github.com/Nekaerth/Strid/game"
//import "github.com/Nekaerth/Strid/network"

func main() {
	game.SetUpGame()
	//state := game.SetUpGame()
	gui.SetUpInterface()
	//network.SetUpNetwork(state)
}
