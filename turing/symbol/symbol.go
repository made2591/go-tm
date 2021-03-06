package symbol

import "fmt"

// BLANK value
const BLANK = uint8(0)

// Symbol interface
type Symbol interface {
	GetValue() interface{}
	Equal(c Symbol) bool
	Copy() Symbol
	IsBlank() bool
	Print() string
	Erase()
	None()
}

// symbol struct
type symbol struct {
	v interface{}
}

// NewSymbol() Create a new symbol given a value,
// otherwise it creates a BLANK Symbol.
func NewSymbol(v ...interface{}) Symbol {

	// create the symbol
	s := &symbol{}

	// assign the BLANK value
	s.v = BLANK

	// if ONLY a single value is passed
	if len(v) == 1 {

		// if it's different from BLANK
		if v[0] != BLANK {

			// set the value
			s.v = v[0]

		}

	}

	return s

}

// GetValue() Return the primitive value of the Symbol
func (s *symbol) GetValue() interface{} {

	return s.v

}

// Equal(c Symbol) compare s to c and return true
// if their primitive value are equal
func (s *symbol) Equal(c Symbol) bool {

	return s.GetValue() == c.GetValue()

}

// Copy() Create a clone of the s Symbol
func (s *symbol) Copy() Symbol {

	return NewSymbol(s.GetValue())

}

// IsBlank() Return true if the Symbol is BLANK Symbol
func (s *symbol) IsBlank() bool {

	return s.v == BLANK

}

// Print() Return a string representation of the primitive
// value contained in the Symbol
func (s *symbol) Print() string {

	return fmt.Sprintf("%d", s.GetValue())

}

// Erase() Change the Symbol value to BLANK
func (s *symbol) Erase() {

	s.v = BLANK

}

// None() Do nothing on the Symbol
func (s *symbol) None() {}
