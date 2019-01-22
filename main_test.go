package main

import (
	"testing"

	set "github.com/made2591/go-tm/set"
	transaction "github.com/made2591/go-tm/turing/machine"
	turingMachine "github.com/made2591/go-tm/turing/machine"
	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

func TestMain(t *testing.T) {

	iss := set.NewSet()
	fss := set.NewSet()
	trs := set.NewSet()

	tr0 := transaction.NewTransaction(state.NewInitialState(), symbol.NewSymbol(uint8(0)), state.NewState("A"), symbol.NewSymbol(uint8(0)), "N")
	tr1 := transaction.NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(0)), state.NewState("B"), symbol.NewSymbol(uint8(1)), "R")
	tr2 := transaction.NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(1)), state.NewState("B"), symbol.NewSymbol(uint8(1)), "L")
	tr3 := transaction.NewTransaction(state.NewState("B"), symbol.NewSymbol(uint8(0)), state.NewState("A"), symbol.NewSymbol(uint8(1)), "L")
	tr4 := transaction.NewTransaction(state.NewState("B"), symbol.NewSymbol(uint8(1)), state.NewFinalState(), symbol.NewSymbol(uint8(1)), "R")
	trs.Add(tr0)
	trs.Add(tr1)
	trs.Add(tr2)
	trs.Add(tr3)
	trs.Add(tr4)

	tm := turingMachine.NewTuringMachine(iss, fss, trs, state.NewState(state.INITIAL), symbol.NewSymbol(uint8(0)))
	tm.Run()

	if tm == nil {
		t.Errorf("Main was incorrect, got: nil")
	}

}
