package main

import (
	"bufio"
	"fmt"
	"os"

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

	tr1 := transaction.NewTransaction(state.NewState(state.INITIAL), symbol.NewSymbol(uint8(0)), state.NewState(uint8(8)), symbol.NewSymbol(uint8(1)), "R")
	tr2 := transaction.NewTransaction(state.NewState(state.INITIAL), symbol.NewSymbol(uint8(1)), state.NewState(uint8(8)), symbol.NewSymbol(uint8(1)), "L")
	tr4 := transaction.NewTransaction(state.NewState(uint8(8)), symbol.NewSymbol(uint8(0)), state.NewState(state.INITIAL), symbol.NewSymbol(uint8(1)), "L")
	tr3 := transaction.NewTransaction(state.NewState(uint8(8)), symbol.NewSymbol(uint8(1)), state.NewState(state.FINAL), symbol.NewSymbol(uint8(1)), "R")
	trs.Add(tr1)
	trs.Add(tr2)
	trs.Add(tr3)
	trs.Add(tr4)

	tm := turingMachine.NewTuringMachine(iss, fss, trs, state.NewState(state.INITIAL), symbol.NewSymbol(uint8(0)))
	for {
		tm.Step()
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		buf.ReadBytes('\n')
		if tm.Computed() {
			fmt.Println("Final state reached")
			break
		}
	}

}
