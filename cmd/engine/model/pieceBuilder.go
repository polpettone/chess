package model

import (
	"github.com/polpettone/chess/cmd/engine/model/foo"
	piece2 "github.com/polpettone/chess/cmd/engine/model/piece"
)

func PieceFrom(symbol string) piece2.Piece {
	if symbol == "" {
		return nil
	}

	if len(symbol) != 2 {
		return nil
	}

	var color foo.Color

	if string(symbol[0]) == "W" {
		color = foo.WHITE
	} else {
		color = foo.BLACK
	}

	var piece piece2.Piece

	switch string(symbol[1]) {
	case "P":
		piece = &piece2.Pawn{Color: color}
	case "R":
		piece = &piece2.Rook{Color: color}
	case "N":
		piece = &piece2.Knight{Color: color}
	case "B":
		piece = &piece2.Bishop{Color: color}
	case "Q":
		piece = &piece2.Queen{Color: color}
	case "K":
		piece = &piece2.King{Color: color}
	default:
		piece = nil
	}
	return piece
}
