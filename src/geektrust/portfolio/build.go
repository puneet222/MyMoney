package portfolio

import (
	"fmt"
	"geektrust/commander"
	"geektrust/common"
)

func BuildPortfolio(commands []*commander.CommandInfo, startYear int) *Portfolio {
	portfolio := &Portfolio{}
	for _, command := range commands {
		investment := NewInvestment(command.GetData().Equity, command.GetData().Debt, command.GetData().Gold)
		switch command.GetName() {
		case common.ALLOCATE:
			portfolio = NewPortfolio(investment, startYear)
		case common.SIP:
			sip := NewSip(command.GetData().Equity, command.GetData().Debt, command.GetData().Gold)
			portfolio.SetSip(sip)
		case common.CHANGE:
			roc := NewChange(command.GetData().Equity, command.GetData().Debt, command.GetData().Gold)
			if command.GetMonth() == common.JANUARY {
				// just update the current investment
				portfolio.GetCurrentInvestment().UpdateInvestment(roc)
			} else {
				// create copy of last investment
				newInvestment := portfolio.GetCurrentInvestment().DeepCopy()
				// add sip
				newInvestment.AddSIP(portfolio.GetSip())
				// update investment based on change
				newInvestment.UpdateInvestment(roc)
				newInvestment.RoundOffInvestment()
				portfolio.AddInvestment(newInvestment)
			}
		case common.BALANCE:
			if command.GetMonth() <= portfolio.GetCurrentMonth() {
				fmt.Println(portfolio.GetInvestment(portfolio.GetCurrentYear(), command.GetMonth()).Output())
			}
		case common.REBALANCE:
			if portfolio.GetLastRebalance() != nil {
				fmt.Println(portfolio.GetLastRebalance().Output())
			} else {
				fmt.Println("CANNOT_REBALANCE")
			}
		}
	}
	return portfolio
}






