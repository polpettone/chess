package piece

import (
	"fmt"
	"github.com/polpettone/chess/cmd/engine/model/foo"
)

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

func (p *Queen) CheckMoveAllowed(current, target foo.Pos) (bool, error) {

	if isDiagonalMove(current, target) {
		return true, nil
	}

	if current.X == target.X || current.Y == target.Y {
		return true, nil
	}
	return false, fmt.Errorf("not allowed")

}
