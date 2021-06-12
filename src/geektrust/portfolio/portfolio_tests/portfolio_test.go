package portfolio_tests

import (
	"geektrust/common"
	"geektrust/portfolio"
	"testing"
)

func getMockPortfolio() *portfolio.Portfolio {
	investment := portfolio.NewInvestment(6000, 3000, 1000)
	return portfolio.NewPortfolio(investment, 2021)
}

//func TestNewPortfolio(t *testing.T) {
//	investment := portfolio.Investment{Equity: 6000, Debt: 3000, Gold: 1000}
//	year := 2021
//	p := portfolio.NewPortfolio(&investment, year)
//	// create expected portfolio
//	investmentHistory := make(map[int]*portfolio.YearlyInvestment)
//	var investments [12]*portfolio.Investment // initialize yearly investments
//	investments[common.JANUARY] = &investment        // add January investment
//	investmentHistory[year] = portfolio.NewYearlyInvestment(2021, investments)
//	expected := portfolio.Portfolio{
//		InvestmentHistory: investmentHistory,
//		sip:               nil,
//		Allocation:        investment.GetAllocation(),
//		LastRebalance:     nil,
//		CurrentMonth:      0,
//		CurrentYear:         year,
//	}
//	if expected.String() != p.String() {
//		t.Errorf("Error while creating new portfolio, expected %v but got %v", expected, p)
//	}
//}

func TestPortfolio_AddInvestment(t *testing.T) {
	investment := portfolio.NewInvestment(6000, 3000, 1000)
	p := portfolio.NewPortfolio(investment, 2021)

	var investments []*portfolio.Investment
	for m := common.FEBRUARY; m <= common.DECEMBER; m++ {
		investments = append(investments, portfolio.NewInvestment(1000, 1000, 1000)) // feb to december
	}
	investments = append(investments, portfolio.NewInvestment(1000, 1000, 1000))    // jan (new year)

	expected := portfolio.NewInvestment(1800, 900, 300)
	for i := 0; i < len(investments); i++ {
		// till june
		p.AddInvestment(investments[i])
		if i == 4 {
			// check june rebalance
			if expected.String() != p.GetCurrentInvestment().String() {
				t.Errorf("Error while adding investment (rebalancing june), expected %v but got %v",
					expected, p.GetCurrentInvestment())
			}
		}
		if i == 10 {
			// check december rebalance
			if expected.String() != p.GetCurrentInvestment().String() {
				t.Errorf("Error while adding investment (rebalancing dec), expected %v but got %v",
					expected, p.GetCurrentInvestment())
			}
		}
		if i == 11 {
			// check for year change
			if p.GetCurrentYear() == 2021 {
				t.Errorf("Year not updated on adding investment, expected 2022 but got %d", p.GetCurrentYear())
			}
			// check for month update
			if p.GetCurrentMonth() != common.JANUARY {
				t.Errorf("Month not updated on adding investment, expected %s but got %s",
					common.JANUARY, p.GetCurrentMonth())
			}
		}
	}


	// check for
}

func TestPortfolio_Rebalance(t *testing.T) {
	investment := portfolio.NewInvestment(6000, 3000, 1000)
	p := portfolio.NewPortfolio(investment, 2021)
	newInvestment := portfolio.NewInvestment(3000, 5000, 2000)
	p.AddInvestment(newInvestment)
	p.Rebalance()
	expected := portfolio.NewInvestment(6000, 3000, 1000)
	if expected.String() != p.GetCurrentInvestment().String() {
		t.Errorf("Error while rebalancing investement, expected %v but got %v",
			expected, p.GetCurrentInvestment())
	}
}

func TestPortfolio_SetSip(t *testing.T) {
	sip := portfolio.NewSip(2000, 1000, 500)
	p := getMockPortfolio()
	p.SetSip(sip)
	if p.GetSip().GetEquity() != 2000 || p.GetSip().GetDebt() != 1000 || p.GetSip().GetGold() != 500 {
		t.Errorf("Error while portfolio SPI, expected %v but got %v", sip, p.GetSip())
	}
}

func TestPortfolio_GetCurrentInvestment(t *testing.T) {
	investment1 := portfolio.NewInvestment(2000, 1000, 500)
	p := getMockPortfolio()
	p.AddInvestment(investment1)
	expectedInvestment := portfolio.NewInvestment(2000, 1000, 500)
	if expectedInvestment.String() != p.GetCurrentInvestment().String() {
		t.Errorf("Error while getting portfolio's current investment expected %v but got %v",
			expectedInvestment, p.GetCurrentInvestment())
	}
}

func TestPortfolio_GetInvestment(t *testing.T) {
	investment := portfolio.NewInvestment(2000, 3000, 4000)
	expected := portfolio.NewInvestment(2000, 3000, 4000)
	p := getMockPortfolio()
	p.AddInvestment(investment) // will add as february's investment
	if expected.String() != p.GetInvestment(p.GetCurrentYear(), common.FEBRUARY).String() {
		t.Errorf("Error while getting specific month's investment, expected %v but got %v",
			expected, p.GetInvestment(p.GetCurrentYear(), common.FEBRUARY))
	}
}