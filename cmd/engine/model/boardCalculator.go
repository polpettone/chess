package model

func FindSquaresAround(board Board, square Square) ([]Square, error) {
	return []Square{
		{
			Piece: PieceFrom("WQ"),
			Pos:   *PositionFromString("D4"),
		},
	}, nil
}
