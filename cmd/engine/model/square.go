package model

type Square struct {
	Piece Piece
	Pos   Pos
}

func SquareFromString(piece string, pos string) Square {
	return Square{Piece: PieceFrom(piece), Pos: *PositionFromString(pos)}
}
