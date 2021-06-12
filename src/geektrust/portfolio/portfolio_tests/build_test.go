package portfolio_tests

import (
	"geektrust/commander"
	"geektrust/common"
	"geektrust/portfolio"
	"testing"
)

var year = 2021

var commands = []*commander.CommandInfo{
	{common.ALLOCATE, common.InvestmentData{Equity: 6000, Debt: 3000, Gold: 1000}, common.Month(0)},
	{common.SIP, common.InvestmentData{Equity: 2000, Debt: 1000, Gold: 500}, common.Month(0)},
	{common.CHANGE, common.InvestmentData{Equity: 4, Debt: 10, Gold: 2}, common.JANUARY},
	{common.CHANGE, common.InvestmentData{Equity: -10, Debt: 40, Gold: 0}, common.FEBRUARY},
	{common.CHANGE, common.InvestmentData{Equity: 12.5, Debt: 12.5, Gold: 12.5}, common.MARCH},
	{common.CHANGE, common.InvestmentData{Equity: 8, Debt: -3, Gold: 7}, common.APRIL},
	{common.BALANCE, common.InvestmentData{}, common.MARCH},
	{common.REBALANCE, common.InvestmentData{}, common.Month(0)},
}

var investments = [12]*portfolio.Investment{
	{6240, 3300, 1020},
	{7416, 6020, 1520},
	{10593, 7897, 2272},
	{13600, 8630, 2966},
}

var yearlyInvestments = portfolio.NewYearlyInvestment(year, investments)

var expectedPortfolio = portfolio.Portfolio{}

func setup() {
	expectedPortfolio.SetInvestmentHistory(map[int]*portfolio.YearlyInvestment{year: yearlyInvestments})
	expectedPortfolio.SetAllocation(&portfolio.Allocation{Equity: 60, Debt: 30, Gold: 10})
	expectedPortfolio.SetSip(&portfolio.SIP{Equity: 2000, Debt: 1000, Gold: 500})
	expectedPortfolio.SetLastRebalance(nil)
	expectedPortfolio.SetCurrentMonth(common.APRIL)
	expectedPortfolio.SetCurrentYear(year)
}

func TestBuildPortfolio(t *testing.T) {
	setup()
	p := portfolio.BuildPortfolio(commands, year)
	if expectedPortfolio.String() != p.String() {
		t.Errorf("Error while creating new portfolio from commands, expected %v but got %v",
			expectedPortfolio.String(), p)
	}
}
