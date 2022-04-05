package model

import (
	"math"
)

type King struct {
	Color Color
}

func (p *King) GetColor() Color {
	return p.Color
}

func (p *King) GetSymbol() string {
	if p.Color == WHITE {
		return "WK"
	} else {
		return "BK"
	}
}

func (p *King) CheckMoveAllowed(current, target Pos) (bool, error) {
	deltaX := math.Abs(float64(current.X) - float64(target.X))
	deltaY := math.Abs(float64(current.Y) - float64(target.Y))
	if deltaX > 1 || deltaY > 1 {
		return false, &MoveError{}
	}
	return true, nil
}
