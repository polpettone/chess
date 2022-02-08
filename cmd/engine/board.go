package engine

import (
	"fmt"
	"reflect"
	"strings"
)

type Square struct {
	Piece Piece
	Pos   Pos
}

type Board struct {
	Fields []*Square
}

func NewEmptyBoard() *Board {
	var fields []*Square

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			fields = append(fields, &Square{Pos: *NewPos(x, y)})
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
					emptyBoard.SetPieceAtPos(*NewPos(x, y), PieceFrom(pieceSymbol))
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
			fields = append(fields, &Square{Pos: *NewPos(x, y)})
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
		board.SetPieceAtPos(*P(k), PieceFrom(v))
	}

	return *board
}

func changePiecesOnBoard(board Board, changes map[string]string) Board {
	for k, v := range changes {
		board.SetPieceAtPos(*P(k), PieceFrom(v))
	}
	return board
}

func (b *Board) GetPieceAtPos(pos Pos) Piece {
	for _, square := range b.Fields {
		if reflect.DeepEqual(square.Pos, pos) {
			return square.Piece
		}
	}
	return nil
}

func (b *Board) SetPieceAtPos(pos Pos, piece Piece) {
	for _, square := range b.Fields {
		if reflect.DeepEqual(square.Pos, pos) {
			square.Piece = piece
		}
	}
}

func (board *Board) Print() string {
	letters := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	numbers := []string{"1", "2", "3", "4", "5", "6", "7", "8"}
	out := "   "

	for x := 0; x < 8; x++ {
		out += fmt.Sprintf(" %s  ", letters[x])
	}
	out += fmt.Sprintln()

	for y := 7; y >= 0; y-- {
		out += fmt.Sprintf("%s ", numbers[y])
		for x := 0; x < 8; x++ {
			piece := board.GetPieceAtPos(*NewPos(x, y))
			if piece != nil {
				out += fmt.Sprintf("[%s]", piece.GetSymbol())
			} else {
				out += "[  ]"
			}
		}
		out += fmt.Sprintf(" %s", numbers[y])
		out += fmt.Sprintln()
	}
	out += "   "
	for x := 0; x < 8; x++ {
		out += fmt.Sprintf(" %s  ", letters[x])
	}

	return out
}
