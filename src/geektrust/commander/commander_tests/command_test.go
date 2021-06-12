package commander_tests

import (
	"geektrust/commander"
	"geektrust/common"
	"reflect"
	"testing"
)

func TestCreateCommandInfo(t *testing.T) {
	investmentData := common.InvestmentData{Equity: 6000, Debt: 3000, Gold: 1000}
	command := common.ALLOCATE
	month := common.JANUARY
	expected := commander.NewCommandInfo(command, investmentData, month)
	data := []string{"6000", "3000", "1000"}
	ci := commander.CreateCommandInfo(command, data) // dereference as this function gives the address
	if !reflect.DeepEqual(ci, expected) {
		t.Errorf("Error while creating CommandInfo expected %v but got %v", expected, ci)
	}
}
