package machine

import (
	"strings"
	"testing"

	set "github.com/made2591/go-tm/set"
	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

func TestNewTuringMachine(t *testing.T) {

	iss := set.NewSet()
	s := state.NewState("A")
	iss.Add(s)

	fss := set.NewSet()
	trs := set.NewSet()

	tr := NewTransaction(state.NewState("B"), symbol.NewSymbol(uint8(4)), state.NewState("C"), symbol.NewSymbol(uint8(6)), "P")
	trs.Add(tr)

	as := s

	tm := NewTuringMachine(iss, fss, trs, as, symbol.NewSymbol(uint8(4)))
	if tm == nil {
		t.Errorf("NewTuringMachine() was incorrect, got: nil")
	}

}

func TestStep(t *testing.T) {

	iss := set.NewSet()
	fss := set.NewSet()
	trs := set.NewSet()

	tr0 := NewTransaction(state.NewInitialState(), symbol.NewSymbol(uint8(0)), state.NewState("A"), symbol.NewSymbol(uint8(0)), "N")
	tr1 := NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(0)), state.NewState("B"), symbol.NewSymbol(uint8(1)), "R")
	tr2 := NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(1)), state.NewState("B"), symbol.NewSymbol(uint8(1)), "L")
	tr3 := NewTransaction(state.NewState("B"), symbol.NewSymbol(uint8(0)), state.NewState("A"), symbol.NewSymbol(uint8(1)), "L")
	tr4 := NewTransaction(state.NewState("B"), symbol.NewSymbol(uint8(1)), state.NewFinalState(), symbol.NewSymbol(uint8(1)), "R")
	trs.Add(tr0)
	trs.Add(tr1)
	trs.Add(tr2)
	trs.Add(tr3)
	trs.Add(tr4)

	tm := NewTuringMachine(iss, fss, trs, state.NewState(state.INITIAL), symbol.NewSymbol(uint8(0)))
	if s := tm.Step(); !strings.EqualFold(s.GetIdentifier().(string), "A") {
		t.Errorf("Step() was incorrect, got: %s wanted: %s", s.GetIdentifier().(string), "A")
	}

}

func TestComputed(t *testing.T) {

	tm := NewTuringMachine(nil, nil, nil, state.NewState(state.FINAL), symbol.NewSymbol(uint8(0)))
	if !tm.Computed() {
		t.Errorf("Computed() was incorrect, got: %t, want: %t.", false, true)
	}

}

func TestRun(t *testing.T) {

	iss := set.NewSet()
	fss := set.NewSet()
	trs := set.NewSet()

	tr0 := NewTransaction(state.NewInitialState(), symbol.NewSymbol(uint8(0)), state.NewState("A"), symbol.NewSymbol(uint8(0)), "N")
	tr1 := NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(0)), state.NewState("B"), symbol.NewSymbol(uint8(1)), "R")
	tr2 := NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(1)), state.NewState("B"), symbol.NewSymbol(uint8(1)), "L")
	tr3 := NewTransaction(state.NewState("B"), symbol.NewSymbol(uint8(0)), state.NewState("A"), symbol.NewSymbol(uint8(1)), "L")
	tr4 := NewTransaction(state.NewState("B"), symbol.NewSymbol(uint8(1)), state.NewFinalState(), symbol.NewSymbol(uint8(1)), "R")
	trs.Add(tr0)
	trs.Add(tr1)
	trs.Add(tr2)
	trs.Add(tr3)
	trs.Add(tr4)

	tm := NewTuringMachine(iss, fss, trs, state.NewState(state.INITIAL), symbol.NewSymbol(uint8(0)))
	tm.Run()
	if s := tm.GetActualState(); !strings.EqualFold(s.GetIdentifier().(string), state.FINAL) {
		t.Errorf("Run() was incorrect, got: %s wanted: %s", s.GetIdentifier().(string), state.FINAL)
	}

}

func TestGetActualSymbol(t *testing.T) {

	tm := NewTuringMachine(nil, nil, nil, state.NewState(state.INITIAL), symbol.NewSymbol(uint8(0)))
	if c := tm.GetActualSymbol(); c.GetValue() != uint8(0) {
		t.Errorf("GetActualSymbol() was incorrect, got: %d wanted: %d", c.GetValue(), uint8(0))
	}

}

func TestGetActualState(t *testing.T) {

	tm := NewTuringMachine(nil, nil, nil, state.NewState(state.INITIAL), symbol.NewSymbol(uint8(0)))
	if c := tm.GetActualState(); !strings.EqualFold(c.GetIdentifier().(string), state.INITIAL) {
		t.Errorf("GetActualState() was incorrect, got: %s wanted: %s", c.GetIdentifier().(string), state.INITIAL)
	}

}

func TestMoveHeadPointer(t *testing.T) {

	tm := NewTuringMachine(nil, nil, nil, state.NewState(state.INITIAL), symbol.NewSymbol(uint8(0)))
	i := tm.MoveHeadPointer(symbol.NewSymbol(uint8(0)), "N")
	if i != 0 {
		t.Errorf("Run() was incorrect, got: %d wanted: %d", i, 0)
	}

	i = tm.MoveHeadPointer(symbol.NewSymbol(uint8(0)), "R")
	if i != 1 {
		t.Errorf("Run() was incorrect, got: %d wanted: %d", i, 1)
	}

	i = tm.MoveHeadPointer(symbol.NewSymbol(uint8(0)), "L")
	if i != 0 {
		t.Errorf("Run() was incorrect, got: %d wanted: %d", i, 0)
	}

	i = tm.MoveHeadPointer(symbol.NewSymbol(uint8(0)), "L")
	if i != 0 {
		t.Errorf("Run() was incorrect, got: %d wanted: %d", i, 0)
	}

}
