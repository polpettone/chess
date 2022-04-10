package model

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/polpettone/chess/cmd/engine/model/foo"
	"github.com/polpettone/chess/cmd/engine/model/piece"

	"github.com/bclicn/color"
)

type MoveError struct {
	Err        error
	Board      Board
	Piece      piece.Piece
	CurrentPos foo.Pos
	TargetPos  foo.Pos
}

func (m *MoveError) Error() string {
	if m.Err != nil {
		return m.Err.Error()
	} else {
		return "Movement not allowed"
	}
}

type Square struct {
	Piece piece.Piece
	Pos   foo.Pos
}

type Board struct {
	Fields    []*Square
	Movements []Movement
}

type Movement struct {
	From  foo.Pos
	To    foo.Pos
	Piece piece.Piece
}

func MoveFromString(raw string) (*Movement, error) {
	items := strings.Split(raw, " ")
	if len(items) != 3 {
		return nil, fmt.Errorf("invalid raw move")
	}
	piece := PieceFrom(items[0])
	from := foo.PositionFromString(items[1])
	to := foo.PositionFromString(items[2])
	if piece == nil || from == nil || to == nil {
		return nil, fmt.Errorf("invalid raw move")
	}
	return &Movement{Piece: piece, From: *from, To: *to}, nil
}

func (b *Board) GetPieceAtPos(pos foo.Pos) piece.Piece {
	for _, square := range b.Fields {
		if reflect.DeepEqual(square.Pos, pos) {
			return square.Piece
		}
	}
	return nil
}

func (b *Board) SetPieceAtPos(pos foo.Pos, piece piece.Piece) {
	for _, square := range b.Fields {
		if reflect.DeepEqual(square.Pos, pos) {
			square.Piece = piece
		}
	}
}

func (b *Board) MovePiece(current, target foo.Pos, piece piece.Piece) (piece.Piece, error) {

	allowed, err := piece.CheckMoveAllowed(current, target)

	if !allowed {
		return nil, err
	}

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
			b.Movements = append(b.Movements, Movement{From: current, To: target, Piece: piece})
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
			pos := foo.NewPos(x, y)

			colorize := false
			if colorizedPositions != nil {
				for _, colorizedPosition := range colorizedPositions {
					p := foo.PositionFromString(colorizedPosition)
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
