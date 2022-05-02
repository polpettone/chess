package model

import (
	"math"
	"reflect"
)

func getMostFarStraightPositionsFor(pos Pos) []Pos {

	positions := []Pos{}

	x0 := *NewPos(0, pos.Y)
	if !reflect.DeepEqual(x0, pos) {
		positions = append(positions, x0)
	}

	x7 := *NewPos(7, pos.Y)
	if !reflect.DeepEqual(x7, pos) {
		positions = append(positions, x7)
	}

	y0 := *NewPos(pos.X, 0)
	if !reflect.DeepEqual(y0, pos) {
		positions = append(positions, y0)
	}

	y7 := *NewPos(pos.X, 7)
	if !reflect.DeepEqual(y7, pos) {
		positions = append(positions, y7)
	}

	return positions
}

func isDiagonalMove(from, to Pos) bool {
	return math.Abs(float64(from.X)-float64(to.X)) ==
		math.Abs(float64(from.Y)-float64(to.Y))
}

func isDiagonalMoveDistanceOne(from, to Pos) bool {
	absX := math.Abs(float64(from.X) - float64(to.X))
	absY := math.Abs(float64(from.Y) - float64(to.Y))
	return absX == absY && absY == 1 && absX == 1
}

func getAllPositionsBetween(from, to Pos) []Pos {
	positions := []Pos{}

	if from.X == to.X && from.Y < to.Y {
		for y := from.Y + 1; y < to.Y; y++ {
			positions = append(positions, *NewPos(from.X, y))
		}
	}

	if from.X == to.X && from.Y > to.Y {
		for y := to.Y + 1; y < from.Y; y++ {
			positions = append(positions, *NewPos(from.X, y))
		}
	}

	if from.Y == to.Y && from.X < to.X {
		for n := from.X + 1; n < to.X; n++ {
			positions = append(positions, *NewPos(n, from.Y))
		}
	}

	if from.Y == to.Y && from.X > to.X {
		for n := to.X + 1; n < from.X; n++ {
			positions = append(positions, *NewPos(n, from.Y))
		}
	}

	if isDiagonalMove(from, to) {

		if to.X < from.X && to.Y < from.Y {
			m := to.Y + 1
			for n := to.X + 1; n < from.X; n++ {
				positions = append(positions, *NewPos(n, m))
				m++
			}
		}
		if to.X > from.X && to.Y > from.Y {
			m := from.Y + 1
			for n := from.X + 1; n < to.X; n++ {
				positions = append(positions, *NewPos(n, m))
				m++
			}
		}

		if to.X < from.X && to.Y > from.Y {
			m := to.Y - 1
			for n := to.X + 1; n < from.X; n++ {
				positions = append(positions, *NewPos(n, m))
				m--
			}
		}

		if to.X > from.X && to.Y < from.Y {
			m := from.Y - 1
			for n := from.X + 1; n < to.X; n++ {
				positions = append(positions, *NewPos(n, m))
				m--
			}
		}
	}

	return positions
}
