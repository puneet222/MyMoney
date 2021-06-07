package portfolio_tests

import (
	"geektrust/commander"
	"geektrust/common"
	"geektrust/portfolio"
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

var investments = []*portfolio.Investment{
	{6240, 3300, 1020},
	{7416, 6020, 1520},
	{10593, 7897, 2272},
	{13600, 8630, 2966},
}

var expectedPortfolio = portfolio.Portfolio{
	InvestmentHistory: [][]*portfolio.Investment{investments},
	Sip:               &portfolio.SIP{Equity: 2000, Debt: 1000, Gold: 500},
	Allocation:        &portfolio.Allocation{Equity: 60, Debt: 30, Gold: 10},
	LastRebalance:     nil,
	CurrentMonth:      common.APRIL,
	CurrentYearIndex:  0,
	StartYear:         year,
}

func TestBuildPortfolio(t *testing.T) {
	p := portfolio.BuildPortfolio(commands, year)
	if expectedPortfolio.String() != p.String() {
		t.Errorf("Error while creating new portfolio from commands, expected %v but got %v",
			expectedPortfolio.String(), p)
	}
}
