package piece

import (
	"fmt"
	"github.com/polpettone/chess/cmd/engine/model/foo"
)

type Queen struct {
	Color foo.Color
}

func (p *Queen) GetColor() foo.Color {
	return p.Color
}

func (p *Queen) GetSymbol() string {
	if p.Color == foo.WHITE {
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
