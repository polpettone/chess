package piece

import (
	"github.com/polpettone/chess/cmd/engine/model/foo"
	"math"
)

func isDiagonalMove(current, target foo.Pos) bool {
	return math.Abs(float64(current.X)-float64(target.X)) == math.Abs(float64(current.Y)-float64(target.Y))
}
