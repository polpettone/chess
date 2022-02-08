package engine

import (
	"fmt"
	"strconv"
)

type Color int

const (
	BLACK Color = iota
	WHITE
)

type Pos struct {
	X int
	Y int
}

func P(v string) *Pos {
	var x int
	var y int

	if len(v) != 2 {
		return nil
	}

	switch string(v[0]) {
	case "A":
		x = 0
	case "B":
		x = 1
	case "C":
		x = 2
	case "D":
		x = 3
	case "E":
		x = 4
	case "F":
		x = 5
	case "G":
		x = 6
	case "H":
		x = 7
	default:
		return nil
	}

	switch string(v[1]) {

	case "1":
		y = 0
	case "2":
		y = 1
	case "3":
		y = 2
	case "4":
		y = 3
	case "5":
		y = 4
	case "6":
		y = 5
	case "7":
		y = 6
	case "8":
		y = 7
	default:
		return nil
	}

	return NewPos(x, y)
}

func (s Pos) String() string {
	return fmt.Sprintf("(%s, %s)", strconv.Itoa(s.X), strconv.Itoa(s.Y))
}

func NewPos(x, y int) *Pos {
	return &Pos{X: x, Y: y}
}

type Piece interface {
	Move(current, target Pos, board Board) (*Board, error)
	GetColor() Color
	GetSymbol() string
}

type Queen struct {
	Color Color
}

func (p *Queen) GetColor() Color {
	return p.Color
}

func (p *Queen) GetSymbol() string {
	if p.Color == WHITE {
		return "WQ"
	} else {
		return "BQ"
	}
}

func (p *Queen) Move(current, target Pos, board Board) (*Board, error) {
	return &board, nil
}

type King struct {
	Color Color
}

func (p *King) GetColor() Color {
	return p.Color
}

func (p *King) GetSymbol() string {
	if p.Color == WHITE {
		return "WK"
	} else {
		return "BK"
	}
}

func (p *King) Move(current, target Pos, board Board) (*Board, error) {
	return &board, nil
}

type Bishop struct {
	Color Color
}

func (p *Bishop) GetColor() Color {
	return p.Color
}

func (p *Bishop) GetSymbol() string {
	if p.Color == WHITE {
		return "WB"
	} else {
		return "BB"
	}
}

func (p *Bishop) Move(current, target Pos, board Board) (*Board, error) {
	return &board, nil
}

type Rook struct {
	Color Color
}

func (p *Rook) GetColor() Color {
	return p.Color
}

func (p *Rook) GetSymbol() string {
	if p.Color == WHITE {
		return "WR"
	} else {
		return "BR"
	}
}

func (p *Rook) Move(current, target Pos, board Board) (*Board, error) {
	return &board, nil
}

type Knight struct {
	Color Color
}

func (p *Knight) GetColor() Color {
	return p.Color
}

func (p *Knight) GetSymbol() string {
	if p.Color == WHITE {
		return "WN"
	} else {
		return "BN"
	}
}

func (p *Knight) Move(current, target Pos, board Board) (*Board, error) {
	return &board, nil
}

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
