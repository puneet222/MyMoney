package commander

import (
	"geektrust/common"
	"reflect"
	"testing"
)


func TestGenerateCommands(t *testing.T) {
	data := []byte("ALLOCATE 6000 3000 1000\nSIP 2000 1000 500\nCHANGE 4.00% 10.00% 2.15% FEBRUARY\nBALANCE MARCH\nREBALANCE\n")
	expected := []*CommandInfo{
		{common.ALLOCATE, InvestmentData{6000, 3000, 1000}, common.Month(0)},
		{common.SIP, InvestmentData{2000, 1000, 500}, common.Month(0)},
		{common.CHANGE, InvestmentData{4.00, 10.00, 2.15}, common.FEBRUARY},
		{common.BALANCE, InvestmentData{}, common.MARCH},
		{common.REBALACE, InvestmentData{}, common.Month(0)},
	}
	xci := GenerateCommands(data)
	if !reflect.DeepEqual(xci, expected) {
		t.Errorf("Error while generating commands expected %v but got %v", expected, xci)
	}
}

func TestGenerateInvestmentData(t *testing.T) {
	data := []string{"10.00%", "-20.45%", "3.56%"}
	expected := InvestmentData{10.00, -20.45, 3.56}
	id := GenerateInvestmentData(data)
	if !reflect.DeepEqual(id, expected) {
		t.Errorf("Error while generating investment data expected %v but got %v", expected, id)
	}
	// failing test case
	data = []string{"1Q.00%", "-2W.45%", "3C.56%"}
	expected = InvestmentData{10.00, -20.45, 3.56}
	id = GenerateInvestmentData(data)
	if reflect.DeepEqual(id, expected) {
		t.Errorf("Error data does not supposed to be equal %v & %v", expected, id)
	}
}

func TestNewCommandInfo(t *testing.T) {
	investmentData := InvestmentData{6000, 3000, 1000}
	command := common.ALLOCATE
	month := common.JANUARY
	expected := CommandInfo{command, investmentData, month}
	data := []string{"6000", "3000", "1000"}
	ci := *(NewCommandInfo(command, data)) // dereference as this function gives the address
	if !reflect.DeepEqual(ci, expected) {
		t.Errorf("Error while creating CommandInfo expected %v but got %v", expected, ci)
	}
}

func TestVerifyDataSize(t *testing.T) {
	allocate := common.ALLOCATE
	data := []string{"1000", "2000", "3000"}
	if !VerifyDataSize(allocate, data) {
		t.Errorf("Error on verify data  size for %s Command", allocate)
	}
	sip := common.SIP
	if !VerifyDataSize(sip, data) {
		t.Errorf("Error on verify data  size for %s Command", sip)
	}
	balance := common.BALANCE
	if VerifyDataSize(balance, data) {
		t.Errorf("Error on verify data  size for %s Command", balance)
	}
	rebalance := common.REBALACE
	if VerifyDataSize(rebalance, data) {
		t.Errorf("Error on verify data  size for %s Command", balance)
	}
	data = append(data, "TEST")
	change := common.CHANGE
	if !VerifyDataSize(change, data) {
		t.Errorf("Error on verify data  size for %s Command", balance)
	}
}