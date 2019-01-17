package state

import (
	"testing"
)

func TestNewInitialState(t *testing.T) {

	s := NewInitialState()

	if s == nil {
		t.Errorf("NewInitialState was incorrect, got: nil")
	}

	if i := s.GetIdentifier(); i != INITIAL {
		t.Errorf("NewInitialState was incorrect, got: %d, want: %+v.", i, INITIAL)
	}

}

func TestNewState(t *testing.T) {

	s := NewState(uint8(1))

	if s == nil {
		t.Errorf("NewState was() incorrect, got: nil")
	}

	if i := s.GetIdentifier(); i != uint8(1) {
		t.Errorf("NewState was() incorrect, got: %d, want: %+v.", i, uint8(1))
	}

}

func TestNewFinalState(t *testing.T) {

	s := NewFinalState()

	if s == nil {
		t.Errorf("NewFinalState() was incorrect, got: nil")
	}

	if i := s.GetIdentifier(); i != FINAL {
		t.Errorf("NewFinalState() was incorrect, got: %d, want: %+v.", i, FINAL)
	}

}

func TestIsInitial(t *testing.T) {

	state := NewState(INITIAL)

	if state == nil {
		t.Errorf("IsInitial() was incorrect, got: nil")
	}

	if i := state.IsInitial(); i != true {
		t.Errorf("IsInitial() was incorrect, got: %t, want: %t.", i, true)
	}

}

func TestIsFinal(t *testing.T) {

	s := NewState(FINAL)

	if s == nil {
		t.Errorf("IsFinal() was incorrect, got: nil")
	}

	if i := s.IsFinal(); i != true {
		t.Errorf("IsFinal() was incorrect, got: %t, want: %t.", i, true)
	}

}

func TestGetIdentifier(t *testing.T) {

	s := NewState(uint8(1))

	if s == nil {
		t.Errorf("GetIdentifier was incorrect, got: nil")
	}

	if i := s.GetIdentifier(); i != uint8(1) {
		t.Errorf("GetIdentifier was incorrect, got: %d, want: %d.", i, uint8(1))
	}

}

func TestEqual(t *testing.T) {

	s1 := NewState(uint8(1))
	s2 := NewState(uint8(1))
	if s1.Equal(s2) == false {
		t.Errorf("Equal() was incorrect, got: %t, want: %t.", false, true)
	}

	s1 = NewState(uint8(1))
	s2 = NewState(uint8(3))
	if s1.Equal(s2) == true {
		t.Errorf("Equal() was incorrect, got: %t, want: %t.", true, false)
	}

}
