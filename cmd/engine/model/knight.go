package model

import (
	"fmt"
	"math"
)

func (p *Knight) Move(current, target Pos, board Board) (*Board, error) {

	deltaX := math.Abs(float64(current.X) - float64(target.X))
	deltaY := math.Abs(float64(current.Y) - float64(target.Y))

	if deltaX <= 2 && deltaY <= 2 && (deltaX+deltaY) == 3 {
		_, err := board.MovePiece(current, target, p)
		if err != nil {
			return &board, err
		}
		return &board, nil
	} else {
		msg := "move not allowed from %s to %s"
		return &board, &MoveError{
			Err: fmt.Errorf(msg, current, target),
		}
	}
}

func (p *Knight) CheckMoveAllowed(current, target Pos) (bool, error) {

	deltaX := math.Abs(float64(current.X) - float64(target.X))
	deltaY := math.Abs(float64(current.Y) - float64(target.Y))

	if deltaX <= 2 && deltaY <= 2 && (deltaX+deltaY) == 3 {
		return true, nil
	} else {
		msg := "move not allowed from %s to %s"
		return false, &MoveError{
			Err: fmt.Errorf(msg, current, target),
		}
	}
}

type Knight struct {
	Color Color
}

func (p *Knight) GetColor() Color {
	return p.Color
}

func (p *Knight) GetSymbol() string {
	if p.Color == WHITE {
		return "WN"
	} else {
		return "BN"
	}
}
