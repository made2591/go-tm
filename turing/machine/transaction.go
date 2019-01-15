package machine

import (
	"strings"

	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

var ACTIONS = [...]string{"P", "E", "N"}

type Transaction interface {
	Validate(m TuringMachine) bool
	Simulate() (state.State, symbol.Symbol, string)
	GetCurrentState() state.State
	GetSymbolScanned() symbol.Symbol
	GetNewState() state.State
	GetSymbolWritten() symbol.Symbol
	GetAction() string
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
			return m.GetActualSymbol().Equal(t.symbolScanned)
		}
	}
	return false
}

func (t *transaction) Simulate() (state.State, symbol.Symbol, string) {
	return t.newState, t.symbolWritten, t.action
}

func (t *transaction) GetCurrentState() state.State {
	return t.currentState
}

func (t *transaction) GetSymbolScanned() symbol.Symbol {
	return t.symbolScanned
}

func (t *transaction) GetNewState() state.State {
	return t.newState
}

func (t *transaction) GetSymbolWritten() symbol.Symbol {
	return t.symbolWritten
}

func (t *transaction) GetAction() string {
	return t.action
}
