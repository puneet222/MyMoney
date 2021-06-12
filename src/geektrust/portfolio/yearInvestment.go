package portfolio

import "geektrust/common"

type YearlyInvestment struct {
	Year int
	Investments [12]*Investment
}

func NewYearlyInvestment(year int, investments [12]*Investment) *YearlyInvestment {
	return &YearlyInvestment{Year: year, Investments: investments}
}

func (y *YearlyInvestment) UpdateInvestment(month common.Month, investment *Investment) {
	y.Investments[month] = investment
}

func (y *YearlyInvestment) GetInvestment(month common.Month) *Investment {
	return y.Investments[month]
}
