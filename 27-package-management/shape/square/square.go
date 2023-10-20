package square

import "errors"

func New(s float32) (sq *Square, err error) {

	if s == 0 {
		return nil, errors.New("invalid side")
	}
	sq = new(Square)
	sq.S = s
	return sq, nil
}

type Square struct {
	//s float32 // unexported so cannot be called outside of this package
	S float32
}

func (s *Square) Area() float32 {
	return s.S * s.S
}

func (s *Square) Perimeter() float32 {
	return 4 * s.S
}
