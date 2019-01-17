package machine

import (
	set "github.com/made2591/go-tm/set"
	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

// TuringMachine interface
type TuringMachine interface {
	Run()
	Step() state.State
	Computed() bool
	GetActualSymbol() symbol.Symbol
	GetActualState() state.State
	MoveHeadPointer() int64
}

// turingMachine struct
type turingMachine struct {
	initialStates set.Set
	finalStates   set.Set
	transactions  set.Set
	actualState   state.State
	headPointer   int
	tape          []symbol.Symbol
}

// NewTuringMachine() Return a new TuringMachine
func NewTuringMachine(iss set.Set, fss set.Set, trs set.Set, as state.State, fs symbol.Symbol) TuringMachine {

	tm := &turingMachine{}
	tm.initialStates = iss
	tm.finalStates = fss
	tm.transactions = trs
	tm.actualState = as
	tm.headPointer = 0
	tm.tape = make([]symbol.Symbol, 0)
	tm.tape = append(tm.tape, fs)

	return tm

}

// Run() Execute a TuringMachine until it move to a final State
func (tm *turingMachine) Run() {

	for !tm.Computed() {
		tm.Step()
	}

}

// Computed() Check if a TuringMachine reached a final State
func (tm *turingMachine) Computed() bool {

	return tm.actualState.IsFinal()

}

// Step() Check if a TuringMachine reached a final State
func (tm *turingMachine) Step() state.State {

	for _, t := range tm.transactions.Iterator() {
		tm.Execute(t.(Transaction))
		return tm.actualState
	}

	return tm.actualState

}

// Execute() Check if a TuringMachine reached a final State
func (tm *turingMachine) Execute(t Transaction) state.State {

	if t.Validate(tm) {
		tm.MoveHeadPointer(t.GetMoveTape())
		tm.actualState = t.GetNewState()
	}
	return tm.actualState

}

// GetActualSymbol() Check if a TuringMachine reached a final State
func (tm *turingMachine) GetActualSymbol() symbol.Symbol {

	return tm.tape[tm.headPointer]

}

// GetActualState() Check if a TuringMachine reached a final State
func (tm *turingMachine) GetActualState() state.State {

	return tm.actualState

}

// MoveHeadPointer() Move the head pointer to
func (tm *turingMachine) MoveHeadPointer() int64 {

	//TODO

}
