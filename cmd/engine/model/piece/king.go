package piece

import (
	"fmt"
	"github.com/polpettone/chess/cmd/engine/model/foo"
	"math"
)

type King struct {
	Color foo.Color
}

func (p *King) GetColor() foo.Color {
	return p.Color
}

func (p *King) GetSymbol() string {
	if p.Color == foo.WHITE {
		return "WK"
	} else {
		return "BK"
	}
}

func (p *King) CheckMoveAllowed(current, target foo.Pos) (bool, error) {
	deltaX := math.Abs(float64(current.X) - float64(target.X))
	deltaY := math.Abs(float64(current.Y) - float64(target.Y))
	if deltaX > 1 || deltaY > 1 {
		return false, fmt.Errorf("not allowed")
	}
	return true, nil
}
