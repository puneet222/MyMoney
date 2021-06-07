package common_tests

import (
	"geektrust/common"
	"testing"
)

var commandTest = []struct {
	name    string
	command common.Command
}{
	{"ALLOCATE", common.ALLOCATE},
	{"SIP", common.SIP},
	{"CHANGE", common.CHANGE},
	{"BALANCE", common.BALANCE},
	{"REBALANCE", common.REBALANCE},
}

var monthTest = []struct {
	name  string
	month common.Month
}{
	{"JANUARY", common.JANUARY},
	{"FEBRUARY", common.FEBRUARY},
	{"MARCH", common.MARCH},
	{"APRIL", common.APRIL},
	{"MAY", common.MAY},
	{"JUNE", common.JUNE},
	{"JULY", common.JULY},
	{"AUGUST", common.AUGUST},
	{"SEPTEMBER", common.SEPTEMBER},
	{"OCTOBER", common.OCTOBER},
	{"NOVEMBER", common.NOVEMBER},
	{"DECEMBER", common.DECEMBER},
}

func TestGetCommand(t *testing.T) {
	for _, c := range commandTest {
		if common.GetCommand(c.name) != c.command {
			t.Errorf("command mismatch expected %v but got %v", c.command, common.GetCommand(c.name))
		}
	}
	// should panic
	assertPanic(t, func() { common.GetCommand("TEST")})
}

func TestGetMonth(t *testing.T) {
	for _, c := range monthTest {
		if common.GetMonth(c.name) != c.month {
			t.Errorf("month mismatch expected %v but got %v", c.month, common.GetMonth(c.name))
		}
	}
	// should panic
	assertPanic(t, func() { common.GetMonth("MAYTEMBER")})
}

func TestCommand_String(t *testing.T) {
	for _, c := range commandTest {
		command := c.command
		if command.String() != c.name {
			t.Errorf("error while converting command to string expected %s got %s", c.name, command.String())
		}
	}
}

func TestMonth_String(t *testing.T) {
	for _, m := range monthTest {
		month := m.month
		if month.String() != m.name {
			t.Errorf("error while converting month to string expected %s got %s", m.name, month.String())
		}
	}
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
