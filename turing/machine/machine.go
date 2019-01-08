package machine

import (
	set "github.com/made2591/go-tm/set"
	state "github.com/made2591/go-tm/turing/state"
)

type TuringMachine interface {
	Run()
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

func (tm *turingMachine) Run() {
	for !tm.Computed() {
		tm.Step()
	}
}

func (tm *turingMachine) Computed() bool {
	return tm.actualState.IsFinal()
}

func (tm *turingMachine) Step() state.State {
	t := NewTransaction(nil, nil, nil, nil, 0)
	for _, t = range tm.transactions.Iterator() {
		return t.Execute()
	}
	return nil
}
