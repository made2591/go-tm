package state

import (
	"testing"

	"github.com/made2591/go-tm/set"
)

func TestNewInitialState(t *testing.T) {
	state := NewInitialState()
	if state == nil {
		t.Errorf("NewInitialState was incorrect, got: nil")
	}
	if v := state.GetValue(); v != INITIAL {
		t.Errorf("NewInitialState was incorrect, got: %d, want: %+v.", v, INITIAL)
	}
}

func TestNewState(t *testing.T) {
	state := NewState(uint8(1), set.NewSet())
	if state == nil {
		t.Errorf("NewState was incorrect, got: nil")
	}
	if v := state.GetValue(); v != uint8(1) {
		t.Errorf("NewState was incorrect, got: %d, want: %+v.", v, uint8(1))
	}
}

func TestIsInitial(t *testing.T) {
	state := NewState(INITIAL, set.NewSet())
	if state == nil {
		t.Errorf("IsInitial was incorrect, got: nil")
	}
	if v := state.IsInitial(); v != true {
		t.Errorf("IsInitial was incorrect, got: %t, want: %t.", v, true)
	}
}

func TestIsFinal(t *testing.T) {
	state := NewState(FINAL, set.NewSet())
	if state == nil {
		t.Errorf("IsFinal was incorrect, got: nil")
	}
	if v := state.IsFinal(); v != true {
		t.Errorf("IsFinal was incorrect, got: %t, want: %t.", v, true)
	}
}

func TestGetValue(t *testing.T) {
	state := NewState(uint8(1), set.NewSet())
	if state == nil {
		t.Errorf("GetValue was incorrect, got: nil")
	}
	if v := state.GetValue(); v != uint8(1) {
		t.Errorf("GetValue was incorrect, got: %d, want: %d.", v, uint8(1))
	}
}
