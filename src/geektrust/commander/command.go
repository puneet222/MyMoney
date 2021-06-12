package commander

import (
	"fmt"
	"geektrust/common"
)

type CommandInfo struct {
	Name common.Command
	Data common.InvestmentData
	Month common.Month
}

func NewCommandInfo(command common.Command, data []string) *CommandInfo {
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
