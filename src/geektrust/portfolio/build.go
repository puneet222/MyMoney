package portfolio

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"geektrust/commander"
	"geektrust/common"
	"strings"
)

type Investment struct {
	Equity float64
	Debt float64
	Gold float64
}

func (i *Investment) DeepCopy() *Investment {
	// note: no error handling below
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	err := encoder.Encode(i)
	if err != nil {
		panic("error while deep copying (encoding) investment object")
	}

	decoder := gob.NewDecoder(&b)
	copy := Investment{}
	err = decoder.Decode(&copy)
	if err != nil {
		panic("error while deep copying (decoding) investment object")
	}
	return &copy
}

func (i *Investment) AddSIP(sip SIP) {
	i.Equity += sip.Equity
	i.Debt += sip.Debt
	i.Gold += sip.Gold
}

func (i *Investment) UpdateInvestment(change Change) {
	i.Equity = i.Equity + (change.Equity/100)*i.Equity
	i.Debt = i.Debt + (change.Debt/100)*i.Debt
	i.Gold = i.Gold + (change.Gold/100)*i.Gold
}

func (i *Investment) GetTotalInvestment() float64 {
	return i.Equity + i.Debt + i.Gold
}

func (i *Investment) RoundOffInvestment() {
	i.Equity = float64(int(i.Equity))
	i.Debt = float64(int(i.Debt))
	i.Gold = float64(int(i.Gold))
}

func (i *Investment) GetAllocation() *Allocation {
	equityAllocation := (i.Equity*100)/i.GetTotalInvestment()
	debtAllocation := (i.Debt*100)/i.GetTotalInvestment()
	goldAllocation := (i.Gold*100)/i.GetTotalInvestment()
	return &Allocation{equityAllocation, debtAllocation, goldAllocation}
}

func (i *Investment) String() string {
	return fmt.Sprintf("Equity %f | Debt %f | Gold %f", i.Equity, i.Debt, i.Gold)
}

type SIP Investment

type Change Investment

type Allocation Investment

type Portfolio struct {
	investmentHistory [][]*Investment
	sip SIP
	allocation *Allocation
	lastRebalance *Investment
	currentMonth common.Month
	currentYearIndex int
	startYear int
}

func NewPortfolio(investment *Investment, startYear int) *Portfolio {
	investmentHistory := make([][]*Investment, 0)
	investments := make([]*Investment, 12, 12)
	investments[common.JANUARY] = investment  // add January investment
	investmentHistory = append(investmentHistory, investments)
	fmt.Println("investemnt history", common.Month(0), startYear)
	p := &Portfolio{
		investmentHistory,
		SIP(Investment{}),
		investment.GetAllocation(),
		&Investment{},
		common.Month(0),
		0,
		startYear,
	}
	fmt.Println(p)
	return p
}

func (p *Portfolio) String() string {
	sb := strings.Builder{}
	sb.WriteString("---------------------------------------------------------------\n")
	for yearIndex := 0; yearIndex < len(p.investmentHistory); yearIndex++ {
		for _, investment := range p.investmentHistory[yearIndex] {
			if investment != nil {
				sb.WriteString("| ")
				sb.WriteString(investment.String())
				sb.WriteString(" |\n")
			}
		}
	}
	sb.WriteString(fmt.Sprintf("Month: %s SIP: %v\n", p.currentMonth, p.sip))
	sb.WriteString("---------------------------------------------------------------\n")
	return sb.String()
}

func (p *Portfolio) setSip(sip SIP) {
	p.sip = sip
}

func (p *Portfolio) GetInvestment(year int, month common.Month) *Investment {
	return p.investmentHistory[year][month]
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

func (p *Portfolio) addInvestment(investment *Investment) {
	p.currentMonth++ // increment the month
	if p.currentMonth == 12 {
		// year got changed
		p.currentMonth = common.JANUARY
		p.currentYearIndex++
		p.investmentHistory = append(p.investmentHistory, []*Investment{investment})
	} else {
		p.investmentHistory[p.currentYearIndex][p.currentMonth] = investment
	}
	// check if re-balancing required
	if p.currentMonth == common.JUNE || p.currentMonth == common.DECEMBER {
		p.Rebalance()
	}
}

func (p *Portfolio) getCurrentInvestment() *Investment {
	return p.investmentHistory[p.currentYearIndex][p.currentMonth]
}

func BuildPortfolio(commands []*commander.CommandInfo, startYear int) *Portfolio {
	portfolio := &Portfolio{}
	for _, command := range commands {
		investment := Investment{command.Data.Equity, command.Data.Debt, command.Data.Gold}
		switch command.Name {
		case common.ALLOCATE:
			portfolio = NewPortfolio(&investment, startYear)
		case common.SIP:
			sip := SIP{command.Data.Equity, command.Data.Debt, command.Data.Gold}
			portfolio.setSip(sip) // converting investment to sip type
		case common.CHANGE:
			roc := Change{command.Data.Equity, command.Data.Debt, command.Data.Gold}
			if command.Month == common.JANUARY {
				// just update the current investment
				portfolio.getCurrentInvestment().UpdateInvestment(roc)
			} else {
				// add sip
				newInvestment := portfolio.getCurrentInvestment().DeepCopy()
				newInvestment.AddSIP(portfolio.sip)
				// update investment based on change
				newInvestment.UpdateInvestment(roc)
				newInvestment.RoundOffInvestment()
				portfolio.addInvestment(newInvestment)
			}
		case common.BALANCE:
			if command.Month <= portfolio.currentMonth {
				fmt.Println(portfolio.GetInvestment(portfolio.currentYearIndex, portfolio.currentMonth))
			}
		case common.REBALACE:
			if portfolio.lastRebalance != nil {
				fmt.Println(portfolio.lastRebalance)
			} else {
				fmt.Println("CANNOT_REBALANCE")
			}
		}
		fmt.Println(portfolio)
	}
	return portfolio
}






