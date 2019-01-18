package symbol

import "testing"

func TestNewSymbol(t *testing.T) {

	// test empty symbol creation
	s := NewSymbol()
	if s == nil {
		t.Errorf("NewSymbol was incorrect, got: nil, want: %v.", s)
	}

	// test not empty symbol creation
	v1 := uint8(1)
	s = NewSymbol(v1)
	if s == nil {
		t.Errorf("NewSymbol was incorrect, got: nil, want: %v.", s)
	}

	// test explicit BLANK symbol creation
	v2 := BLANK
	s = NewSymbol(v2)
	if s == nil {
		t.Errorf("NewSymbol was incorrect, got: nil, want: %v.", s)
	}

}

func TestGetValue(t *testing.T) {

	// test empty symbol creation
	s := NewSymbol()
	if v := s.GetValue(); v != uint8(BLANK) {
		t.Errorf("GetValue() was incorrect, got: %d, want: %d.", v, BLANK)
	}

	// test not empty symbol creation
	v1 := uint8(1)
	s = NewSymbol(v1)
	if v := s.GetValue(); v != v1 {
		t.Errorf("GetValue() was incorrect, got: %d, want: %d.", v, v1)
	}

	// test explicit BLANK symbol creation
	v2 := BLANK
	s = NewSymbol(v2)
	if v := s.GetValue(); v != v2 {
		t.Errorf("GetValue() was incorrect, got: %d, want: %d.", v, v2)
	}

}

func TestEqual(t *testing.T) {

	// test equal between two BLANK Symbol
	s1 := NewSymbol()
	s2 := NewSymbol()
	if s1.Equal(s2) == false {
		t.Errorf("Equal() was incorrect, got: %t, want: %t.", false, true)
	}

	// test equal between two Equal Symbol
	s1 = NewSymbol(uint8(1))
	s2 = NewSymbol(uint8(1))
	if s1.Equal(s2) == false {
		t.Errorf("Equal() was incorrect, got: %t, want: %t.", false, true)
	}

	// test equal between two NOT Equal Symbol
	s1 = NewSymbol(uint8(1))
	s2 = NewSymbol(uint8(3))
	if s1.Equal(s2) == true {
		t.Errorf("Equal() was incorrect, got: %t, want: %t.", true, false)
	}

}

func TestCopy(t *testing.T) {

	// test empty symbol creation
	s1 := NewSymbol()
	s2 := s1.Copy()
	if s1.Equal(s2) == false {
		t.Errorf("Copy() was incorrect, got: %t, want: %t.", false, true)
	}

}

func TestIsBlank(t *testing.T) {

	// test IsBlank() over a BLANK Symbol
	s1 := NewSymbol()
	if s1.IsBlank() == false {
		t.Errorf("IsBlank() was incorrect, got: %t, want: %t.", s1, true)
	}

	s2 := NewSymbol(uint8(1))
	// test NewSymbol() over a NOT BLANK Symbol
	if s2.IsBlank() == true {
		t.Errorf("IsBlank() was incorrect, got: %t, want: %t.", s2, false)
	}

}

func TestErase(t *testing.T) {

	// test erase a Symbol
	s := NewSymbol(uint8(1))
	s.Erase()
	if v := s.GetValue(); v != BLANK {
		t.Errorf("Erase was incorrect, got: %d, want: %d.", v, uint8(BLANK))
	}

}
