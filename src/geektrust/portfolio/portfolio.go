package portfolio

import (
	"geektrust/common"
	"strconv"
	"strings"
)

type Portfolio struct {
	investmentHistory [][]*Investment // to store year also
	sip *SIP
	allocation *Allocation
	lastRebalance *Investment
	currentMonth common.Month
	currentYearIndex int
	startYear int
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
	p.currentMonth++ // increment the month
	if p.currentMonth == 12 {
		// year got changed
		p.currentMonth = common.JANUARY // initialize month with January
		p.currentYearIndex++  // increment year
		// initialize new year's investment
		investments := make([]*Investment, 12, 12)
		// assign current investment to current month of year
		investments[p.currentMonth] = investment
		// update Portfolio
		p.investmentHistory = append(p.investmentHistory, investments)
	} else {
		p.investmentHistory[p.currentYearIndex][p.currentMonth] = investment
	}
	// check if re-balancing required
	if p.currentMonth == common.JUNE || p.currentMonth == common.DECEMBER {
		p.Rebalance()
	}
}

func (p *Portfolio) Rebalance() {
	currentInvestment := p.getCurrentInvestment()
	totalInvestment := currentInvestment.GetTotalInvestment()
	rebalancedEquity := totalInvestment * (p.allocation.Equity/100)
	rebalancedDebt := totalInvestment * (p.allocation.Debt/100)
	rebalancedGold := totalInvestment * (p.allocation.Gold/100)
	investment := &Investment{
		rebalancedEquity,
		rebalancedDebt,
		rebalancedGold,
	}
	investment.RoundOffInvestment()
	p.investmentHistory[p.currentYearIndex][p.currentMonth] = investment
	p.lastRebalance = p.getCurrentInvestment()
}

func (p *Portfolio) setSip(sip *SIP) {
	p.sip = sip
}

func (p *Portfolio) getCurrentInvestment() *Investment {
	return p.investmentHistory[p.currentYearIndex][p.currentMonth]
}

func (p *Portfolio) GetInvestment(year int, month common.Month) *Investment {
	return p.investmentHistory[year][month]
}

// to print current state of portfolio
func (p *Portfolio) String() string {
	sb := strings.Builder{}
	for yearIndex := 0; yearIndex < len(p.investmentHistory); yearIndex++ {
		year := p.startYear + yearIndex
		header := "--------   " + strconv.Itoa(year) + "   --------\n"
		sb.WriteString(header)
		for _, investment := range p.investmentHistory[yearIndex] {
			if investment != nil {
				sb.WriteString(investment.String())
			}
		}
	}
	sb.WriteString(p.sip.String())
	sb.WriteString(p.allocation.String())
	sb.WriteString("---------------------------------------------------------\n")
	return sb.String()
}