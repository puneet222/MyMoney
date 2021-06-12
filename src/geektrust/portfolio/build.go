package portfolio

import (
	"fmt"
	"geektrust/commander"
	"geektrust/common"
)

func BuildPortfolio(commands []*commander.CommandInfo, startYear int) *Portfolio {
	portfolio := &Portfolio{}
	for _, command := range commands {
		investment := Investment{command.Data.Equity, command.Data.Debt, command.Data.Gold}
		switch command.Name {
		case common.ALLOCATE:
			portfolio = NewPortfolio(&investment, startYear)
		case common.SIP:
			sip := SIP{command.Data.Equity, command.Data.Debt, command.Data.Gold}
			portfolio.SetSip(&sip)
		case common.CHANGE:
			roc := Change{command.Data.Equity, command.Data.Debt, command.Data.Gold}
			if command.Month == common.JANUARY {
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
			if command.Month <= portfolio.GetCurrentMonth() {
				fmt.Println(portfolio.GetInvestment(portfolio.GetCurrentYear(), command.Month).Output())
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






