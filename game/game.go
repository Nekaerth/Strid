package game

func SetUpGame() *GameState {
	state := &GameState{}
	state.Commands = make([]Command, 0)
	return state
}

type GameState struct {
	Commands []Command
}