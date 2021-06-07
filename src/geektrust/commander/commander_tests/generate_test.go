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
		{common.ALLOCATE, commander.InvestmentData{Equity: 6000, Debt: 3000, Gold: 1000}, common.Month(0)},
		{common.SIP, commander.InvestmentData{Equity: 2000, Debt: 1000, Gold: 500}, common.Month(0)},
		{common.CHANGE, commander.InvestmentData{Equity: 4.00, Debt: 10.00, Gold: 2.15}, common.FEBRUARY},
		{common.BALANCE, commander.InvestmentData{}, common.MARCH},
		{common.REBALANCE, commander.InvestmentData{}, common.Month(0)},
	}
	xci := commander.GenerateCommands(data)
	if !reflect.DeepEqual(xci, expected) {
		t.Errorf("Error while generating commands expected %v but got %v", expected, xci)
	}
}

func TestGenerateInvestmentData(t *testing.T) {
	data := []string{"10.00%", "-20.45%", "3.56%"}
	expected := commander.InvestmentData{Equity: 10.00, Debt: -20.45, Gold: 3.56}
	id := commander.GenerateInvestmentData(data)
	if !reflect.DeepEqual(id, expected) {
		t.Errorf("Error while generating investment data expected %v but got %v", expected, id)
	}
	// failing test case
	data = []string{"1Q.00%", "-2W.45%", "3C.56%"}
	expected = commander.InvestmentData{Equity: 10.00, Debt: -20.45, Gold: 3.56}
	id = commander.GenerateInvestmentData(data)
	if reflect.DeepEqual(id, expected) {
		t.Errorf("Error data does not supposed to be equal %v & %v", expected, id)
	}
}

func TestNewCommandInfo(t *testing.T) {
	investmentData := commander.InvestmentData{Equity: 6000, Debt: 3000, Gold: 1000}
	command := common.ALLOCATE
	month := common.JANUARY
	expected := commander.CommandInfo{Name: command, Data: investmentData, Month: month}
	data := []string{"6000", "3000", "1000"}
	ci := *(commander.NewCommandInfo(command, data)) // dereference as this function gives the address
	if !reflect.DeepEqual(ci, expected) {
		t.Errorf("Error while creating CommandInfo expected %v but got %v", expected, ci)
	}
}

func TestVerifyDataSize(t *testing.T) {
	allocate := common.ALLOCATE
	data := []string{"1000", "2000", "3000"}
	if !commander.VerifyDataSize(allocate, data) {
		t.Errorf("Error on verify data  size for %s Command", allocate)
	}
	sip := common.SIP
	if !commander.VerifyDataSize(sip, data) {
		t.Errorf("Error on verify data  size for %s Command", sip)
	}
	balance := common.BALANCE
	if commander.VerifyDataSize(balance, data) {
		t.Errorf("Error on verify data  size for %s Command", balance)
	}
	rebalance := common.REBALANCE
	if commander.VerifyDataSize(rebalance, data) {
		t.Errorf("Error on verify data  size for %s Command", balance)
	}
	data = append(data, "TEST")
	change := common.CHANGE
	if !commander.VerifyDataSize(change, data) {
		t.Errorf("Error on verify data  size for %s Command", balance)
	}
}