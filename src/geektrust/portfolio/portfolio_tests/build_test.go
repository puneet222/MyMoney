package portfolio_tests

import (
	"geektrust/commander"
	"geektrust/common"
	"geektrust/portfolio"
	"testing"
)

var year = 2021

var commands []*commander.CommandInfo

var investments [12]*portfolio.Investment

var yearlyInvestments *portfolio.YearlyInvestment

var expectedPortfolio = portfolio.Portfolio{}

func setup() {
	commands = append(commands, commander.NewCommandInfo(common.ALLOCATE, common.InvestmentData{Equity: 6000, Debt: 3000, Gold: 1000}, common.Month(0)))
	commands = append(commands, commander.NewCommandInfo(common.SIP, common.InvestmentData{Equity: 2000, Debt: 1000, Gold: 500}, common.Month(0)))
	commands = append(commands, commander.NewCommandInfo(common.CHANGE, common.InvestmentData{Equity: 4, Debt: 10, Gold: 2}, common.JANUARY))
	commands = append(commands, commander.NewCommandInfo(common.CHANGE, common.InvestmentData{Equity: -10, Debt: 40, Gold: 0}, common.FEBRUARY))
	commands = append(commands, commander.NewCommandInfo(common.CHANGE, common.InvestmentData{Equity: 12.5, Debt: 12.5, Gold: 12.5}, common.MARCH))
	commands = append(commands, commander.NewCommandInfo(common.CHANGE, common.InvestmentData{Equity: 8, Debt: -3, Gold: 7}, common.APRIL))
	commands = append(commands, commander.NewCommandInfo(common.BALANCE, common.InvestmentData{}, common.MARCH))
	commands = append(commands, commander.NewCommandInfo(common.REBALANCE, common.InvestmentData{}, common.Month(0)))

	investments[0] = portfolio.NewInvestment(6240, 3300, 1020)
	investments[1] = portfolio.NewInvestment(7416, 6020, 1520)
	investments[2] = portfolio.NewInvestment(10593, 7897, 2272)
	investments[3] = portfolio.NewInvestment(13600, 8630, 2966)

	yearlyInvestments = portfolio.NewYearlyInvestment(year, investments)

	expectedPortfolio.SetInvestmentHistory(map[int]*portfolio.YearlyInvestment{year: yearlyInvestments})
	expectedPortfolio.SetAllocation(portfolio.NewAllocation(60, 30, 10))
	expectedPortfolio.SetSip(portfolio.NewSip(2000, 1000, 500))
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
