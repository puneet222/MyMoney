package commander_tests

import (
	"geektrust/commander"
	"geektrust/common"
	"reflect"
	"testing"
)

func TestGenerateCommands(t *testing.T) {
	data := []byte("ALLOCATE 6000 3000 1000\nSIP 2000 1000 500\nCHANGE 4.00% 10.00% 2.15% FEBRUARY\nBALANCE MARCH\nREBALANCE\n")
	expected := []*commander.CommandInfo{
		{common.ALLOCATE, common.InvestmentData{Equity: 6000, Debt: 3000, Gold: 1000}, common.Month(0)},
		{common.SIP, common.InvestmentData{Equity: 2000, Debt: 1000, Gold: 500}, common.Month(0)},
		{common.CHANGE, common.InvestmentData{Equity: 4.00, Debt: 10.00, Gold: 2.15}, common.FEBRUARY},
		{common.BALANCE, common.InvestmentData{}, common.MARCH},
		{common.REBALANCE, common.InvestmentData{}, common.Month(0)},
	}
	xci := commander.GenerateCommands(data)
	if !reflect.DeepEqual(xci, expected) {
		t.Errorf("Error while generating commands expected %v but got %v", expected, xci)
	}
}