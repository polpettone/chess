package model

import (
	"fmt"
)

type Square struct {
	Piece Piece
	Pos   Pos
}

func SquareFromString(piece string, pos string) Square {
	return Square{Piece: PieceFrom(piece), Pos: *PositionFromString(pos)}
}

func (s Square) String() string {
	return fmt.Sprintf("%s:%s", s.Piece.GetSymbol(), s.Pos.Print())
}
