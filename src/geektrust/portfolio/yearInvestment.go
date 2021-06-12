package portfolio

import "geektrust/common"

type YearlyInvestment struct {
	year int
	investments [12]*Investment
}

func NewYearlyInvestment(year int, investments [12]*Investment) *YearlyInvestment {
	return &YearlyInvestment{year: year, investments: investments}
}

func (y *YearlyInvestment) GetYear() int {
	return y.year
}

func (y *YearlyInvestment) SetYear(year int) {
	y.year = year
}

func (y *YearlyInvestment) GetInvestments() [12]*Investment {
	return y.investments
}

func (y *YearlyInvestment) GetMonthlyInvestment(month common.Month) *Investment {
	return y.investments[month]
}

func (y *YearlyInvestment) SetInvestments(investments [12]*Investment) {
	y.investments = investments
}

func (y *YearlyInvestment) UpdateInvestment(month common.Month, investment *Investment) {
	y.investments[month] = investment
}


