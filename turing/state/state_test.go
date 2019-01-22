package state

import (
	"strings"
	"testing"
)

func TestNewState(t *testing.T) {

	s := NewState("A")

	if s == nil {
		t.Errorf("NewState was() incorrect, got: nil")
	}

}

func TestNewInitialState(t *testing.T) {

	s := NewInitialState()

	if s == nil {
		t.Errorf("NewInitialState was incorrect, got: nil")
	}

	if i := s.GetIdentifier(); !strings.EqualFold(i.(string), INITIAL) {
		t.Errorf("NewInitialState was incorrect, got: %d, want: %+v.", i, INITIAL)
	}

}

func TestNewFinalState(t *testing.T) {

	s := NewFinalState()

	if s == nil {
		t.Errorf("NewFinalState() was incorrect, got: nil")
	}

	if i := s.GetIdentifier(); !strings.EqualFold(i.(string), FINAL) {
		t.Errorf("NewFinalState() was incorrect, got: %d, want: %+v.", i, FINAL)
	}

}

func TestGetIdentifier(t *testing.T) {

	s := NewState("A")

	if s == nil {
		t.Errorf("GetIdentifier was incorrect, got: nil")
	}

	if i := s.GetIdentifier(); !strings.EqualFold(i.(string), "A") {
		t.Errorf("NewState was() incorrect, got: %s, want: %s", i, "A")
	}

}

func TestIsInitial(t *testing.T) {

	state := NewState(INITIAL)

	if state == nil {
		t.Errorf("IsInitial() was incorrect, got: nil")
	}

	if !state.IsInitial() {
		t.Errorf("IsInitial() was incorrect, got: %t, want: %t.", false, true)
	}

}

func TestIsFinal(t *testing.T) {

	s := NewState(FINAL)

	if s == nil {
		t.Errorf("IsFinal() was incorrect, got: nil")
	}

	if !s.IsFinal() {
		t.Errorf("IsFinal() was incorrect, got: %t, want: %t.", false, true)
	}

}

func TestEqual(t *testing.T) {

	s1 := NewState("A")
	s2 := NewState("A")
	if !s1.Equal(s2) {
		t.Errorf("Equal() was incorrect, got: %t, want: %t.", false, true)
	}

	s1 = NewState("A")
	s2 = NewState("B")
	if s1.Equal(s2) {
		t.Errorf("Equal() was incorrect, got: %t, want: %t.", true, false)
	}

}
