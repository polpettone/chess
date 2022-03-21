package engine

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"

	"github.com/bclicn/color"
)

type MoveError struct {
	Err        error
	Board      Board
	Piece      Piece
	CurrentPos Pos
	TargetPos  Pos
}

func (m *MoveError) Error() string {
	if m.Err != nil {
		return m.Err.Error()
	} else {
		return "Movement not allowed"
	}
}

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

func SaveBoardToFile(path string, board Board) error {
	err := ioutil.WriteFile(path, []byte(board.Print(nil)), 0755)
	return err
}

const newBoard string = ` 
    A   B   C   D   E   F   G   H  
8 [BR][BN][BB][BQ][BK][BB][BN][BR] 8
7 [BP][BP][BP][BP][BP][BP][BP][BP] 7
6 [  ][  ][  ][  ][  ][  ][  ][  ] 6
5 [  ][  ][  ][  ][  ][  ][  ][  ] 5
4 [  ][  ][  ][  ][  ][  ][  ][  ] 4
3 [  ][  ][  ][  ][  ][  ][  ][  ] 3
2 [WP][WP][WP][WP][WP][WP][WP][WP] 2
1 [WR][WN][WB][WQ][WK][WB][WN][WR] 1
    A   B   C   D   E   F   G   H 
`

func LoadBoardFromFileOrCreateNewBoard(path string) (*Board, error) {
	_, err := os.Open(path)
	if errors.Is(err, os.ErrNotExist) {
		return NewBoardFromString(newBoard)
	} else {
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		return NewBoardFromString(string(content))
	}
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

func (b *Board) MovePiece(current, target Pos, piece Piece) (Piece, error) {

	for _, square := range b.Fields {
		if reflect.DeepEqual(square.Pos, current) {

			if !reflect.DeepEqual(square.Piece, piece) {
				errorMsg := "not allowed "
				errorMsg += fmt.Sprintf("No Piece %s at Pos %s", piece.GetSymbol(), current.String())
				return nil, &MoveError{
					Err: fmt.Errorf(errorMsg),
				}
			}
			square.Piece = nil
		}
	}

	for _, square := range b.Fields {
		if reflect.DeepEqual(square.Pos, target) {
			if square.Piece != nil && square.Piece.GetColor() == piece.GetColor() {
				errorMsg := "not allowed "
				errorMsg += fmt.Sprintf("Piece %s is on %s", square.Piece.GetSymbol(), target.String())
				return nil, &MoveError{
					Err:       fmt.Errorf(errorMsg),
					Board:     *b,
					Piece:     piece,
					TargetPos: target,
				}
			}
			beatenPiece := square.Piece
			square.Piece = piece
			return beatenPiece, nil
		}
	}
	return nil, fmt.Errorf("invalid state")
}

func (board *Board) Print(colorizedPositions []string) string {
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
			pos := NewPos(x, y)

			colorize := false
			if colorizedPositions != nil {
				for _, colorizedPosition := range colorizedPositions {
					p := P(colorizedPosition)
					if reflect.DeepEqual(pos, p) {
						colorize = true
					}
				}
			}

			piece := board.GetPieceAtPos(*pos)

			if piece != nil {
				if colorize {
					out += fmt.Sprintf("[%s]", color.Blue(piece.GetSymbol()))
				} else {
					out += fmt.Sprintf("[%s]", piece.GetSymbol())
				}
			} else {
				emptyFieldSymbol := "[  ]"
				if colorize {
					out += color.Blue(emptyFieldSymbol)
				} else {
					out += emptyFieldSymbol
				}
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
