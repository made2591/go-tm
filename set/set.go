package set

import (
	"math/rand"
)

var exists = struct{}{}

type Set interface {
	Add(v interface{})
	Remove(v interface{})
	GetOne() interface{}
	Contains(v interface{}) bool
	Cardinality() int
	Iterator() <-chan interface{}
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

func (s *set) GetOne() interface{} {
	return randIntMapKey
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

func (s *set) Iterator() <-chan interface{} {
	keys := make(chan interface{}, len(s.m))
	for k := range s.m {
		keys <- k
	}
	close(keys)
	return keys
}

func randIntMapKey(m map[interface{}]interface{}) interface{} {
	for _, i := range rand.Perm(len(m)) {
		return i
	}
	return nil
}
