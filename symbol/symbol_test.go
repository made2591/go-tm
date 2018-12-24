package symbol

import "testing"

func TestNewSymbol(t *testing.T) {
	s := NewSymbol()
	if s.v != 0 {
		t.Errorf("NewSymbol was incorrect, got: %d, want: %d.", s.v, uint8(0))
	}
	s = NewSymbol(uint8(1))
	if s.v != uint8(1) {
		t.Errorf("NewSymbol was incorrect, got: %d, want: %d.", s.v, uint8(1))
	}
}
