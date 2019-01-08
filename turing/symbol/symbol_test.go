package symbol

import "testing"

func TestNewSymbol(t *testing.T) {
	s := NewSymbol()
	if v := s.GetValue(); v != uint8(BLANK) {
		t.Errorf("NewSymbol was incorrect, got: %d, want: %d.", v, BLANK)
	}
	s = NewSymbol(uint8(1))
	if v := s.GetValue(); v != uint8(1) {
		t.Errorf("NewSymbol was incorrect, got: %d, want: %d.", v, uint8(1))
	}
}

func TestIsBlank(t *testing.T) {
	s := NewSymbol()
	if s.IsBlank() == false {
		t.Errorf("IsBlank was incorrect, got: %t, want: %t.", false, true)
	}
}

func TestGetValue(t *testing.T) {
	s := NewSymbol(uint8(1))
	if v := s.GetValue(); v != uint8(1) {
		t.Errorf("GetValue was incorrect, got: %d, want: %d.", v, uint8(1))
	}
}

func TestErase(t *testing.T) {
	s := NewSymbol(uint8(1))
	s.Erase()
	if v := s.GetValue(); v != BLANK {
		t.Errorf("Erase was incorrect, got: %d, want: %d.", v, uint8(BLANK))
	}
}
