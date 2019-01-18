package set

var exists = struct{}{}

type Set interface {
	Add(v interface{})
	Remove(v interface{})
	Contains(v interface{}) bool
	Cardinality() int
	Iterator() []interface{}
}

type set struct {
	m map[interface{}]struct{}
}

func NewSet() Set {
	s := &set{}
	s.m = make(map[interface{}]struct{})
	return s
}

func (s *set) Add(value interface{}) {
	s.m[value] = exists
}

func (s *set) Remove(value interface{}) {
	delete(s.m, value)
}

func (s *set) Contains(value interface{}) bool {
	_, c := s.m[value]
	return c
}

func (s *set) Cardinality() int {
	if s.m == nil {
		return 0
	}
	return len(s.m)
}

func (s *set) Iterator() []interface{} {
	keys := make([]interface{}, 0)
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}
