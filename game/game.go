package game

func SetUpGame() *GameState {
	state := new(GameState)
	state.Commands = make([]Command, 0)
	return state
}

type GameState struct {
	Commands []Command
}