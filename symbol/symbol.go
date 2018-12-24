package symbol

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
	s.v = uint8(0)
	if len(v) > 0 {
		s.v = v[0]
	}
	return s
}

func (s *symbol) IsBlank() bool {
	return s.v == 2
}

func (s *symbol) GetValue() uint8 {
	return s.v
}
