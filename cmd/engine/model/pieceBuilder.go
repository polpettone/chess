package model

import piece2 "github.com/polpettone/chess/cmd/engine/model/piece"

func PieceFrom(symbol string) Piece {
	if symbol == "" {
		return nil
	}

	if len(symbol) != 2 {
		return nil
	}

	var color piece2.Color

	if string(symbol[0]) == "W" {
		color = piece2.WHITE
	} else {
		color = piece2.BLACK
	}

	var piece Piece

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
