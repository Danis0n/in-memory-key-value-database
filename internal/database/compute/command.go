package compute

const (
	UnknownCommandID = iota
	GetCommandID     = 1
	SetCommandID     = 2
	DelCommandID     = 3
)

var (
	UnknownCommand = "UNKNOWN"
	SetCommand     = "SET"
	GetCommand     = "GET"
	DelCommand     = "DEL"
)

var commandMap = map[string]int{
	SetCommand:     SetCommandID,
	GetCommand:     GetCommandID,
	DelCommand:     DelCommandID,
	UnknownCommand: UnknownCommandID,
}

func CommandFromName(command string) int {
	commandID, ok := commandMap[command]
	if !ok {
		return UnknownCommandID
	}

	return commandID
}
