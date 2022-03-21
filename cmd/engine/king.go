package engine

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

func (p *King) Move(current, target Pos, board Board) (*Board, error) {

	deltaX := math.Abs(float64(current.X) - float64(target.X))
	deltaY := math.Abs(float64(current.Y) - float64(target.Y))

	if deltaX > 1 || deltaY > 1 {
		return &board, &MoveError{}
	}

	_, err := board.MovePiece(current, target, p)

	if err != nil {
		return &board, err
	}

	return &board, nil
}
