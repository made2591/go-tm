package machine

import (
	"strings"
	"testing"
)

func NewTransaction(t *testing.T) {
	transaction := NewTransaction(nil, nil, NewState(uint8(5)), NewSymbol(uint8(6)), "P")
	if transaction == nil {
		t.Errorf("NewTransaction was incorrect, got: nil")
	}
	newState, symbolWritten, action := transaction.Simulate()
	if newState.GetValue() != uint8(5) || symbolWritten.GetValue() != uint8(6) || strings.EqualFold(action, "P") {
		t.Errorf("NewTransaction was incorrect, got: %d %d %s, want: %d %d %s.", newState, symbolWritten, action)
	}
}
