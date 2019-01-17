package machine

import (
	"strings"
	"testing"

	set "github.com/made2591/go-tm/set"
	state "github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

func TestNewTransaction(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(6)), state.NewState(uint8(5)), symbol.NewSymbol(uint8(6)), "P")
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

	tr := NewTransaction(state.NewState(uint8(3)), symbol.NewSymbol(uint8(4)), state.NewState(uint8(5)), symbol.NewSymbol(uint8(6)), "P")
	trs.Add(tr)

	as := s

	tm := NewTuringMachine(iss, fss, trs, as, symbol.NewSymbol(uint8(4)))
	if v := tr.Validate(tm); v != true {
		t.Errorf("Validate was incorrect, got: %t, want: %t.", v, true)
	}

}

func TestSimulate(t *testing.T) {

	tr := NewTransaction(nil, nil, state.NewState(uint8(5)), symbol.NewSymbol(uint8(6)), "P")
	ns, sw, ac := tr.Simulate()

	if v := ns.GetIdentifier(); v != uint8(5) {
		t.Errorf("Simulate was incorrect, got: %d, want: %d.", v, uint8(5))
	}

	if v := sw.GetValue(); v != uint8(6) {
		t.Errorf("Simulate was incorrect, got: %d, want: %d.", v, uint8(6))
	}

	if strings.EqualFold(ac, "P") == false {
		t.Errorf("Simulate was incorrect, got: %s, want: %s.", ac, "P")
	}

}

func TestGetCurrentState(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(3)), state.NewState(uint8(2)), symbol.NewSymbol(uint8(9)), "P")
	if v := tr.GetCurrentState().GetIdentifier(); v != uint8(5) {
		t.Errorf("GetCurrentState was incorrect, got: %d, want: %d.", v, uint8(5))
	}

}

func TestGetSymbolScanned(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(3)), state.NewState(uint8(2)), symbol.NewSymbol(uint8(9)), "P")
	if v := tr.GetSymbolScanned().GetValue(); v != uint8(3) {
		t.Errorf("GetSymbolScanned was incorrect, got: %d, want: %d.", v, uint8(3))
	}

}

func TestGetNewState(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(3)), state.NewState(uint8(2)), symbol.NewSymbol(uint8(9)), "P")
	if v := tr.GetNewState().GetIdentifier(); v != uint8(2) {
		t.Errorf("GetNewState was incorrect, got: %d, want: %d.", v, uint8(2))
	}

}

func TestGetSymbolWritten(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(3)), state.NewState(uint8(2)), symbol.NewSymbol(uint8(9)), "P")
	if v := tr.GetSymbolWritten().GetValue(); v != uint8(9) {
		t.Errorf("GetSymbolWritten was incorrect, got: %d, want: %d.", v, uint8(9))
	}

}

func TestGetAction(t *testing.T) {

	tr := NewTransaction(state.NewState(uint8(5)), symbol.NewSymbol(uint8(3)), state.NewState(uint8(2)), symbol.NewSymbol(uint8(9)), "P")
	if ac := tr.GetAction(); strings.EqualFold(ac, "P") == false {
		t.Errorf("GetAction was incorrect, got: %s, want: %s.", ac, "P")
	}

}
