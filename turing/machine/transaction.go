package machine

import (
	"strings"

	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

var ACTIONS = [...]string{"P", "E", "N"}

type Transaction interface {
	Validate(m TuringMachine) bool
	Simulate(m TuringMachine) state.State
	Execute(m TuringMachine) state.State
}

type transaction struct {
	currentState  state.State
	symbolScanned symbol.Symbol
	newState      state.State
	symbolWritten symbol.Symbol
	action        string
}

func NewTransaction(currentState state.State, symbolScanned symbol.Symbol, newState state.State, symbolWritten symbol.Symbol, action string) Transaction {
	t := &transaction{}
	t.currentState = currentState
	t.symbolScanned = symbolScanned
	t.newState = newState
	t.symbolWritten = symbolWritten
	t.action = action
	return t
}

func (t *transaction) Validate(m TuringMachine) bool {
	for _, a := range ACTIONS {
		if strings.EqualFold(a, t.action) {
			return true
		}
	}
	return false
}

func (t *transaction) Simulate(m TuringMachine) state.State, symbol.Symbol, string {
	return t.newState, t.symbolWritten, t.action
}
