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
	GetValue() interface{}
}

type state struct {
	v interface{}
	t set.Set
}

func NewInitialState() State {
	s := &state{}
	s.v = INITIAL
	s.t = set.NewSet()
	return s
}

func NewState(v interface{}, t set.Set) State {
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

func (s *state) GetValue() interface{} {
	return s.v
}
