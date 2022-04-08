package piece

func PieceFrom(symbol string) Piece {
	if symbol == "" {
		return nil
	}

	if len(symbol) != 2 {
		return nil
	}

	var color Color

	if string(symbol[0]) == "W" {
		color = WHITE
	} else {
		color = BLACK
	}

	var piece Piece

	switch string(symbol[1]) {
	case "P":
		piece = &Pawn{Color: color}
	case "R":
		piece = &Rook{Color: color}
	case "N":
		piece = &Knight{Color: color}
	case "B":
		piece = &Bishop{Color: color}
	case "Q":
		piece = &Queen{Color: color}
	case "K":
		piece = &King{Color: color}
	default:
		piece = nil
	}
	return piece
}
