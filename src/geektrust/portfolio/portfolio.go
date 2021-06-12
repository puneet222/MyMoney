package portfolio

import (
	"geektrust/common"
	"strconv"
	"strings"
)

type Portfolio struct {
	investmentHistory map[int]*YearlyInvestment
	sip *SIP
	allocation *Allocation
	lastRebalance *Investment
	currentMonth common.Month
	currentYear int
}

func NewPortfolio(investment *Investment, startYear int) *Portfolio {
	investmentHistory := make(map[int]*YearlyInvestment)
	var investments [12]*Investment // initialize yearly investments
	investments[common.JANUARY] = investment  // add January investment
	investmentHistory[startYear] = NewYearlyInvestment(startYear, investments)
	p := &Portfolio{
		investmentHistory,
		nil,
		investment.GetAllocation(),
		nil,
		common.Month(0),
		startYear,
	}
	return p
}

func (p *Portfolio) AddInvestment(investment *Investment) {
	p.IncrementCurrentMonth() // increment the month
	if p.GetCurrentMonth() == 12 {
		// year got changed
		p.SetCurrentMonth(common.JANUARY) // initialize month with January
		p.IncrementCurrentYear()  // increment year
		// initialize new year's investment
		var investments [12]*Investment
		// assign current investment to current month of year
		investments[p.GetCurrentMonth()] = investment
		// update Portfolio
		p.SetYearlyInvestment(p.GetCurrentYear(), NewYearlyInvestment(p.GetCurrentYear(), investments))
	} else {
		p.UpdateCurrentInvestment(investment)
	}
	// check if re-balancing required
	if p.GetCurrentMonth() == common.JUNE || p.GetCurrentMonth() == common.DECEMBER {
		p.Rebalance()
	}
}

func (p *Portfolio) Rebalance() {
	currentInvestment := p.GetCurrentInvestment()
	totalInvestment := currentInvestment.GetTotalInvestment()
	rebalancedEquity := totalInvestment * (p.GetAllocation().GetEquity()/100)
	rebalancedDebt := totalInvestment * (p.GetAllocation().GetDebt()/100)
	rebalancedGold := totalInvestment * (p.GetAllocation().GetGold()/100)
	investment := NewInvestment(rebalancedEquity, rebalancedDebt, rebalancedGold)
	investment.RoundOffInvestment()
	p.UpdateCurrentInvestment(investment)
	p.SetLastRebalance(p.GetCurrentInvestment())
}

func (p *Portfolio) GetInvestmentHistory()	map[int]*YearlyInvestment  {
	return p.investmentHistory
}

func (p *Portfolio) SetInvestmentHistory(ih map[int]*YearlyInvestment)  {
	p.investmentHistory = ih
}

func (p *Portfolio) GetYearlyInvestment(year int) *YearlyInvestment  {
	return p.investmentHistory[year]
}

func (p *Portfolio) GetInvestment(year int, month common.Month) *Investment {
	return p.investmentHistory[year].GetMonthlyInvestment(month)
}

func (p *Portfolio) GetCurrentInvestment() *Investment {
	return p.investmentHistory[p.GetCurrentYear()].GetMonthlyInvestment(p.GetCurrentMonth())
}

func (p *Portfolio) SetYearlyInvestment(year int, investment *YearlyInvestment) {
	p.investmentHistory[year] = investment
}

func (p *Portfolio) UpdateInvestment(year int, month common.Month, investment *Investment) {
	p.investmentHistory[year].UpdateInvestment(month, investment)
}

func (p *Portfolio) UpdateCurrentInvestment(investment *Investment) {
	p.investmentHistory[p.GetCurrentYear()].UpdateInvestment(p.GetCurrentMonth(), investment)
}

func (p *Portfolio) GetSip() *SIP {
	return p.sip
}

func (p *Portfolio) SetSip(sip *SIP) {
	p.sip = sip
}

func (p *Portfolio) GetAllocation() *Allocation {
	return p.allocation
}

func (p *Portfolio) SetAllocation(allocation *Allocation) {
	p.allocation = allocation
}

func (p *Portfolio) GetLastRebalance() *Investment {
	return p.lastRebalance
}

func (p *Portfolio) SetLastRebalance(investment *Investment) {
	p.lastRebalance = investment
}

func (p *Portfolio) GetCurrentMonth() common.Month {
	return p.currentMonth
}

func (p *Portfolio) SetCurrentMonth(month common.Month)  {
	p.currentMonth = month
}

func (p *Portfolio) IncrementCurrentMonth() {
	p.currentMonth++
}

func (p *Portfolio) GetCurrentYear() int {
	return p.currentYear
}

func (p *Portfolio) SetCurrentYear(year int)  {
	p.currentYear = year
}

func (p *Portfolio) IncrementCurrentYear() {
	p.currentYear++
}

// to print current state of portfolio
func (p *Portfolio) String() string {
	sb := strings.Builder{}
	for year, yearlyInvestments := range p.GetInvestmentHistory() {
		header := "--------   " + strconv.Itoa(year) + "   --------\n"
		sb.WriteString(header)
		for _, investment := range yearlyInvestments.GetInvestments() {
			if investment != nil {
				sb.WriteString(investment.String())
			}
		}
	}
	if p.GetSip() != nil {
		sb.WriteString(p.GetSip().String())
	}
	if p.GetAllocation() != nil {
		sb.WriteString(p.GetAllocation().String())
	}
	sb.WriteString("Current Month: " + p.GetCurrentMonth().String())
	sb.WriteString("\n---------------------------------------------------------\n")
	return sb.String()
}