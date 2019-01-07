package machine

import (
	set "github.com/made2591/go-tm/set"
	state "github.com/made2591/go-tm/turing/state"
)

type TuringMachine interface {
	Run() bool
	Step() bool
	Computed() bool
}

type turingMachine struct {
	initialStates set.Set
	finalStates   set.Set
	transactions  set.Set
	actualState   state.State
	headPointer   int
}

func NewTuringMachine(initialStates set.Set, finalStates set.Set, transactions set.Set, actualState state.State) TuringMachine {
	tm := &turingMachine{}
	tm.initialStates = initialStates
	tm.finalStates = finalStates
	tm.transactions = transactions
	tm.actualState = nil
	tm.headPointer = 0
	return tm
}

func (tm *turingMachine) Run() bool {
	for !tm.Computed() {
		tm.Step()
	}
	return true
}

func (tm *turingMachine) Computed() bool {
	return tm.actualState.IsFinal()
}

func (tm *turingMachine) Step() bool {
	for _, transaction := range tm.transactions.Iterator() {

	}
}
