package state

type State interface {
	NewState(n uint, t *Set, s bool, f bool)
	IsInitial() bool
	IsFinal() bool
	NextTransaction() *Transaction
	GetValue() uint8
}

type state struct {
	v uint
	t *Set
	s bool
	f bool
}

func NewEmptyState() *state {
	s := &state{}
	s.v = 0
	s.t = NewSet()
	s.s = false
	s.f = false
	return s
}

func NewState(n uint, t *Set, i bool, f bool) *state {
	s := &state{}
	s.v = 0
	s.t = NewSet()
	s.s = false
	s.f = false
	return s
}

func (s *state) IsInitial() bool {
	return s.s
}

func (s *state) IsFinal() bool {
	return s.f
}

func (s *state) NextTransaction() bool {
	return true
}
