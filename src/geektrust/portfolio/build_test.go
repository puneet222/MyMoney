package portfolio

import (
	"geektrust/commander"
	"geektrust/common"
	"testing"
)

var year = 2021

var commands = []*commander.CommandInfo{
	{common.ALLOCATE, commander.InvestmentData{Equity: 6000, Debt: 3000, Gold: 1000}, common.Month(0)},
	{common.SIP, commander.InvestmentData{Equity: 2000, Debt: 1000, Gold: 500}, common.Month(0)},
	{common.CHANGE, commander.InvestmentData{Equity: 4, Debt: 10, Gold: 2}, common.JANUARY},
	{common.CHANGE, commander.InvestmentData{Equity: -10, Debt: 40, Gold: 0}, common.FEBRUARY},
	{common.CHANGE, commander.InvestmentData{Equity: 12.5, Debt: 12.5, Gold: 12.5}, common.MARCH},
	{common.CHANGE, commander.InvestmentData{Equity: 8, Debt: -3, Gold: 7}, common.APRIL},
	{common.BALANCE, commander.InvestmentData{}, common.MARCH},
	{common.REBALANCE, commander.InvestmentData{}, common.Month(0)},
}

var investments = []*Investment{
	{6240, 3300, 1020},
	{7416, 6020, 1520},
	{10593, 7897, 2272},
	{13600, 8630, 2966},
}

var expectedPortfolio = Portfolio{
	investmentHistory: [][]*Investment{investments},
	sip:               &SIP{2000, 1000, 500},
	allocation:        &Allocation{60, 30, 10},
	lastRebalance:     nil,
	currentMonth:      common.APRIL,
	currentYearIndex:  0,
	startYear:         year,
}

func TestBuildPortfolio(t *testing.T) {
	portfolio := BuildPortfolio(commands, year)
	if expectedPortfolio.String() != portfolio.String() {
		t.Errorf("Error while creating new portfolio from commands, expected %v but got %v",
			expectedPortfolio.String(), portfolio)
	}
}
