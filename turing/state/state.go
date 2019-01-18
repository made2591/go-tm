package state

const (
	// Initial State Identifier
	INITIAL = ^uint8(0) - 1
	// Final State Identifier
	FINAL = ^uint8(0)
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

	return s.i == INITIAL

}

// IsFinal() Check if s State is has FINAL identifier
func (s *state) IsFinal() bool {

	return s.i == FINAL

}

// Equal(c State) Compare s to c and return true
// if their identifier are equal
func (s *state) Equal(c State) bool {

	return s.GetIdentifier() == c.GetIdentifier()

}
