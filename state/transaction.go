package state

import (
	symbol "github.com/made2591/go-tm/symbol"
	turingMachine "github.com/made2591/go-tm/turingMachine"
)

type Transaction interface {
	Validate(m turingMachine.TuringMachine) bool
	Simulate() State
	Execute() State
}

type transaction struct {
	ist State
	isy symbol.Symbol
	fst State
	fsy symbol.Symbol
	tst int8
}

func NewTransaction(ist State, isy symbol.Symbol, fst State, fsy symbol.Symbol, tst int8) Transaction {
	t := &transaction{}
	t.ist = ist
	t.isy = isy
	t.fst = fst
	t.fsy = fsy
	t.tst = tst
	return t
}

func (t *transaction) Validate(m turingMachine.TuringMachine) bool {
	return true
}

func (t *transaction) Simulate() State {
	return t.fst
}

func (t *transaction) Execute() State {
	return t.fst
}
