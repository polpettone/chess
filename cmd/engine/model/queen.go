package model

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

func (p *Queen) Move(current, target Pos, board Board) (*Board, error) {

	if isDiagonalMove(current, target) {
		_, err := board.MovePiece(current, target, p)
		if err != nil {
			return nil, err
		}
		return &board, nil
	}

	if current.X == target.X || current.Y == target.Y {
		_, err := board.MovePiece(current, target, p)
		if err != nil {
			return nil, err
		}
		return &board, nil
	}
	return &board, &MoveError{}
}
