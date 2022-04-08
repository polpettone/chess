package piece

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
	CheckMoveAllowed(current, target Pos) (bool, error)
	GetColor() Color
	GetSymbol() string
}
