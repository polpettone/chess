package model

import (
	"fmt"
	"reflect"

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
		return "Move not allowed"
	}
}

type Board struct {
	Fields    []*Square
	Movements []Move
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

func (b *Board) MovePiece(movement Move) (Piece, error) {

	if movement.From == movement.To {
		return nil,
			&MoveError{
				Err: fmt.Errorf("from is equal to position"),
			}
	}

	if len(b.Movements) > 0 {
		lastMove := b.Movements[len(b.Movements)-1]
		if lastMove.Piece.GetColor() == movement.Piece.GetColor() {
			return nil, fmt.Errorf("%s is not on move", movement.Piece.GetSymbol())
		}
	}

	allowed, err := movement.Piece.CheckMoveAllowed(movement.From, movement.To)

	if !allowed {
		return nil, err
	}

	for _, square := range b.Fields {
		if reflect.DeepEqual(square.Pos, movement.From) {

			if !reflect.DeepEqual(square.Piece, movement.Piece) {
				errorMsg := "not allowed "
				errorMsg += fmt.Sprintf("No %s at %s", movement.Piece.GetSymbol(), movement.From.Print())
				return nil, &MoveError{
					Err: fmt.Errorf(errorMsg),
				}
			}
			square.Piece = nil
		}
	}

	for _, square := range b.Fields {
		if reflect.DeepEqual(square.Pos, movement.To) {
			if square.Piece != nil && square.Piece.GetColor() == movement.Piece.GetColor() {
				errorMsg := "not allowed "
				errorMsg += fmt.Sprintf("Piece %s is on %s", square.Piece.GetSymbol(), movement.To.Print())
				return nil, &MoveError{
					Err:       fmt.Errorf(errorMsg),
					Board:     *b,
					Piece:     movement.Piece,
					TargetPos: movement.To,
				}
			}
			beatenPiece := square.Piece
			square.Piece = movement.Piece
			b.Movements = append(b.Movements, movement)
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
					p := PositionFromString(colorizedPosition)
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

	out += fmt.Sprintf("\n")

	for _, movement := range board.Movements {
		out += fmt.Sprintf("%s\n", movement.Print())
	}

	return out
}
