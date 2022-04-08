package piece

import (
	"fmt"
	"github.com/polpettone/chess/cmd/engine/model/foo"
)

type Pawn struct {
	Color Color
}

func (p *Pawn) GetColor() Color {
	return p.Color
}

func (p *Pawn) GetSymbol() string {
	if p.Color == WHITE {
		return "WP"
	} else {
		return "BP"
	}
}

func (p *Pawn) CheckMoveAllowed(current, target foo.Pos) (bool, error) {
	if p.Color == WHITE {
		if current.Y > target.Y {
			return false, fmt.Errorf("not allowed")
		}
	}

	if p.Color == BLACK {
		if current.Y < target.Y {
			return false, fmt.Errorf("not allowed")
		}
	}

	if current.Y == target.Y && current.X != target.X {
		return false, fmt.Errorf("not allowed")
	}

	if p.Color == WHITE {
		if current.Y == 1 {
			if (target.Y - current.Y) > 2 {
				return false, fmt.Errorf("not allowed")
			}
		} else {
			if (target.Y - current.Y) > 1 {
				return false, fmt.Errorf("not allowed")
			}
		}
	}

	if p.Color == BLACK {
		if current.Y == 6 {
			if (current.Y - target.Y) > 2 {
				return false, fmt.Errorf("not allowed")
			}
		} else {
			if (current.Y - target.Y) > 1 {
				return false, fmt.Errorf("not allowed")
			}
		}
	}
	return true, nil
}
