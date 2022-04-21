package model

import (
	"math"
)

func isDiagonalMove(current, target Pos) bool {
	return math.Abs(float64(current.X)-float64(target.X)) ==
		math.Abs(float64(current.Y)-float64(target.Y))
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
