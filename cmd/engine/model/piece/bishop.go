package piece

import (
	"fmt"
	"github.com/polpettone/chess/cmd/engine/model/foo"
)

type Bishop struct {
	Color foo.Color
}

func (p *Bishop) GetColor() foo.Color {
	return p.Color
}

func (p *Bishop) GetSymbol() string {
	if p.Color == foo.WHITE {
		return "WB"
	} else {
		return "BB"
	}
}

func (p *Bishop) CheckMoveAllowed(current, target foo.Pos) (bool, error) {
	if isDiagonalMove(current, target) {
		return true, nil
	}
	return false, fmt.Errorf("not allowed")

}
