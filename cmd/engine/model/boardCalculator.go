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

	mostFarStraightPositions := getMostFarStraightPositionsFor(square.Pos)

	squaresAround := []Square{}

	for _, p := range mostFarStraightPositions {

		allPositionsBetween := getAllPositionsBetween(p, square.Pos)

		for _, posBetween := range allPositionsBetween {
			pieceAt := board.GetPieceAtPos(posBetween)
			if pieceAt != nil {
				s := Square{Pos: posBetween, Piece: pieceAt}
				squaresAround = append(squaresAround, s)
			}
		}

	}

	return squaresAround, nil
}
