package commander

import (
	"fmt"
	"geektrust/common"
)

type CommandInfo struct {
	name common.Command
	data common.InvestmentData
	month common.Month
}

func NewCommandInfo(name common.Command, data common.InvestmentData, month common.Month) *CommandInfo {
	return &CommandInfo{name: name, data: data, month: month}
}

func CreateCommandInfo(command common.Command, data []string) *CommandInfo {
	if !common.VerifyDataSize(command, data) {
		fmt.Errorf("data size is not correct based on command: %v", command.String())
	}
	n := len(data) // n -> size of data
	var investmentData common.InvestmentData // update  according to command
	var month common.Month  				// update according to command
	switch command {
	case common.ALLOCATE:
		investmentData = common.GenerateInvestmentData(data)
	case common.SIP:
		investmentData = common.GenerateInvestmentData(data)
	case common.CHANGE:
		investmentData = common.GenerateInvestmentData(data[:n-1])
		month = common.GetMonth(data[n-1])
	case common.BALANCE:
		month = common.GetMonth(data[0]) // balance command only has month
	case common.REBALANCE:
		// do nothing
	}
	return &CommandInfo{command, investmentData, month}
}

func (c CommandInfo) GetName() common.Command {
	return c.name
}

func (c *CommandInfo) GetData() common.InvestmentData {
	return c.data
}

func (c *CommandInfo) GetMonth() common.Month {
	return c.month
}
