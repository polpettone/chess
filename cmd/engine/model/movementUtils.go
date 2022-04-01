package model

import (
	"math"
)

func isDiagonalMove(current, target Pos) bool {
	return math.Abs(float64(current.X)-float64(target.X)) == math.Abs(float64(current.Y)-float64(target.Y))
}
