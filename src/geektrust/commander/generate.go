package commander

import (
	"fmt"
	"geektrust/common"
	"strconv"
	"strings"
)

type InvestmentData struct {
	Equity float64
	Debt float64
	Gold float64
}

type CommandInfo struct {
	Name common.Command
	Data InvestmentData
	Month common.Month
}

func NewCommandInfo(command common.Command, data []string) *CommandInfo {
	if !VerifyDataSize(command, data) {
		fmt.Errorf("data size is not correct based on command: %v", command.String())
	}
	n := len(data) // n -> size of data
	investmentData := InvestmentData{} // update  according to command
	var month common.Month  // update according to command
	switch command {
	case common.ALLOCATE:
		investmentData = GenerateInvestmentData(data)
	case common.SIP:
		investmentData = GenerateInvestmentData(data)
	case common.CHANGE:
		investmentData = GenerateInvestmentData(data[:n-1])
		month = common.GetMonth(data[n-1])
	case common.BALANCE:
		month = common.GetMonth(data[0]) // balance command only has month
	case common.REBALACE:
		// do nothing
	}
	return &CommandInfo{command, investmentData, month}
}

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
			commands = append(commands, NewCommandInfo(common.ALLOCATE, commandData))
		case "SIP":
			commands = append(commands, NewCommandInfo(common.SIP, commandData))
		case "CHANGE":
			commands = append(commands, NewCommandInfo(common.CHANGE, commandData))
		case "BALANCE":
			commands = append(commands, NewCommandInfo(common.BALANCE, commandData))
		case "REBALANCE":
			commands = append(commands, NewCommandInfo(common.REBALACE, commandData))
		}
	}
	return commands
}

func GenerateInvestmentData(data []string) InvestmentData {
	// clean data
	for i, d := range data {
		data[i] = strings.ReplaceAll(d, "%", "")
	}
	equity, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		fmt.Errorf("error on parsing float (equity) %v", err)
	}
	debt, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		fmt.Errorf("error on parsing float (debt) %v", err)
	}
	gold, err := strconv.ParseFloat(data[2], 64)
	if err != nil {
		fmt.Errorf("error on parsing float (gold) %v", err)
	}
	return InvestmentData{equity, debt, gold}
}

func VerifyDataSize(command common.Command, data []string) bool {
	n := len(data)
	switch command {
	case common.ALLOCATE:
		return n == 3
	case common.SIP:
		return n == 3
	case common.CHANGE:
		return n == 4
	case common.BALANCE:
		return n == 1
	case common.REBALACE:
		return n == 0
	}
	return false
}
