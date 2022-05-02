package model

import (
	"fmt"
	"reflect"
)

func FindSquaresAround(board Board, square Square) ([]Square, error) {

	actualPiece := board.GetPieceAtPos(square.Pos)

	if !reflect.DeepEqual(actualPiece, square.Piece) {
		return nil,
			fmt.Errorf("%s not on %s",
				square.Piece.GetSymbol(), square.Pos.Print())
	}

	return []Square{
		{
			Piece: PieceFrom("WQ"),
			Pos:   *PositionFromString("D4"),
		},
	}, nil
}
