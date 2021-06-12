package commander

import (
	"geektrust/common"
	"strings"
)

func GenerateCommands(data []byte) []*CommandInfo {
	commands := make([]*CommandInfo, 0)
	content := string(data)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		command := tokens[0]
		commandData := tokens[1:]
		switch command {
		case "ALLOCATE":
			commands = append(commands, CreateCommandInfo(common.ALLOCATE, commandData))
		case "SIP":
			commands = append(commands, CreateCommandInfo(common.SIP, commandData))
		case "CHANGE":
			commands = append(commands, CreateCommandInfo(common.CHANGE, commandData))
		case "BALANCE":
			commands = append(commands, CreateCommandInfo(common.BALANCE, commandData))
		case "REBALANCE":
			commands = append(commands, CreateCommandInfo(common.REBALANCE, commandData))
		}
	}
	return commands
}