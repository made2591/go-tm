package machine

import (
	"strings"
	"testing"

	set "github.com/made2591/go-tm/set"
	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

func TestNewTransaction(t *testing.T) {

	tr := NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(6)), state.NewState("A"), symbol.NewSymbol(uint8(6)), "R")
	if tr == nil {
		t.Errorf("NewTransaction() was incorrect, got: nil")
	}

}

func TestValidate(t *testing.T) {

	iss := set.NewSet()
	s := state.NewState("A")
	iss.Add(s)

	fss := set.NewSet()
	trs := set.NewSet()

	tr := NewTransaction(s, symbol.NewSymbol(uint8(4)), state.NewState("B"), symbol.NewSymbol(uint8(6)), "R")
	trs.Add(tr)

	as := s

	tm := NewTuringMachine(iss, fss, trs, as, symbol.NewSymbol(uint8(4)))
	if v := tr.Validate(tm); !v {
		t.Errorf("Validate() was incorrect, got: %t, want: %t.", v, true)
	}

}

func TestSimulate(t *testing.T) {

	tr := NewTransaction(nil, nil, state.NewState("B"), symbol.NewSymbol(uint8(6)), "R")
	ns, sw, mt := tr.Simulate()

	if v := ns.GetIdentifier(); !strings.EqualFold(v.(string), "B") {
		t.Errorf("Simulate() was incorrect, got: %s, want: %s.", v, "B")
	}

	if v := sw.GetValue(); v != uint8(6) {
		t.Errorf("Simulate() was incorrect, got: %d, want: %d.", v, uint8(6))
	}

	if !strings.EqualFold(mt, "R") {
		t.Errorf("Simulate() was incorrect, got: %s, want: %s.", mt, "R")
	}

}

func TestGetCurrentState(t *testing.T) {

	tr := NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(3)), state.NewState("B"), symbol.NewSymbol(uint8(9)), "R")
	if v := tr.GetCurrentState().GetIdentifier(); !strings.EqualFold(v.(string), "A") {
		t.Errorf("GetCurrentState() was incorrect, got: %s, want: %s.", v, "A")
	}

}

func TestGetSymbolScanned(t *testing.T) {

	tr := NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(3)), state.NewState("B"), symbol.NewSymbol(uint8(9)), "R")
	if v := tr.GetSymbolScanned().GetValue(); v != uint8(3) {
		t.Errorf("GetSymbolScanned() was incorrect, got: %d, want: %d.", v, uint8(3))
	}

}

func TestGetNewState(t *testing.T) {

	tr := NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(3)), state.NewState("B"), symbol.NewSymbol(uint8(9)), "R")
	if v := tr.GetNewState().GetIdentifier(); !strings.EqualFold(v.(string), "B") {
		t.Errorf("NewState() was incorrect, got: %s, want: %s.", v, "B")
	}

}

func TestGetSymbolWritten(t *testing.T) {

	tr := NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(3)), state.NewState("B"), symbol.NewSymbol(uint8(9)), "R")
	if v := tr.GetSymbolWritten().GetValue(); v != uint8(9) {
		t.Errorf("GetSymbolWritten() was incorrect, got: %d, want: %d.", v, uint8(9))
	}

}

func TestGetMoveTape(t *testing.T) {

	tr := NewTransaction(state.NewState("A"), symbol.NewSymbol(uint8(3)), state.NewState("B"), symbol.NewSymbol(uint8(9)), "R")
	if mt := tr.GetMoveTape(); !strings.EqualFold(mt, "R") {
		t.Errorf("GetMoveTape was incorrect, got: %s, want: %s.", mt, "R")
	}

}
