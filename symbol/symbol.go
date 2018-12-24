package symbol

const BLANK = uint8(2)

type Symbol interface {
	NewSymbol(v ...uint8) *symbol
	IsBlank() bool
	GetValue() uint8
}

type symbol struct {
	v uint8
}

func NewSymbol(v ...uint8) *symbol {
	s := &symbol{}
	s.v = uint8(2)
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

func (s *symbol) GetValue() uint8 {
	return s.v
}
