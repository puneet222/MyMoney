package portfolio

import (
	"geektrust/common"
	"strconv"
	"strings"
)

type Portfolio struct {
	InvestmentHistory [][]*Investment // to store year also
	Sip *SIP
	Allocation *Allocation
	LastRebalance *Investment
	CurrentMonth common.Month
	CurrentYearIndex int
	StartYear int
}

func NewPortfolio(investment *Investment, startYear int) *Portfolio {
	investmentHistory := make([][]*Investment, 0)
	investments := make([]*Investment, 12, 12) // initialize yearly investments
	investments[common.JANUARY] = investment  // add January investment
	investmentHistory = append(investmentHistory, investments)
	p := &Portfolio{
		investmentHistory,
		nil,
		investment.GetAllocation(),
		nil,
		common.Month(0),
		0,
		startYear,
	}
	return p
}

func (p *Portfolio) AddInvestment(investment *Investment) {
	p.CurrentMonth++ // increment the month
	if p.CurrentMonth == 12 {
		// year got changed
		p.CurrentMonth = common.JANUARY // initialize month with January
		p.CurrentYearIndex++  // increment year
		// initialize new year's investment
		investments := make([]*Investment, 12, 12)
		// assign current investment to current month of year
		investments[p.CurrentMonth] = investment
		// update Portfolio
		p.InvestmentHistory = append(p.InvestmentHistory, investments)
	} else {
		p.InvestmentHistory[p.CurrentYearIndex][p.CurrentMonth] = investment
	}
	// check if re-balancing required
	if p.CurrentMonth == common.JUNE || p.CurrentMonth == common.DECEMBER {
		p.Rebalance()
	}
}

func (p *Portfolio) Rebalance() {
	currentInvestment := p.GetCurrentInvestment()
	totalInvestment := currentInvestment.GetTotalInvestment()
	rebalancedEquity := totalInvestment * (p.Allocation.Equity/100)
	rebalancedDebt := totalInvestment * (p.Allocation.Debt/100)
	rebalancedGold := totalInvestment * (p.Allocation.Gold/100)
	investment := &Investment{
		rebalancedEquity,
		rebalancedDebt,
		rebalancedGold,
	}
	investment.RoundOffInvestment()
	p.InvestmentHistory[p.CurrentYearIndex][p.CurrentMonth] = investment
	p.LastRebalance = p.GetCurrentInvestment()
}

func (p *Portfolio) SetSip(sip *SIP) {
	p.Sip = sip
}

func (p *Portfolio) GetCurrentInvestment() *Investment {
	return p.InvestmentHistory[p.CurrentYearIndex][p.CurrentMonth]
}

func (p *Portfolio) GetInvestment(year int, month common.Month) *Investment {
	return p.InvestmentHistory[year][month]
}

// to print current state of portfolio
func (p *Portfolio) String() string {
	sb := strings.Builder{}
	for yearIndex := 0; yearIndex < len(p.InvestmentHistory); yearIndex++ {
		year := p.StartYear + yearIndex
		header := "--------   " + strconv.Itoa(year) + "   --------\n"
		sb.WriteString(header)
		for _, investment := range p.InvestmentHistory[yearIndex] {
			if investment != nil {
				sb.WriteString(investment.String())
			}
		}
	}
	if p.Sip != nil {
		sb.WriteString(p.Sip.String())
	}
	if p.Allocation != nil {
		sb.WriteString(p.Allocation.String())
	}
	sb.WriteString("Current Month: " + p.CurrentMonth.String())
	sb.WriteString("\n---------------------------------------------------------\n")
	return sb.String()
}