package model

import (
	"fmt"
)

type Bishop struct {
	Color Color
}

func (p *Bishop) GetColor() Color {
	return p.Color
}

func (p *Bishop) GetSymbol() string {
	if p.Color == WHITE {
		return "WB"
	} else {
		return "BB"
	}
}

func (p *Bishop) CheckMoveAllowed(current, target Pos) (bool, error) {
	if isDiagonalMove(current, target) {
		return true, nil
	}
	return false, fmt.Errorf("not allowed")

}
