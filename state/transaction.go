package state

import (
	symbol "github.com/made2591/go-tm/symbol"
)

type Transaction interface {
	Validate() bool
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

func NewTransaction() Transaction {
	t := &transaction{}
	return t
}

func (t *transaction) Validate() bool {
	return true
}

func (t *transaction) Simulate() State {
	return &state{}
}

func (t *transaction) Execute() State {
	return &state{}
}
