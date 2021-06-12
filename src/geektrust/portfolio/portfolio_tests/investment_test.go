package portfolio_tests

import (
	"geektrust/portfolio"
	"testing"
)

func getMockInvestment() *portfolio.Investment {
	return  portfolio.NewInvestment(6000, 3000, 1000)
}

func TestInvestment_AddSIP(t *testing.T) {
	investment := getMockInvestment()
	sip := portfolio.NewSip(1500, 1000, 500)
	expected := portfolio.NewInvestment(7500, 4000, 1500)
	investment.AddSIP(sip)
	if expected.String() != investment.String() {
		t.Errorf("Error on adding SIP to investment expected %v but got %v", expected, investment)
	}
}

func TestInvestment_UpdateInvestment(t *testing.T) {
	investment := getMockInvestment()
	change := portfolio.NewChange(-5, 4, 10)
	expected := portfolio.NewInvestment(5700, 3120, 1100)
	investment.UpdateInvestment(change)
	if expected.String() != investment.String() {
		t.Errorf("Error on adding change to investment expected %v but got %v", expected, investment)
	}
}

func TestAllocation_String(t *testing.T) {
	investment := getMockInvestment()
	expected := portfolio.NewAllocation(60, 30, 10)
	if expected.String() != investment.GetAllocation().String() {
		t.Errorf("Error while getting allocation expected %v but got %v", expected, investment)
	}
}

func TestInvestment_RoundOffInvestment(t *testing.T) {
	investment := portfolio.NewInvestment(6509.987, 2345.54, 425.56)
	expected := portfolio.NewInvestment(6509, 2345, 425)
	investment.RoundOffInvestment()
	if expected.String() != investment.String() {
		t.Errorf("Error while rounding off expected %v but got %v", expected, investment)
	}
}

func TestInvestment_DeepCopy(t *testing.T) {
	investment := getMockInvestment()
	icopy := investment.DeepCopy()
	icopy.SetEquity(100000)
	if icopy.String() == investment.String() {
		t.Errorf("Error on deep copying - returns the same reference")
	}
}

func TestInvestment_Output(t *testing.T) {
	investment := portfolio.NewInvestment(2345, 5437, 892)
	expected := "2345 5437 892"
	if investment.Output() != expected {
		t.Errorf("error in investment output expected %s but got %s", expected, investment.Output())
	}
}


