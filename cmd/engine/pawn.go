package engine

import (
	"fmt"
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

func (p *Pawn) Move(current, target Pos, board Board) (*Board, error) {

	pieceOnCurrentPos := board.GetPieceAtPos(current)
	if pieceOnCurrentPos == nil {
		return nil, fmt.Errorf("no piece at current pos")
	}

	if pieceOnCurrentPos.GetSymbol() != p.GetSymbol() {
		return nil, fmt.Errorf("wrong piece on current pos")
	}

	if p.Color == WHITE {
		if current.Y > target.Y {
			return nil, fmt.Errorf("not allowed")
		}
	}

	if p.Color == BLACK {
		if current.Y < target.Y {
			return nil, fmt.Errorf("not allowed")
		}
	}

	if current.Y == target.Y && current.X != target.X {
		return nil, fmt.Errorf("not allowed")
	}

	if current.X == target.X && board.GetPieceAtPos(target) != nil {
		return nil, fmt.Errorf("not allowed")
	}

	if p.Color == WHITE {
		if current.Y == 1 {
			if (target.Y - current.Y) > 2 {
				return nil, fmt.Errorf("not allowed")
			} else {
				targetPiece := board.GetPieceAtPos(target)
				if targetPiece != nil && targetPiece.GetColor() == WHITE {
					return nil, fmt.Errorf("not allowed")
				}
				board.SetPieceAtPos(target, p)
				board.SetPieceAtPos(current, nil)
			}
		} else {
			if (target.Y - current.Y) > 1 {
				return nil, fmt.Errorf("not allowed")
			} else {
				targetPiece := board.GetPieceAtPos(target)
				if targetPiece != nil && targetPiece.GetColor() == WHITE {
					return nil, fmt.Errorf("not allowed")
				}
				board.SetPieceAtPos(target, p)
				board.SetPieceAtPos(current, nil)
			}
		}
	}

	if p.Color == BLACK {
		if current.Y == 6 {
			if (current.Y - target.Y) > 2 {
				return nil, fmt.Errorf("not allowed")
			} else {
				targetPiece := board.GetPieceAtPos(target)
				if targetPiece != nil && targetPiece.GetColor() == BLACK {
					return nil, fmt.Errorf("not allowed")
				}
				board.SetPieceAtPos(target, p)
				board.SetPieceAtPos(current, nil)
			}
		} else {
			if (current.Y - target.Y) > 1 {
				return nil, fmt.Errorf("not allowed")
			} else {
				targetPiece := board.GetPieceAtPos(target)
				if targetPiece != nil && targetPiece.GetColor() == BLACK {
					return nil, fmt.Errorf("not allowed")
				}
				board.SetPieceAtPos(target, p)
				board.SetPieceAtPos(current, nil)
			}
		}
	}

	return &board, nil
}
