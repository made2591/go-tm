package set

var exists = struct{}{}

type Set interface {
	Add(v string)
	Remove(v string)
	Contains(v string) bool
	Cardinality() int
}

type set struct {
	m map[string]struct{}
}

func NewSet() Set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string) {
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *set) Cardinality() int {
	if s.m == nil {
		return 0
	}
	return len(s.m)
}
