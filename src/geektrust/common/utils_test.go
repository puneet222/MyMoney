package common

import (
	"testing"
)

var commandTest = []struct {
	name    string
	command Command
}{
	{"ALLOCATE", ALLOCATE},
	{"SIP", SIP},
	{"CHANGE", CHANGE},
	{"BALANCE", BALANCE},
	{"REBALANCE", REBALANCE},
}

var monthTest = []struct {
	name  string
	month Month
}{
	{"JANUARY", JANUARY},
	{"FEBRUARY", FEBRUARY},
	{"MARCH", MARCH},
	{"APRIL", APRIL},
	{"MAY", MAY},
	{"JUNE", JUNE},
	{"JULY", JULY},
	{"AUGUST", AUGUST},
	{"SEPTEMBER", SEPTEMBER},
	{"OCTOBER", OCTOBER},
	{"NOVEMBER", NOVEMBER},
	{"DECEMBER", DECEMBER},
}

func TestGetCommand(t *testing.T) {
	for _, c := range commandTest {
		if GetCommand(c.name) != c.command {
			t.Errorf("command mismatch expected %v but got %v", c.command, GetCommand(c.name))
		}
	}
	// should panic
	assertPanic(t, func() {GetCommand("TEST")})
}

func TestGetMonth(t *testing.T) {
	for _, c := range monthTest {
		if GetMonth(c.name) != c.month {
			t.Errorf("month mismatch expected %v but got %v", c.month, GetMonth(c.name))
		}
	}
	// should panic
	assertPanic(t, func() {GetMonth("MAYTEMBER")})
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
