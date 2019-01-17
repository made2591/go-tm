package machine

import (
	"strings"

	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

var ACTIONS = [...]string{"R", "L", "N"}

// Transaction interface
type Transaction interface {
	Validate(m TuringMachine) bool
	Simulate() (state.State, symbol.Symbol, string)
	GetCurrentState() state.State
	GetSymbolScanned() symbol.Symbol
	GetNewState() state.State
	GetSymbolWritten() symbol.Symbol
	GetMoveTape() string
}

// transaction struct
type transaction struct {
	currentState  state.State
	symbolScanned symbol.Symbol
	newState      state.State
	symbolWritten symbol.Symbol
	moveTape      string
}

// NewTransaction() Create a new Transaction with given
// currentState, symbolScanned, symbolWritten and moveTape action
func NewTransaction(currentState state.State, symbolScanned symbol.Symbol, newState state.State, symbolWritten symbol.Symbol, moveTape string) Transaction {

	t := &transaction{}
	t.currentState = currentState
	t.symbolScanned = symbolScanned
	t.newState = newState
	t.symbolWritten = symbolWritten
	t.moveTape = moveTape

	return t

}

// Validate(m TuringMachine) Validate the transaction t with respect
// to the TuringMachine m. A Transaction to be considered valid
// need to be defined with a valid ACTIONS, the current state of the
// TuringMachine must be equal to the state in which the transaction
// is activated and the head of the TuringMachine must point to the
// the same symbol scanned by the transaction. If all of this three
// condition are verified, Validate(m TuringMachine) over t returns
// true; it returns otherwise
func (t *transaction) Validate(m TuringMachine) bool {

	// check if moveTape action is allowed
	for _, a := range ACTIONS {
		if strings.EqualFold(a, t.moveTape) {
			// check if actual state and scanned symbol match with transaction
			return t.currentState.Equal(m.GetActualState()) && t.symbolScanned.Equal(m.GetActualSymbol())
		}
	}

	return false

}

// Simulate() Return the state in which the transaction
// will bring the TuringMachine, the symbol written and the moveTape action
// taken over the symbol pointed by the TuringMachine's head. Pay attention
// that this method DOES NOT accept any TuringMachine, thus it doesn't
// verify that the transaction is doable over any specific TuringMachine.
// It returns a State, a Symbol and an moveTape action
func (t *transaction) Simulate() (state.State, symbol.Symbol, string) {

	return t.newState, t.symbolWritten, t.moveTape

}

// GetCurrentState() Return the state needed to activate the transaction
func (t *transaction) GetCurrentState() state.State {

	return t.currentState

}

// GetSymbolScanned() Return the Symbol scanned
func (t *transaction) GetSymbolScanned() symbol.Symbol {

	return t.symbolScanned

}

// GetNewState() Return the new State in which the
// Transaction bring the TuringMachine in which is applied
func (t *transaction) GetNewState() state.State {

	return t.newState

}

// GetSymbolWritten() Return the new Symbol written by the
// Transaction bring the TuringMachine in which is applied
func (t *transaction) GetSymbolWritten() symbol.Symbol {

	return t.symbolWritten

}

// GetMoveTape() Return the moveTape action done by the
// Transaction over the Symbol scanned
func (t *transaction) GetMoveTape() string {

	return t.moveTape

}
