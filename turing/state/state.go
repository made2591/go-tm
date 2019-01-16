package state

import (
	set "github.com/made2591/go-tm/set"
)

const (
	// Initial State Identifier
	INITIAL = ^uint8(0) - 1
	// Final State Identifier
	FINAL = ^uint8(0)
)

// State interface
type State interface {
	IsInitial() bool
	IsFinal() bool
	GetIdentifier() interface{}
}

// state struct
type state struct {
	i interface{}
	t set.Set
}

// NewInitialState() Create a new State of Initial Identifier
func NewInitialState() State {

	s := &state{}
	s.i = INITIAL
	return s

}

// NewState() Create a new State of Initial Identifier
func NewState(v interface{}) State {

	s := &state{}
	s.i = v
	return s

}

// NewFinalState() Create a new State of Final Identifier
func NewFinalState() State {

	s := &state{}
	s.i = FINAL
	return s

}

// IsInitial() Create a new State of Final Identifier
func (s *state) IsInitial() bool {

	return s.i == INITIAL

}

// IsInitial() Create a new State of Final Identifier
func (s *state) IsFinal() bool {

	return s.i == FINAL

}

// GetIdentifier() Create a new State of Final Identifier
func (s *state) GetIdentifier() interface{} {

	return s.i

}
