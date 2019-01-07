package transaction

import (
	turingMachine "github.com/made2591/go-tm/turing/machine"
	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

type Transaction interface {
	Validate(m turingMachine.TuringMachine) bool
	Simulate() state.State
	Execute() state.State
}

type transaction struct {
	ist state.State
	isy symbol.Symbol
	fst state.State
	fsy symbol.Symbol
	tst int8
}

func NewTransaction(ist state.State, isy symbol.Symbol, fst state.State, fsy symbol.Symbol, tst int8) Transaction {
	t := &transaction{}
	t.ist = ist
	t.isy = isy
	t.fst = fst
	t.fsy = fsy
	t.tst = tst
	return t
}

func (t *transaction) Validate(m turingMachine.TuringMachine) bool {
	return false
}

func (t *transaction) Simulate() state.State {
	return t.fst
}

func (t *transaction) Execute() state.State {
	return t.fst
}
