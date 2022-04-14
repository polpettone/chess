package model

type Piece interface {
	CheckMoveAllowed(current, target Pos) (bool, error)
	GetColor() Color
	GetSymbol() string
}
