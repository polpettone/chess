package engine

type Bishop struct {
	Color Color
}

func (p *Bishop) GetColor() Color {
	return p.Color
}

func (p *Bishop) GetSymbol() string {
	if p.Color == WHITE {
		return "WB"
	} else {
		return "BB"
	}
}

func (p *Bishop) Move(current, target Pos, board Board) (*Board, error) {

	if isDiagonalMove(current, target) {
		_, err := board.MovePiece(current, target, p)
		if err != nil {
			return nil, err
		}
		return &board, nil
	}

	return &board, &MoveError{}
}
