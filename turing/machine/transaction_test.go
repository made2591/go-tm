package machine

import (
	"strings"
	"testing"

	set "github.com/made2591/go-tm/set"
	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

func TestNewTransaction(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(6)), state.NewState(uint8(5)), symbol.NewSymbol(uint8(6)), "R")
	if tr == nil {
		t.Errorf("NewTransaction was incorrect, got: nil")
	}

}

func TestValidate(t *testing.T) {

	iss := set.NewSet()
	s := state.NewState(uint8(4))
	iss.Add(s)

	fss := set.NewSet()
	trs := set.NewSet()

	tr := NewTransaction(s, symbol.NewSymbol(uint8(4)), state.NewState(uint8(5)), symbol.NewSymbol(uint8(6)), "R")
	trs.Add(tr)

	as := s

	tm := NewTuringMachine(iss, fss, trs, as, symbol.NewSymbol(uint8(4)))
	if v := tr.Validate(tm); !v {
		t.Errorf("Validate was incorrect, got: %t, want: %t.", v, true)
	}

}

func TestSimulate(t *testing.T) {

	tr := NewTransaction(nil, nil, state.NewState(uint8(5)), symbol.NewSymbol(uint8(6)), "R")
	ns, sw, mt := tr.Simulate()

	if v := ns.GetIdentifier(); v != uint8(5) {
		t.Errorf("Simulate was incorrect, got: %d, want: %d.", v, uint8(5))
	}

	if v := sw.GetValue(); v != uint8(6) {
		t.Errorf("Simulate was incorrect, got: %d, want: %d.", v, uint8(6))
	}

	if !strings.EqualFold(mt, "R") {
		t.Errorf("Simulate was incorrect, got: %s, want: %s.", mt, "R")
	}

}

func TestGetCurrentState(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(3)), state.NewState(uint8(2)), symbol.NewSymbol(uint8(9)), "R")
	if v := tr.GetCurrentState().GetIdentifier(); v != uint8(5) {
		t.Errorf("GetCurrentState was incorrect, got: %d, want: %d.", v, uint8(5))
	}

}

func TestGetSymbolScanned(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(3)), state.NewState(uint8(2)), symbol.NewSymbol(uint8(9)), "R")
	if v := tr.GetSymbolScanned().GetValue(); v != uint8(3) {
		t.Errorf("GetSymbolScanned was incorrect, got: %d, want: %d.", v, uint8(3))
	}

}

func TestGetNewState(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(3)), state.NewState(uint8(2)), symbol.NewSymbol(uint8(9)), "R")
	if v := tr.GetNewState().GetIdentifier(); v != uint8(2) {
		t.Errorf("GetNewState was incorrect, got: %d, want: %d.", v, uint8(2))
	}

}

func TestGetSymbolWritten(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(3)), state.NewState(uint8(2)), symbol.NewSymbol(uint8(9)), "R")
	if v := tr.GetSymbolWritten().GetValue(); v != uint8(9) {
		t.Errorf("GetSymbolWritten was incorrect, got: %d, want: %d.", v, uint8(9))
	}

}

func TestGetMoveTape(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(3)), state.NewState(uint8(2)), symbol.NewSymbol(uint8(9)), "R")
	if mt := tr.GetMoveTape(); !strings.EqualFold(mt, "R") {
		t.Errorf("GetMoveTape was incorrect, got: %s, want: %s.", mt, "R")
	}

}
