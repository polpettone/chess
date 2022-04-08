package piece

import (
	"fmt"
	"github.com/polpettone/chess/cmd/engine/model/foo"
)

type Rook struct {
	Color foo.Color
}

func (p *Rook) GetColor() foo.Color {
	return p.Color
}

func (p *Rook) GetSymbol() string {
	if p.Color == foo.WHITE {
		return "WR"
	} else {
		return "BR"
	}
}

func (p *Rook) CheckMoveAllowed(current, target foo.Pos) (bool, error) {
	if current.X != target.X {
		if current.Y != target.Y {
			return false, fmt.Errorf("not allowed")
		}
	}
	return true, nil
}
