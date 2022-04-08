package piece

import (
	"fmt"
	"github.com/polpettone/chess/cmd/engine/model/foo"
	"math"
)

func (p *Knight) CheckMoveAllowed(current, target foo.Pos) (bool, error) {

	deltaX := math.Abs(float64(current.X) - float64(target.X))
	deltaY := math.Abs(float64(current.Y) - float64(target.Y))

	if deltaX <= 2 && deltaY <= 2 && (deltaX+deltaY) == 3 {
		return true, nil
	} else {
		msg := "move not allowed from %s to %s"
		return false, fmt.Errorf(msg, current, target)
	}
}

type Knight struct {
	Color foo.Color
}

func (p *Knight) GetColor() foo.Color {
	return p.Color
}

func (p *Knight) GetSymbol() string {
	if p.Color == foo.WHITE {
		return "WN"
	} else {
		return "BN"
	}
}
