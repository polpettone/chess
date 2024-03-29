package model

import (
	"fmt"
)

type Rook struct {
	Color Color
}

func (p *Rook) GetColor() Color {
	return p.Color
}

func (p *Rook) GetSymbol() string {
	if p.Color == WHITE {
		return "WR"
	} else {
		return "BR"
	}
}

func (p *Rook) CheckMoveAllowed(current, target Pos) (bool, error) {
	if current.X != target.X {
		if current.Y != target.Y {
			return false, fmt.Errorf("not allowed")
		}
	}
	return true, nil
}
