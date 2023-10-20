package rect

import "errors"

var PackageName string // this can be accessed outside of the pacakge which is called exported
var moduleName string  // this cannot be accessed since it starts with lowercase which is called unexported

type Rect struct {
	L, B float32
}

func New(l, b float32) (*Rect, error) { // similar to a constructor
	if l == 0 || b == 0 {
		return nil, errors.New("invalid length or bredth")
	}
	return &Rect{L: l, B: b}, nil
}
