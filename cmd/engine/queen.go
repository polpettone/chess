package engine

import "math"

type Queen struct {
	Color Color
}

func (p *Queen) GetColor() Color {
	return p.Color
}

func (p *Queen) GetSymbol() string {
	if p.Color == WHITE {
		return "WQ"
	} else {
		return "BQ"
	}
}

func (p *Queen) Move(current, target Pos, board Board) (*Board, error) {

	if math.Abs(float64(current.X)-float64(target.X)) == math.Abs(float64(current.Y)-float64(target.Y)) {
		return &board, nil
	}

	if current.X == target.X || current.Y == target.Y {
		return &board, nil
	}

	return &board, &MoveError{}
}
