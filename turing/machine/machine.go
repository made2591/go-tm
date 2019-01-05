package machine

import (
	set "github.com/made2591/go-tm/set"
	state "github.com/made2591/go-tm/turing/state"
)

type TuringMachine interface {
	Run() bool
	Step() bool
}

type turingMachine struct {
	is  set.Set
	fs  set.Set
	act state.State
	pnt int
}

func NewTuringMachine(is set.Set, fs set.Set) TuringMachine {
	tm := &turingMachine{}
	tm.is = is
	tm.fs = fs
	tm.act = nil
	tm.pnt = 0
	return tm
}

func (tm *turingMachine) Run() bool {
	return true
}

func (tm *turingMachine) Step() bool {
	return true
}
