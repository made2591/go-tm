package turingMachine

import (
	"github.com/made2591/go-tm/set"
	"github.com/made2591/go-tm/state"
)

type TuringMachine interface {
	Run()
	Step()
}

type turingMachine struct {
	is  set.Set
	fs  set.Set
	act state.State
	tst int
}

func NewTuringMachine(is set.Set, fs set.Set) State {
	tm := &turingMachine{}
	tm.is = is
	tm.fs = fs
	tm.act = is.
	tm.tst = 0
	return s
}
