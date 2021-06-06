package portfolio

import (
	"geektrust/common"
	"testing"
)

func getMockPortfolio() *Portfolio{
	investment := Investment{6000, 3000, 1000}
	return NewPortfolio(&investment, 2021)
}

func TestNewPortfolio(t *testing.T) {
	investment := Investment{6000, 3000, 1000}
	year := 2021
	portfolio := NewPortfolio(&investment, year)
	// create expected portfolio
	investmentHistory := make([][]*Investment, 0)
	investments := make([]*Investment, 12, 12) // initialize yearly investments
	investments[common.JANUARY] = &investment  // add January investment
	investmentHistory = append(investmentHistory, investments)
	expected := Portfolio{
		investmentHistory: investmentHistory,
		sip:               nil,
		allocation:        investment.GetAllocation(),
		lastRebalance:     nil,
		currentMonth:      0,
		currentYearIndex:  0,
		startYear:         year,
	}
	if expected.String() != portfolio.String() {
		t.Errorf("Error while creating new portfolio, expected %v but got %v", expected, portfolio)
	}
}

func TestPortfolio_AddInvestment(t *testing.T) {
	investment := Investment{6000, 3000, 1000}
	portfolio := NewPortfolio(&investment, 2021)
	investments := []*Investment{
		{1000, 1000, 1000},    // feb
		{1000, 1000, 1000},    // mar
		{1000, 1000, 1000},    // apr
		{1000, 1000, 1000},    // may
		{1000, 1000, 1000},    // jun
		{1000, 1000, 1000},    // jul
		{1000, 1000, 1000},    // aug
		{1000, 1000, 1000},    // sep
		{1000, 1000, 1000},    // oct
		{1000, 1000, 1000},    // nov
		{1000, 1000, 1000},    // dec
		{1000, 1000, 1000},    // jan (new year)
	}
	expected := Investment{1800, 900, 300}
	for i := 0; i < len(investments); i++ {
		// till june
		portfolio.AddInvestment(investments[i])
		if i == 4 {
			// check june rebalance
			if expected.String() != portfolio.GetCurrentInvestment().String() {
				t.Errorf("Error while adding investment (rebalancing june), expected %v but got %v",
					expected, portfolio.GetCurrentInvestment())
			}
		}
		if i == 10 {
			// check december rebalance
			if expected.String() != portfolio.GetCurrentInvestment().String() {
				t.Errorf("Error while adding investment (rebalancing dec), expected %v but got %v",
					expected, portfolio.GetCurrentInvestment())
			}
		}
		if i == 11 {
			// check for year change
			if portfolio.currentYearIndex != 1 {
				t.Errorf("Year not updated on adding investment, expected 1 but got %d", portfolio.currentYearIndex)
			}
			// check for month update
			if portfolio.currentMonth != common.JANUARY {
				t.Errorf("Month not updated on adding investment, expected %s but got %s",
					common.JANUARY, portfolio.currentMonth)
			}
		}
	}


	// check for
}

func TestPortfolio_Rebalance(t *testing.T) {
	investment := Investment{6000, 3000, 1000}
	portfolio := NewPortfolio(&investment, 2021)
	newInvestment := Investment{3000, 5000, 2000}
	portfolio.AddInvestment(&newInvestment)
	portfolio.Rebalance()
	expected := Investment{6000, 3000, 1000}
	if expected.String() != portfolio.GetCurrentInvestment().String() {
		t.Errorf("Error while rebalancing investement, expected %v but got %v",
			expected, portfolio.GetCurrentInvestment())
	}
}

func TestPortfolio_SetSip(t *testing.T) {
	sip := SIP{2000, 1000, 500}
	portfolio := getMockPortfolio()
	portfolio.SetSip(&sip)
	if portfolio.sip.Equity != 2000 || portfolio.sip.Debt != 1000 || portfolio.sip.Gold != 500 {
		t.Errorf("Error while portfolio SPI, expected %v but got %v", sip, portfolio.sip)
	}
}

func TestPortfolio_GetCurrentInvestment(t *testing.T) {
	investment1 := Investment{2000, 1000, 500}
	portfolio := getMockPortfolio()
	portfolio.AddInvestment(&investment1)
	expectedInvestment := Investment{2000, 1000, 500}
	if expectedInvestment.String() != portfolio.GetCurrentInvestment().String() {
		t.Errorf("Error while getting portfolio's current investment expected %v but got %v",
			expectedInvestment, portfolio.GetCurrentInvestment())
	}
}

func TestPortfolio_GetInvestment(t *testing.T) {
	investment := Investment{2000, 3000, 4000}
	expected := Investment{2000, 3000, 4000}
	portfolio := getMockPortfolio()
	portfolio.AddInvestment(&investment) // will add as february's investment
	if expected.String() != portfolio.GetInvestment(portfolio.currentYearIndex, common.FEBRUARY).String() {
		t.Errorf("Error while getting specific month's investment, expected %v but got %v",
			expected, portfolio.GetInvestment(portfolio.currentYearIndex, common.FEBRUARY))
	}
}