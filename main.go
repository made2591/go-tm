package main

import (
	"github.com/made2591/go-tm/set"
	transaction "github.com/made2591/go-tm/turing/machine"
	turingMachine "github.com/made2591/go-tm/turing/machine"
	"github.com/made2591/go-tm/turing/state"
	"github.com/made2591/go-tm/turing/symbol"
)

func main() {

	iss := set.NewSet()
	fss := set.NewSet()
	trs := set.NewSet()

	tr2 := transaction.NewTransaction(state.NewState(uint8(0)), symbol.NewSymbol(uint8(1)), state.NewState(uint8(1)), symbol.NewSymbol(uint8(6)), "L")
	tr1 := transaction.NewTransaction(state.NewState(state.INITIAL), symbol.NewSymbol(uint8(0)), state.NewState(uint8(1)), symbol.NewSymbol(uint8(6)), "R")
	tr4 := transaction.NewTransaction(state.NewState(uint8(1)), symbol.NewSymbol(uint8(1)), state.NewState(state.FINAL), symbol.NewSymbol(uint8(6)), "L")
	tr3 := transaction.NewTransaction(state.NewState(uint8(1)), symbol.NewSymbol(uint8(0)), state.NewState(uint8(1)), symbol.NewSymbol(uint8(6)), "R")
	trs.Add(tr1)
	trs.Add(tr2)
	trs.Add(tr3)
	trs.Add(tr4)

	tm := turingMachine.NewTuringMachine(iss, fss, trs, state.NewState(state.INITIAL), symbol.NewSymbol(uint8(0)))
	tm.Run()

}
