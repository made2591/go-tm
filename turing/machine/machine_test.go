package machine

import (
	"testing"

	"github.com/made2591/go-tm/set"
	"github.com/made2591/go-tm/turing/state"
	symbol "github.com/made2591/go-tm/turing/symbol"
)

func TestNewTuringMachine(t *testing.T) {
	initialStates := set.NewSet()
	s := state.NewInitialState()
	initialStates.Add(s)
	finalStates := set.NewSet()
	transactions := set.NewSet()
	tr := NewTransaction(state.NewState(uint8(3)), symbol.NewSymbol(uint8(4)), state.NewState(uint8(5)), symbol.NewSymbol(uint8(6)), "P")
	transactions.Add(tr)
	actualState := s
	tm := NewTuringMachine(initialStates, finalStates, transactions, actualState, symbol.NewSymbol(uint8(4)))
	if tm == nil {
		t.Errorf("NewTuringMachine was incorrect, got: nil")
	}
	// newState, symbolWritten, action := transaction.Simulate()
	// if newState.GetValue() != uint8(5) || symbolWritten.GetValue() != uint8(6) || strings.EqualFold(action, "P") {
	// 	t.Errorf("NewTransaction was incorrect, got: %d %d %s, want: %d %d %s.", newState, symbolWritten, action)
	// }
}
