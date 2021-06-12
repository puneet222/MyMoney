package portfolio

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Investment struct {
	Equity float64
	Debt float64
	Gold float64
}

type Change Investment

type Allocation Investment

type SIP Investment

func (i *Investment) AddSIP(sip *SIP) {
	i.Equity += sip.Equity
	i.Debt += sip.Debt
	i.Gold += sip.Gold
}

func (i *Investment) UpdateInvestment(change Change) {
	i.Equity = i.Equity + (change.Equity/100)*i.Equity
	i.Debt = i.Debt + (change.Debt/100)*i.Debt
	i.Gold = i.Gold + (change.Gold/100)*i.Gold
}

func (i *Investment) GetAllocation() *Allocation {
	equityAllocation := (i.Equity*100)/i.GetTotalInvestment()
	debtAllocation := (i.Debt*100)/i.GetTotalInvestment()
	goldAllocation := (i.Gold*100)/i.GetTotalInvestment()
	return &Allocation{equityAllocation, debtAllocation, goldAllocation}
}

func (i *Investment) GetTotalInvestment() float64 {
	return i.Equity + i.Debt + i.Gold
}

func (i *Investment) RoundOffInvestment() {
	i.Equity = float64(int(i.Equity))
	i.Debt = float64(int(i.Debt))
	i.Gold = float64(int(i.Gold))
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
	icopy := Investment{}
	err = decoder.Decode(&icopy)
	if err != nil {
		panic("error while deep copying (decoding) investment object")
	}
	return &icopy
}

func (i *Investment) String() string {
	return fmt.Sprintf("Equity %f | Debt %f | Gold %f\n", i.Equity, i.Debt, i.Gold)
}

func (s *SIP) String() string {
	return fmt.Sprintf("SIPs for Equity: %d Debt: %d Gold: %d\n", int(s.Equity), int(s.Debt), int(s.Gold))
}

func (a *Allocation) String() string {
	return fmt.Sprintf("Allocations Equity: %d%% Debt: %d%% Gold: %d%%\n", int(a.Equity), int(a.Debt), int(a.Gold))
}

func (i *Investment) Output() string {
	return fmt.Sprintf("%d %d %d", int(i.Equity), int(i.Debt), int(i.Gold))
}
