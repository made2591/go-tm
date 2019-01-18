package machine

import (
	"fmt"
	"strings"

	set "github.com/made2591/go-tm/set"
	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

const EmptyTapeLength = 5

// TuringMachine interface
type TuringMachine interface {
	Step() state.State
	Computed() bool
	Run()
	GetActualSymbol() symbol.Symbol
	GetActualState() state.State
	MoveHeadPointer(s symbol.Symbol, m string) int
	Print() string
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

// Run() Execute a TuringMachine step until it move to a final State
func (tm *turingMachine) Run() {

	for !tm.Computed() {
		tm.Step()
	}

}

// Computed() Check if a TuringMachine reached a final State
func (tm *turingMachine) Computed() bool {

	return tm.actualState.IsFinal()

}

// Step() Loop over transactions defined in the machine, take
// without any guarantee about order the first valid transaction
// and execute it over the Turing Machine.
func (tm *turingMachine) Step() state.State {

	for _, t := range tm.transactions.Iterator() {
		if t.(Transaction).Validate(tm) {
			return tm.Execute(t.(Transaction))
		}
	}
	return tm.actualState

}

// Execute() Apply a given transaction to the Turing Machine
// and return the new actual state.
func (tm *turingMachine) Execute(t Transaction) state.State {

	// fmt.Println(t.Print())
	tm.MoveHeadPointer(t.GetSymbolWritten(), t.GetMoveTape())
	tm.actualState = t.GetNewState()
	return tm.actualState

}

// GetActualSymbol() Get the Symbol pointed by the head pointer on the tape
func (tm *turingMachine) GetActualSymbol() symbol.Symbol {

	return tm.tape[tm.headPointer]

}

// GetActualState() Get the actual State of the Turing machine
func (tm *turingMachine) GetActualState() state.State {

	return tm.actualState

}

// MoveHeadPointer() Write the Symbol in the actual head pointer position
// and move the head pointer. Adjustment of the tape is done automatically
// if the head pointer go out of bound.
func (tm *turingMachine) MoveHeadPointer(s symbol.Symbol, m string) int {

	if strings.EqualFold(m, "R") {
		tm.tape[tm.headPointer] = s
		tm.tape = append(tm.tape, symbol.NewSymbol())
		tm.headPointer++
	} else {
		if strings.EqualFold(m, "L") {
			tm.tape[tm.headPointer] = s
			tm.headPointer--
			if tm.headPointer < 0 {
				tm.tape = append([]symbol.Symbol{symbol.NewSymbol()}, tm.tape...)
				tm.headPointer = 0
			}
		}
	}
	return tm.headPointer

}

// Print() Return a string representation of the Turing machine
func (tm *turingMachine) Print() string {

	s := fmt.Sprintf("Actual state of TM:    %d", tm.actualState.GetIdentifier())
	s = strings.Join([]string{s, fmt.Sprintf("Actual symbol on head: %d", tm.tape[tm.headPointer].GetValue())}, "\n")
	s = strings.Join([]string{s, fmt.Sprintf("Actual tape pointer:   %d", tm.headPointer)}, "\n")
	s = strings.Join([]string{s, ""}, "\n")
	for i := 0; i < EmptyTapeLength/2; i++ {
		s = strings.Join([]string{s, fmt.Sprintf("%d", symbol.BLANK)}, "")
	}
	for _, sy := range tm.tape {
		s = strings.Join([]string{s, fmt.Sprintf("%d", sy.GetValue())}, "")
	}
	for i := 0; i < EmptyTapeLength/2; i++ {
		s = strings.Join([]string{s, fmt.Sprintf("%d", symbol.BLANK)}, "")
	}
	s = strings.Join([]string{s, ""}, "\n")
	for i := 0; i < len(tm.tape)+EmptyTapeLength; i++ {
		if i == tm.headPointer+(EmptyTapeLength/2) {
			s = strings.Join([]string{s, "^"}, "")
		} else {
			s = strings.Join([]string{s, " "}, "")
		}
	}
	return s

}
