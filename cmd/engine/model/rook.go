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

func (p *Rook) Move(current, target Pos, board Board) (*Board, error) {

	if current.X != target.X {
		if current.Y != target.Y {
			return nil, fmt.Errorf("not allowed")
		}
	}

	_, err := board.MovePiece(current, target, p)

	if err != nil {
		return &board, err
	}

	return &board, nil
}
