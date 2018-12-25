package state

import (
	"testing"
)

func TestNewInitialState(t *testing.T) {
	state := NewInitialState()
	if v := state.GetValue(); v != INITIAL {
		t.Errorf("NewInitialState was incorrect, got: %d, want: %+v.", v, INITIAL)
	}
}
