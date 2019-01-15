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
	tape          []symbol.Symbol
}

func NewTuringMachine(initialStates set.Set, finalStates set.Set, transactions set.Set, actualState state.State, firstSymbol symbol.Symbol) TuringMachine {
	tm := &turingMachine{}
	tm.initialStates = initialStates
	tm.finalStates = finalStates
	tm.transactions = transactions
	tm.actualState = actualState
	tm.headPointer = 0
	tm.tape = make([]symbol.Symbol, 0)
	tm.tape = append(tm.tape, firstSymbol)
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
	for _, t := range tm.transactions.Iterator() {
		tm.Execute(t.(Transaction))
		return tm.actualState
	}
	return tm.actualState
}

func (tm *turingMachine) Execute(t Transaction) state.State {
	if t.Validate(tm) {
		switch t.GetAction() {
		case "P":
			tm.GetActualSymbol().Print()
		case "E":
			tm.GetActualSymbol().Erase()
		default:
			tm.GetActualSymbol().None()
		}

	}
	tm.actualState = t.GetNewState()
	return tm.actualState
}

func (tm *turingMachine) GetActualSymbol() symbol.Symbol {
	return tm.tape[tm.headPointer]
}

func (tm *turingMachine) GetActualState() state.State {
	return tm.actualState
}
