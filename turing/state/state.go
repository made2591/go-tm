package state

import "strings"

const (
	// Initial State Identifier
	INITIAL = "INITIAL"
	// Final State Identifier
	FINAL = "FINAL"
)

// State interface
type State interface {
	GetIdentifier() interface{}
	IsInitial() bool
	IsFinal() bool
	Equal(s State) bool
}

// state struct
type state struct {
	i interface{}
}

// NewState() Create a new State with initial identifier
func NewState(v interface{}) State {

	s := &state{}
	s.i = v
	return s

}

// NewInitialState() Create a new State with INITIAL Identifier
func NewInitialState() State {

	s := &state{}
	s.i = INITIAL
	return s

}

// NewFinalState() Create a new State with FINAL Identifier
func NewFinalState() State {

	s := &state{}
	s.i = FINAL
	return s

}

// GetIdentifier() Get identifier of a State
func (s *state) GetIdentifier() interface{} {

	return s.i

}

// IsInitial() Check if s State has INITIAL identifier
func (s *state) IsInitial() bool {

	return strings.EqualFold(s.i.(string), INITIAL)

}

// IsFinal() Check if s State is has FINAL identifier
func (s *state) IsFinal() bool {

	return strings.EqualFold(s.i.(string), FINAL)

}

// Equal(c State) Compare s to c and return true
// if their identifier are equal
func (s *state) Equal(c State) bool {

	return strings.EqualFold(s.GetIdentifier().(string), c.GetIdentifier().(string))

}
