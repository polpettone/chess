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
