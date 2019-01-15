package symbol

import "fmt"

const BLANK = uint8(2)

type Symbol interface {
	IsBlank() bool
	GetValue() interface{}
	Print()
	Erase()
	None()
	Equal(c Symbol) bool
}

type symbol struct {
	v interface{}
}

func NewSymbol(v ...interface{}) Symbol {
	s := &symbol{}
	s.v = BLANK
	if len(v) == 1 {
		if v[0] != BLANK {
			s.v = v[0]
		}
	}
	return s
}

func (s *symbol) IsBlank() bool {
	return s.v == BLANK
}

func (s *symbol) GetValue() interface{} {
	return s.v
}

func (s *symbol) Print() {
	fmt.Print(s.v)
}

func (s *symbol) Erase() {
	s.v = BLANK
}

func (s *symbol) None() {
}

func (s *symbol) Equal(c Symbol) bool {
	return s.GetValue() == c.GetValue()
}
