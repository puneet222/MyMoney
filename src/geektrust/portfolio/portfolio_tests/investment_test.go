package portfolio_tests

import (
	"geektrust/portfolio"
	"testing"
)

func getMockInvestment() *portfolio.Investment {
	return &portfolio.Investment{Equity: 6000, Debt: 3000, Gold: 1000}
}

func TestInvestment_AddSIP(t *testing.T) {
	investment := getMockInvestment()
	sip := portfolio.SIP{Equity: 1500, Debt: 1000, Gold: 500}
	expected := portfolio.Investment{Equity: 7500, Debt: 4000, Gold: 1500}
	investment.AddSIP(&sip)
	if expected.String() != investment.String() {
		t.Errorf("Error on adding SIP to investment expected %v but got %v", expected, investment)
	}
}

func TestInvestment_UpdateInvestment(t *testing.T) {
	investment := getMockInvestment()
	change := portfolio.Change{Equity: -5, Debt: 4, Gold: 10}
	expected := portfolio.Investment{Equity: 5700, Debt: 3120, Gold: 1100}
	investment.UpdateInvestment(change)
	if expected.String() != investment.String() {
		t.Errorf("Error on adding change to investment expected %v but got %v", expected, investment)
	}
}

func TestAllocation_String(t *testing.T) {
	investment := getMockInvestment()
	expected := portfolio.Allocation{Equity: 60, Debt: 30, Gold: 10}
	if expected.String() != investment.GetAllocation().String() {
		t.Errorf("Error while getting allocation expected %v but got %v", expected, investment)
	}
}

func TestInvestment_RoundOffInvestment(t *testing.T) {
	investment := portfolio.Investment{Equity: 6509.987, Debt: 2345.54, Gold: 425.56}
	expected := portfolio.Investment{Equity: 6509, Debt: 2345, Gold: 425}
	investment.RoundOffInvestment()
	if expected.String() != investment.String() {
		t.Errorf("Error while rounding off expected %v but got %v", expected, investment)
	}
}

func TestInvestment_DeepCopy(t *testing.T) {
	investment := getMockInvestment()
	icopy := investment.DeepCopy()
	icopy.Equity = 100000
	if icopy.String() == investment.String() {
		t.Errorf("Error on deep copying - returns the same reference")
	}
}

func TestInvestment_Output(t *testing.T) {
	investment := portfolio.Investment{Equity: 2345, Debt: 5437, Gold: 892}
	expected := "2345 5437 892"
	if investment.Output() != expected {
		t.Errorf("error in investment output expected %s but got %s", expected, investment.Output())
	}
}


