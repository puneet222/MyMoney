package common

import (
	"fmt"
	"strconv"
	"strings"
)

type Month int

const (
	JANUARY Month = iota
	FEBRUARY
	MARCH
	APRIL
	MAY
	JUNE
	JULY
	AUGUST
	SEPTEMBER
	OCTOBER
	NOVEMBER
	DECEMBER
)

func (m Month) String() string {
	return [...]string{"JANUARY", "FEBRUARY", "MARCH", "APRIL", "MAY", "JUNE", "JULY", "AUGUST", "SEPTEMBER", "OCTOBER", "NOVEMBER", "DECEMBER"}[m]
}

func GetMonth(month string) Month {
	switch month {
	case "JANUARY":
		return JANUARY
	case "FEBRUARY":
		return FEBRUARY
	case "MARCH":
		return MARCH
	case "APRIL":
		return APRIL
	case "MAY":
		return MAY
	case "JUNE":
		return JUNE
	case "JULY":
		return JULY
	case "AUGUST":
		return AUGUST
	case "SEPTEMBER":
		return SEPTEMBER
	case "OCTOBER":
		return OCTOBER
	case "NOVEMBER":
		return NOVEMBER
	case "DECEMBER":
		return DECEMBER
	default:
		panic("month not supported")
	}
}

type Command int

const (
	ALLOCATE Command = iota
	SIP
	CHANGE
	BALANCE
	REBALANCE
)

func (c Command) String() string {
	return [...]string{"ALLOCATE", "SIP", "CHANGE", "BALANCE", "REBALANCE"}[c]
}

func GetCommand(command string) Command {
	switch command {
	case "ALLOCATE":
		return ALLOCATE
	case "SIP":
		return SIP
	case "CHANGE":
		return CHANGE
	case "BALANCE":
		return BALANCE
	case "REBALANCE":
		return REBALANCE
	default:
		panic("command not supported")
	}
}

type InvestmentData struct {
	Equity float64
	Debt float64
	Gold float64
}

func GenerateInvestmentData(data []string) InvestmentData {
	// clean data
	for i, d := range data {
		data[i] = strings.ReplaceAll(d, "%", "")
	}
	equity, err := strconv.ParseFloat(data[0], 64)
	if err != nil {
		fmt.Errorf("error on parsing float (equity) %v", err)
	}
	debt, err := strconv.ParseFloat(data[1], 64)
	if err != nil {
		fmt.Errorf("error on parsing float (debt) %v", err)
	}
	gold, err := strconv.ParseFloat(data[2], 64)
	if err != nil {
		fmt.Errorf("error on parsing float (gold) %v", err)
	}
	return InvestmentData{Equity: equity, Debt: debt, Gold: gold}
}

func VerifyDataSize(command Command, data []string) bool {
	n := len(data)
	switch command {
	case ALLOCATE:
		return n == 3
	case SIP:
		return n == 3
	case CHANGE:
		return n == 4
	case BALANCE:
		return n == 1
	case REBALANCE:
		return n == 0
	}
	return false
}
