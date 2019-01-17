package machine

import (
	"testing"

	set "github.com/made2591/go-tm/set"
	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

func TestNewTuringMachine(t *testing.T) {

	iss := set.NewSet()
	s := state.NewState(uint8(4))
	iss.Add(s)

	fss := set.NewSet()
	trs := set.NewSet()

	tr := NewTransaction(state.NewState(uint8(3)), symbol.NewSymbol(uint8(4)), state.NewState(uint8(5)), symbol.NewSymbol(uint8(6)), "P")
	trs.Add(tr)

	as := s

	tm := NewTuringMachine(iss, fss, trs, as, symbol.NewSymbol(uint8(4)))
	if tm == nil {
		t.Errorf("NewTuringMachine was incorrect, got: nil")
	}

}

func TestRun(t *testing.T) {

	iss := set.NewSet()
	s := state.NewState(uint8(4))
	iss.Add(s)

	fss := set.NewSet()
	trs := set.NewSet()

	tr1 := NewTransaction(state.NewState(uint8(0)), symbol.NewSymbol(uint8(0)), state.NewState(uint8(1)), symbol.NewSymbol(uint8(6)), "P")
	tr2 := NewTransaction(state.NewState(uint8(0)), symbol.NewSymbol(uint8(1)), state.NewState(uint8(1)), symbol.NewSymbol(uint8(6)), "P")
	tr3 := NewTransaction(state.NewState(uint8(1)), symbol.NewSymbol(uint8(0)), state.NewState(uint8(1)), symbol.NewSymbol(uint8(6)), "P")
	tr3 := NewTransaction(state.NewState(uint8(1)), symbol.NewSymbol(uint8(1)), state.NewState(state.FINAL), symbol.NewSymbol(uint8(6)), "P")
	trs.Add(tr1)
	trs.Add(tr2)
	trs.Add(tr3)
	A	B
0	1RB	1LA
1	1LB	1RH

	as := s

	tm := NewTuringMachine(iss, fss, trs, as, symbol.NewSymbol(uint8(0)))
	if tm == nil {
		t.Errorf("NewTuringMachine was incorrect, got: nil")
	}

}

func TestStep(t *testing.T) {

}

func TestComputed(t *testing.T) {

}

func TestGetActualSymbol(t *testing.T) {

}

func TestGetActualState(t *testing.T) {

}
