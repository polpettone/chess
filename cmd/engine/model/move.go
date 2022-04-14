package model

import (
	"fmt"
	"strings"
)

type Move struct {
	From  Pos
	To    Pos
	Piece Piece
}

func (m Move) Print() string {
	return fmt.Sprintf("%s %s %s", m.Piece.GetSymbol(), m.From.Print(), m.To.Print())
}

func MoveFromString(raw string) (*Move, error) {
	items := strings.Split(raw, " ")
	if len(items) != 3 {
		return nil, fmt.Errorf("invalid raw move")
	}
	piece := PieceFrom(items[0])
	from := PositionFromString(items[1])
	to := PositionFromString(items[2])
	if piece == nil || from == nil || to == nil {
		return nil, fmt.Errorf("invalid raw move")
	}
	return &Move{Piece: piece, From: *from, To: *to}, nil
}
