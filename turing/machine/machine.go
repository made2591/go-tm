package machine

import (
	set "github.com/made2591/go-tm/set"
	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

type TuringMachine interface {
	Run()
	Step() state.State
	Computed() bool
	GetActualSymbol() symbol.Symbol
}

type turingMachine struct {
	initialStates set.Set
	finalStates   set.Set
	transactions  set.Set
	actualState   state.State
	headPointer   int
	tape          []*symbol.Symbol
}

func NewTuringMachine(initialStates set.Set, finalStates set.Set, transactions set.Set, actualState state.State) TuringMachine {
	tm := &turingMachine{}
	tm.initialStates = initialStates
	tm.finalStates = finalStates
	tm.transactions = transactions
	tm.actualState = initialStates.GetOne()
	tm.headPointer = 0
	tm.tape = make([]*symbol.Symbol{})
	tm.tape = append(tm.tape, tm.actualState)
	return tm
}

func (tm *turingMachine) Run() {
	for !tm.Computed() {
		tm.Step()
	}
}

func (tm *turingMachine) Computed() bool {
	return tm.actualState.IsFinal()
}

func (tm *turingMachine) Step() state.State {
	for t := range tm.transactions.Iterator() {
		tm.Execute(t)
		return tm.actualState
	}
	return tm.actualState
}

func (tm *turingMachine) Execute(m TuringMachine) state.State {
	if t.Validate(m) {
		switch t.action {
		case "P":
			m.GetActualSymbol().Print()
		case "E":
			m.GetActualSymbol().Erase()
		default:
			m.GetActualSymbol().None()
		}

	}
	return m.GetActualState()
}

func (tm *turingMachine) GetActualSymbol() symbol.Symbol {
	return tm.tape[headPointer]
}

func (tm *turingMachine) GetActualState() state.State {
	return tm.actualState
}
