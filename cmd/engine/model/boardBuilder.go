package model

import (
	"github.com/polpettone/chess/cmd/engine/model/piece"
	"strings"
)

func NewEmptyBoard() *Board {
	var fields []*Square
	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			fields = append(fields, &Square{Pos: *piece.NewPos(x, y)})
		}
	}
	board := &Board{Fields: fields}
	return board
}

func NewBoardFromString(value string) (*Board, error) {
	slice := strings.Split(value, "\n")
	y := 7
	emptyBoard := NewEmptyBoard()
	for _, line := range slice {
		if strings.Contains(line, "[") {
			lineSlice := strings.Split(line, "[")
			x := 0
			for _, l := range lineSlice {
				if strings.Contains(l, "]") {
					pieceSymbol := l[0:2]
					emptyBoard.SetPieceAtPos(*piece.NewPos(x, y), piece.PieceFrom(pieceSymbol))
					x = x + 1
				}
			}
			y = y - 1
		}
	}
	return emptyBoard, nil
}

func NewBoard() Board {
	var fields []*Square

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			fields = append(fields, &Square{Pos: *piece.NewPos(x, y)})
		}
	}

	board := &Board{Fields: fields}

	setup := map[string]string{

		"A2": "WP",
		"B2": "WP",
		"C2": "WP",
		"D2": "WP",
		"E2": "WP",
		"F2": "WP",
		"G2": "WP",
		"H2": "WP",

		"A7": "BP",
		"B7": "BP",
		"C7": "BP",
		"D7": "BP",
		"E7": "BP",
		"F7": "BP",
		"G7": "BP",
		"H7": "BP",

		"A1": "WR",
		"B1": "WN",
		"C1": "WB",
		"D1": "WQ",
		"E1": "WK",
		"F1": "WB",
		"G1": "WN",
		"H1": "WR",

		"A8": "BR",
		"B8": "BN",
		"C8": "BB",
		"D8": "BQ",
		"E8": "BK",
		"F8": "BB",
		"G8": "BN",
		"H8": "BR",
	}

	for k, v := range setup {
		board.SetPieceAtPos(*piece.P(k), piece.PieceFrom(v))
	}

	return *board
}
