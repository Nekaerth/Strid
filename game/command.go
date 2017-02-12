package game

type Command struct {
	CommandStrings []string
}

func executeCommand(command Command) {
	if command.CommandStrings == nil {
		return
	}
	switch command.CommandStrings[0] {
	case "move":
		move(command)
	case "buy":
		buy(command)
	default:
		//TODO
	}
}

func move(command Command) {
	//TODO
}

func buy(command Command) {
	//TODO
}
