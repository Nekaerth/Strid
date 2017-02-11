package main

import "github.com/nequilich/gocto"
import "github.com/Nekaerth/Strid/gui"
import "github.com/Nekaerth/Strid/game"

func main() {
	gui.SetUpInterface()
	game.SetUpGame()
	SetUpInputs()
}