package ast

import (
	"fmt"

	"github.com/tomratford/day-19/token"
)

type Part struct {
	X int
	M int
	A int
	S int
}

func (p Part) GetValue(s token.Type) (int, error) {
	switch s {
	case "XPART":
		return p.X, nil
	case "MPART":
		return p.M, nil
	case "APART":
		return p.A, nil
	case "SPART":
		return p.S, nil
	default:
		return 0, fmt.Errorf("Part %q does not exist", s)
	}
}

func (p Part) Sum() int {
	return p.X + p.M + p.A + p.S
}
