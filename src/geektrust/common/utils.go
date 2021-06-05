package common

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
	REBALACE
)

func (c Command) String() string {
	return [...]string{"ALLOCATE", "SIP", "CHANGE", "BALANCE", "REBALACE"}[c]
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
		return REBALACE
	default:
		panic("command not supported")
	}
}