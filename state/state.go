package state

import (
	set "github.com/made2591/go-tm/set"
)

const (
	INITIAL = ^uint8(0) - 1
	FINAL   = ^uint8(0)
)

type State interface {
	IsInitial() bool
	IsFinal() bool
	GetValue() uint8
	NextTransaction() Transaction
}

type state struct {
	v uint8
	t set.Set
}

func NewInitialState() State {
	s := &state{}
	s.v = INITIAL
	s.t = set.NewSet()
	return s
}

func NewState(v uint8, t set.Set) State {
	s := &state{}
	s.v = v
	s.t = t
	return s
}

func (s *state) IsInitial() bool {
	return s.v == INITIAL
}

func (s *state) IsFinal() bool {
	return s.v == FINAL
}

func (s *state) GetValue() uint8 {
	return s.v
}

func (s *state) NextTransaction() Transaction {
	// TODO
	return &transaction{}
}
